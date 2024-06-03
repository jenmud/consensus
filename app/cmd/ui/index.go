package ui

import (
	"context"
	"embed"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

//go:embed templates/*.tmpl
var embedded embed.FS

//go:embed static/*.js
var static embed.FS

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func project(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl", "templates/project.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	project := Project{
		ID:    id,
		Title: "Project 1",
	}

	if err := tmpl.Execute(w, project); err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItems(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{
		{
			Title:   "Ticket 1",
			Content: "Ticket 1 content",
		},
		{
			Title:   "Ticket 2",
			Content: "Ticket 2 content",
		},
		{
			Title:   "Ticket 3",
			Content: "Ticket 3 content",
		},
		{
			Title:   "Ticket 4",
			Content: "Ticket Project 4 content",
		},
	}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItemsBacklog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{
		{
			Title:   "Ticket 3",
			Content: "Ticket 3 content",
		},
		{
			Title:   "Ticket 4",
			Content: "Ticket Project 4 content",
		},
	}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItemsInProgress(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{
		{
			Title:   "Ticket 2",
			Content: "Ticket 2 content",
		},
	}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItemsCodeReview(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{
		{
			Title:   "Ticket 1",
			Content: "Ticket 1 content",
		},
	}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItemsTesting(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func projectItemsDone(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/project-items.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := []Card{}

	if err := tmpl.Execute(w, items); err != nil {
		slog.Error("Failed to render index page", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// registerRoutes registers the routes for the HTTP server.
func registerRoutes(mux *chi.Mux) {
	mux.Get("/", index)

	mux.Route("/projects/{id:^[1-9]+}", func(r chi.Router) {
		r.Get("/", project)
		r.Route("/items", func(r chi.Router) {
			r.Get("/", projectItems)
			r.HandleFunc("/backlog", projectItemsBacklog)
			r.HandleFunc("/inprogress", projectItemsInProgress)
			r.HandleFunc("/codereview", projectItemsCodeReview)
			r.HandleFunc("/testing", projectItemsTesting)
			r.HandleFunc("/done", projectItemsDone)
		})
	})

	staticFS := http.FS(static)
	fileServer := http.FileServer(staticFS)
	mux.Handle("/static/*", http.StripPrefix("/", fileServer))
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
	mux := chi.NewRouter()
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
