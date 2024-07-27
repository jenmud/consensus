/*
Copyright © 2024 Jenda Mudron
*/
package cmd

import (
	"context"
	"log/slog"
	"net"

	"github.com/jenmud/consensus/app/cmd/ui"
	"github.com/jenmud/consensus/business/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// uiCmd represents the dashboard command
var uiCmd = &cobra.Command{
	Use:     "ui",
	Aliases: []string{"dashboard", "UI"},
	Short:   "UI for the Consensus. Provides a web-based dashboard for projects, features and tasks.",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		logger := slog.With(
			slog.Group(
				"server",
				slog.String("dsn", viper.GetString("dsn")),
				slog.String("address", viper.GetString("address")),
				slog.String("service", viper.GetString("service")),
			),
		)

		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, "tcp", addr)
			}),
		}

		conn, err := grpc.Dial(viper.GetString("service"), opts...)
		if err != nil {
			logger.Error("failed to start ui server", slog.String("reason", err.Error()))
			panic(err)
		}

		client := service.NewConsensusClient(conn)

		errChan := make(chan error)
		go func() {
			errChan <- ui.ListenAndServe(ctx, viper.GetString("address"), client, logger)
		}()

		select {
		case <-ctx.Done():
		case e := <-errChan:
			slog.Error("Shutdown...", slog.String("reason", e.Error()))
		}

	},
}

func init() {
	webCmd.AddCommand(uiCmd)
	uiCmd.Flags().StringP("address", "a", "0.0.0.0:8080", "Address is the address on which the server should listen for incoming requests.")
	uiCmd.Flags().StringP("service", "s", "0.0.0.0:8000", "Service is the consensus service address to connect to.")
}
