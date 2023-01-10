/*
Copyright © 2023 Jenda Mudron

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/jenmud/consensus/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func sqliteClient(dsn string) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	return client, client.Schema.Create(context.Background())
}

// entrypoint is the main entrypoint to the applications.
func entrypoint(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "consensus",
	Short: "Consensus helps building and maintaining software by describing the expected behavior.",
	Long: `Consensus helps building and maintaining software by describing the expected behavior.
It focuses on the high level expectations rather then low level details encouraging a wider
participation from stakeholders, collaborators, and developers.`,
	RunE: entrypoint,
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
	rootCmd.Flags().String("db", "file:consensus.sqlite?mode=memory&cache=shared&_fk=1", "DSN string for connecting to the database")
	rootCmd.Flags().String("dialect", "sqlite3", "Dialect used for the sql driver")
	rootCmd.Flags().Bool("debug", false, "enable sql debugging")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
