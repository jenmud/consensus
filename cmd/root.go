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
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jenmud/consensus/ent"
	"github.com/jenmud/consensus/graph"
	"github.com/jenmud/consensus/graph/generated"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// sqliteClient returns a sqlite3 ent client.
func sqliteClient(dsn string) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	return client, client.Schema.Create(context.Background())
}

// Defining the Graphql handler
func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: graph.New(client)},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandleer returns a GraphQL handler.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("Consensus GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// servePlayground serves the GraphQL playground.
func servePlayground(ctx context.Context, addr string, client *ent.Client, debug bool) error {
	go func() {
		if !debug {
			gin.SetMode(gin.ReleaseMode)
		}

		r := gin.Default()
		r.POST("/query", graphqlHandler(client))
		r.GET("/", playgroundHandler())

		s := &http.Server{
			Addr:           addr,
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		log.Printf("listening and accepting client connections on %s", addr)
		s.ListenAndServe()
	}()

	<-ctx.Done()
	return nil
}

// entrypoint is the main entrypoint to the applications.
func entrypoint(cmd *cobra.Command, args []string) error {
	var client *ent.Client
	var err error

	debugging, err := strconv.ParseBool(cmd.Flags().Lookup("debug").Value.String())
	if err != nil {
		return err
	}

	dialect := cmd.Flags().Lookup("dialect").Value.String()
	dsn := cmd.Flags().Lookup("db").Value.String()

	switch strings.ToLower(dialect) {
	case "sqlite3":
		client, err = sqliteClient(dsn)
	}

	if err != nil {
		return err
	}

	/*
		client.Use(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					start := time.Now()

					defer func() {
						log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
					}()

					v, err := next.Mutate(ctx, m)
					if err != nil {
						return v, err
					}

					switch t := v.(type) {
					case *ent.Comment:
						added := m.Fields()
						id := t.ID
						text := t.Text
						for _, f := range added {
							o, err := m.OldField(ctx, f)
							if err != nil {
								continue
							}
							log.Printf("Op=%s\tType=%s\tConcreteType=%T\tID:%d\tText:%s\tAddedFields: %v\tOld: %s\n", m.Op(), m.Type(), m, id, text, added, o)
						}
					case *ent.Epic:
						added := m.Fields()
						id := t.ID
						name := t.Name
						desp := t.Description
						log.Printf("Op=%s\tType=%s\tConcreteType=%T\tID:%d\tText:%s\tDescription:%s\tAddedFields: %v\n", m.Op(), m.Type(), m, id, name, desp, added)
					}

					log.Printf("--> Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
					return v, err
				})
			},
		)
	*/

	addr := cmd.Flags().Lookup("server").Value.String()
	return servePlayground(cmd.Context(), addr, client, debugging)
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
	rootCmd.Flags().StringP("server", "s", "0.0.0.0:8080", "address to listen and accept client connections")
	rootCmd.Flags().String("db", "file:consensus.sqlite?mode=memory&cache=shared&_fk=1", "DSN string for connecting to the database")
	rootCmd.Flags().String("dialect", "sqlite3", "dialect used for the sql driver")
	rootCmd.Flags().Bool("debug", false, "enable sql debugging")
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
}
