package ui

import (
	"context"
	"embed"
	"encoding/json"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jenmud/consensus/business/service"
)

//go:embed templates/*.tmpl
var embedded embed.FS

//go:embed static/*.js
var static embed.FS

// index renders the index page.
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

// login renders the login page.
func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/login.tmpl")
	if err != nil {
		slog.Error("Failed to render login page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		slog.Error("Failed to render login page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// projects renders the projects page.
func projects(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/projects.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, ok := r.Context().Value(serviceCtx).(service.ConsensusClient)
	if !ok {
		slog.Error("failed to get client from context")
		http.Error(w, "failed to get consensus service client", http.StatusInternalServerError)
		return
	}

	projects, err := client.GetProjects(r.Context(), &service.ProjectsReq{})
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If they are asking for JSON, then return the JSON response.
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(projects); err != nil {
			slog.Error("Failed to encode project", slog.String("reason", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Otherwise, render the HTML response.
	if err := tmpl.Execute(w, CoreProjectsToProjects(projects.GetProjects()...)); err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// users renders the users page.
func users(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/users.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, ok := r.Context().Value(serviceCtx).(service.ConsensusClient)
	if !ok {
		slog.Error("failed to get client from context")
		http.Error(w, "failed to get consensus service client", http.StatusInternalServerError)
		return
	}

	users, err := client.GetUsers(r.Context(), &service.GetUsersReq{})
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If they are asking for JSON, then return the JSON response.
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(users); err != nil {
			slog.Error("Failed to encode project", slog.String("reason", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Otherwise, render the HTML response.
	if err := tmpl.Execute(w, CoreUsersToUsers(users.GetUsers()...)); err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// registerRoutes registers the routes for the HTTP server.
func registerRoutes(mux *chi.Mux) {
	mux.Get("/", index)
	mux.Get("/login", login)

	mux.Route("/users", func(r chi.Router) {
		r.Get("/", users)
		r.Route("/projects/{id:^[1-9]+}", func(r chi.Router) {
			//r.Get("/", projectItems)
			//r.HandleFunc("/backlog", projectItemsBacklog)
			//r.HandleFunc("/inprogress", projectItemsInProgress)
			//r.HandleFunc("/codereview", projectItemsCodeReview)
			//r.HandleFunc("/testing", projectItemsTesting)
			//r.HandleFunc("/done", projectItemsDone)
		})
	})

	mux.Route("/projects", func(r chi.Router) {
		r.Get("/", projects)
		r.Route("/projects/{id:^[1-9]+}", func(r chi.Router) {
			r.Route("/items", func(r chi.Router) {
				//r.Get("/", projectItems)
				//r.HandleFunc("/backlog", projectItemsBacklog)
				//r.HandleFunc("/inprogress", projectItemsInProgress)
				//r.HandleFunc("/codereview", projectItemsCodeReview)
				//r.HandleFunc("/testing", projectItemsTesting)
				//r.HandleFunc("/done", projectItemsDone)
			})
		})
	})

	staticFS := http.FS(static)
	fileServer := http.FileServer(staticFS)
	mux.Handle("/static/*", http.StripPrefix("/", fileServer))
}

type serviceCtxKey string

var serviceCtx = serviceCtxKey("service")

// Run starts the HTTP server and serves the static files and the index page.
// It takes an address as a parameter, which is the address on which the server
// should listen for incoming requests. The address should be in the format
// `host:port`.
//
// Returns an error if there was a problem starting the server.
func ListenAndServe(ctx context.Context, addr string, client service.ConsensusClient, logger *slog.Logger) error {
	if logger == nil {
		logger = slog.Default()
	}

	mux := chi.NewRouter()
	registerRoutes(mux)

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			serviceCtx := context.WithValue(ctx, serviceCtx, client)
			logger.Info("starting ui server")
			return serviceCtx
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			logger.Info("peer connection", slog.String("peer", c.RemoteAddr().String()))
			return ctx
		},
	}

	logger.Info("starting http server and accepting client connections")
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	logger.Info("shutting down http server")
	return server.Shutdown(ctx)
}
