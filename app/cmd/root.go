/*
Copyright © 2024 Jenda Mudron

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/jenmud/consensus/business/service"
	"github.com/jenmud/consensus/foundation/data/sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "consensus",
	Short: "Consensus is a project management tool for teams.",
	Long: `Consensus is a project management tool for teams that provides a platform for
collaboration and progress tracking. It is built around behavior-driven development
(BDD) and is language agnostic, making it a great fit for teams of any size
or composition.
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
		slog.Info("starting Consensus...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		listener, err := net.Listen("tcp", viper.GetString("address"))
		if err != nil {
			slog.Error("failed to start Consensus", slog.String("reason", err.Error()))
			return
		}

		defer listener.Close()

		logger := slog.With(
			slog.String("address", listener.Addr().String()),
			slog.String("dsn", viper.GetString("dsn")),
		)

		done := make(chan error, 1)

		go func() {
			db, err := sqlite.NewDB(viper.GetString("dsn"))
			if err != nil {
				done <- err
				return
			}

			ss := service.New(db)
			opts := []grpc.ServerOption{}
			server := grpc.NewServer(opts...)
			service.RegisterConsensusServer(server, ss)
			logger.Info("started Consensus")
			done <- server.Serve(listener)
		}()

		select {
		case <-ctx.Done():
			logger.Info("stopped Consensus", slog.String("reason", ctx.Err().Error()))
		case e := <-done:
			slog.Error("stopped Consensus", slog.String("reason", e.Error()))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.consensus.yaml)")
	rootCmd.Flags().StringP("address", "a", ":8000", "Address to listen and accept connections on.")
	rootCmd.Flags().StringP("dsn", "d", "file:consensus.sqlite", "Data source name for the database.")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".app" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".consensus")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
