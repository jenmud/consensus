package ui

import (
	"context"
	"embed"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"time"
)

//go:embed templates/*.tmpl
var embedded embed.FS

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// registerRoutes registers the routes for the HTTP server.
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", index)

	staticFS := http.FileServer(http.FS(embedded))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFS))
}

// Run starts the HTTP server and serves the static files and the index page.
// It takes an address as a parameter, which is the address on which the server
// should listen for incoming requests. The address should be in the format
// `host:port`.
//
// Returns an error if there was a problem starting the server.
func ListenAndServe(ctx context.Context, addr string, logger *slog.Logger) error {
	if logger == nil {
		logger = slog.Default()
	}

	slogger := logger.With(slog.String("address", addr))
	mux := http.NewServeMux()
	registerRoutes(mux)

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			serverCtx := context.WithoutCancel(ctx)
			slogger.Info("starting ui server", slog.String("address", l.Addr().String()))
			return serverCtx
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			slogger.Info("peer connection", slog.String("peer", c.RemoteAddr().String()))
			return ctx
		},
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return server.Shutdown(ctx)
}
