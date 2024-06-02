/*
Copyright © 2024 Jenda Mudron
*/
package cmd

import (
	"log/slog"

	"github.com/jenmud/consensus/app/cmd/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uiCmd represents the dashboard command
var uiCmd = &cobra.Command{
	Use:     "ui",
	Aliases: []string{"dashboard", "UI"},
	Short:   "UI for the Consensus. Provides a web-based dashboard for projects, features and tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		err := make(chan error)
		go func() {
			err <- ui.ListenAndServe(ctx, viper.GetString("address"), nil)
		}()

		select {
		case <-ctx.Done():
		case e := <-err:
			slog.Error("Shutdown...", slog.String("reason", e.Error()))
		}

	},
}

func init() {
	webCmd.AddCommand(uiCmd)
	uiCmd.Flags().StringP("address", "a", "0.0.0.0:8080", "Address is the address on which the server should listen for incoming requests.")
	viper.BindPFlags(uiCmd.Flags())
}
