package ui

import (
	"context"
	"embed"
	"encoding/json"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jenmud/consensus/business/service"
	"github.com/jenmud/consensus/foundation/crypto"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed templates/*.tmpl
var embedded embed.FS

//go:embed static/*.js
var static embed.FS

// secret is the shared secret for all tokens.
var secret string
var tokenAuth *jwtauth.JWTAuth

func init() {
	secret = crypto.Secret()
	if s := os.Getenv("CONCENSUS_SECRET"); s != "" {
		secret = s
	}

	jwtOpts := []jwt.ValidateOption{
		jwt.WithAcceptableSkew(5 * time.Minute),
	}

	tokenAuth = jwtauth.New("HS256", []byte(secret), nil, jwtOpts...)
}

// index renders the index page.
func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl", "templates/login.tmpl")
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

// registerUserForm is the register user form.
func registerUserForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl", "templates/register-form.tmpl")
	if err != nil {
		slog.Error("Failed to render page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Otherwise, render the HTML response.
	if err := tmpl.Execute(w, nil); err != nil {
		slog.Error("Failed to render page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// registerUser creates a new user.
func registerUser(w http.ResponseWriter, r *http.Request) {
	client, ok := r.Context().Value(serviceCtx).(service.ConsensusClient)
	if !ok {
		slog.Error("failed to get client from context")
		http.Error(w, "failed to get consensus service client", http.StatusInternalServerError)
		return
	}

	user := &service.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	switch r.FormValue("role") {
	case "admin":
		user.Role = service.Role_ADMIN
	case "user":
		user.Role = service.Role_USER
	}

	_, err := client.CreateUser(r.Context(), user)

	if err != nil {
		slog.Error("Failed to authenticate user", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// login logs the account in and returns the JWT token.
func login(w http.ResponseWriter, r *http.Request) {
	client, ok := r.Context().Value(serviceCtx).(service.ConsensusClient)
	if !ok {
		slog.Error("failed to get client from context")
		http.Error(w, "failed to get consensus service client", http.StatusInternalServerError)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	resp, err := client.AuthenticateUser(r.Context(), &service.AuthReq{
		Email:    email,
		Password: password,
	})

	if err != nil {
		slog.Error("Failed to authenticate user", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If we get here we have a valid user, so generate the JWT token

	claims := map[string]any{
		"exp":        time.Now().Add(5 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
		"sub":        resp.Id,
		"user_id":    resp.Id,
		"first_name": resp.FirstName,
		"last_name":  resp.LastName,
		"email":      email,
		"role":       resp.Role.String(),
	}

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		slog.Error("Failed to generate JWT token", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt", // must be "jwt" to be searchable by the jwtauth.Varifier
		Value:    token,
		Expires:  time.Now().Add(7 * 24 * time.Hour), // 7 days
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/users", http.StatusSeeOther) // TODO: testing with the /users endpoint
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

func LoggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())

		if token != nil && jwt.Validate(token) == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

		next.ServeHTTP(w, r)
	})
}

func UnloggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())

		if token == nil || jwt.Validate(token) != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

		next.ServeHTTP(w, r)
	})
}

// registerRoutes registers the routes for the HTTP server.
func registerRoutes(mux *chi.Mux) {

	// PUBLIC ROUTES
	mux.Get("/", index)
	mux.Post("/login", login)
	mux.Get("/register", registerUserForm)
	mux.Post("/register", registerUser)

	// PROTECTED ROUTES

	// create a base router that will be used by all sub-routers which
	// will redirect to the login page if the user is not authenticated.
	mux.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(UnloggedInRedirector)
		//r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/users", func(r chi.Router) {
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

		r.Route("/projects", func(r chi.Router) {
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
