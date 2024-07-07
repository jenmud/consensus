package ui

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jenmud/consensus/business/service"
	"github.com/jenmud/consensus/foundation/crypto"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// defaultJWTExpiry is the default expiry time for JWT tokens.
const defaultJWTExpiry = 15 * time.Minute

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
	tmpl, err := template.ParseFS(embedded, "templates/index.tmpl", "templates/nav.tmpl", "templates/swimlanes.tmpl")
	if err != nil {
		slog.Error("Failed to render index page", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If we get here we have a valid user, so generate the JWT token
	token, err := JWTFromCtx(r.Context())
	if err != nil {
		slog.Error("Failed to generate JWT", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := Payload{
		Project: Project{
			ID:          1,
			Title:       "SomeFakeProject",
			Description: "Some fake project for UI testing",
			Backlog: []Card{
				{
					ID:      1,
					Title:   "SomeFakeCard",
					Content: "Some fake card for UI testing",
				},
				{
					ID:      2,
					Title:   "SomeFakeCard-2",
					Content: "Second fake card for UI testing",
				},
			},
			InProgress: []Card{
				{
					ID:      3,
					Title:   "SomeFakeCard-3",
					Content: "Third fake card for UI testing",
				},
			},
			CodeReview: []Card{},
			Testing: []Card{
				{
					ID:      4,
					Title:   "SomeFakeCard-4",
					Content: "Fourth fake card for UI testing",
				},
				{
					ID:      5,
					Title:   "SomeFakeCard-5",
					Content: "Fifth fake card for UI testing",
				},
			},
			Done: []Card{
				{
					ID:      6,
					Title:   "SomeFakeCard-6",
					Content: "Sixth fake card for UI testing",
				},
				{
					ID:      7,
					Title:   "SomeFakeCard-7",
					Content: "Seventh fake card for UI testing",
				},
				{
					ID:      8,
					Title:   "SomeFakeCard-8",
					Content: "Eighth fake card for UI testing",
				},
			},
			Owner: User{},
		},
		JWT: token,
	}

	if err := tmpl.Execute(w, payload); err != nil {
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

// registerUserPOST creates a new user.
func registerUserPOST(w http.ResponseWriter, r *http.Request) {
	client, ok := r.Context().Value(serviceCtx).(service.ConsensusClient)
	if !ok {
		slog.Error("failed to get client from context")
		http.Error(w, "failed to get consensus service client", http.StatusInternalServerError)
		return
	}

	password := r.FormValue("password")
	confirm := r.FormValue("confirm_password")

	if password != confirm {
		slog.Error("failed to confirm the password")
		http.Error(w, "failed to confirm the password", http.StatusUnauthorized)
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

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// login logs the account in and returns the JWT token.
func login(w http.ResponseWriter, r *http.Request) {
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

func loginFormPOST(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// If we get here we have a valid user, so generate the JWT token
	token := JWT{
		ID:        uuid.NewString(),
		Audience:  []string{"Consensus"},
		Subject:   resp.Id,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(defaultJWTExpiry),
		User: User{
			ID:        resp.Id,
			FirstName: resp.FirstName,
			LastName:  resp.LastName,
			Email:     email,
			Role:      resp.Role.String(),
		},
	}

	_, tokenAuth, err := tokenAuth.Encode(token.AsMap())
	if err != nil {
		slog.Error("Failed to generate JWT token", slog.String("reason", err.Error()))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenAuth))

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt", // must be "jwt" to be searchable by the jwtauth.Varifier
		Value:    tokenAuth,
		Expires:  time.Now().Add(7 * 24 * time.Hour), // 7 days
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logoutFormPOST(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt", // must be "jwt" to be searchable by the jwtauth.Varifier
		Expires: time.Now(),
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
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
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			slog.Error("failed to get token from context", slog.String("reason", err.Error()))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		next.ServeHTTP(w, r)
	})
}

// registerRoutes registers the routes for the HTTP server.
func registerRoutes(mux *chi.Mux) {

	// PUBLIC ROUTES
	mux.Get("/login", login)
	mux.Post("/login", loginFormPOST)
	mux.Get("/register", registerUserForm)
	mux.Post("/register", registerUserPOST)
	mux.Post("/logout", logoutFormPOST)

	//mux.Get("/", index) // TOOD: Remove me
	// PROTECTED ROUTES
	mux.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(UnloggedInRedirector)
		r.Get("/", index)
		// TODO: complete the following
		// r.Get("/projects", projects)                  // all projects view
		// r.Post("/projects", projectsPOST)             // add or remove a project
		// r.Get("/projects/{id}/swimlanes", swimlanes)  // project swimlanes
		// r.Post("/projects/{id}/backlog", backlogPOST) // add or remove task to the backlog
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
