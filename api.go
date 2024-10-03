package main

import (
	"log"
	"net/http"

	"github.com/buraksaglam089/tool-tube/handlers"
	"github.com/buraksaglam089/tool-tube/services/auth"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

type APIServer struct {
	listenAddr string
	db         *gorm.DB
}

func NewAPIServer(listenAddr string, db *gorm.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Credentials"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
	sessionStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: "secret",
		MaxAge:     60 * 60 * 24 * 7,
		Secure:     false,
		HttpOnly:   true,
	})
	authService := auth.NewAuthService(sessionStore)
	h := handlers.NewHandler(s.db, authService)

	router.Get("/foo", h.HandleFoo)
	router.Post("/user", h.CreateNewUser)

	//Auth
	router.Get("/auth/{provider}", h.HandleProvideLogin)
	router.Get("/api/auth/{provider}/callback", h.HandleAuthCallbackFunction)
	router.Get("/auth/me", h.GetCurrentUser)
	router.Post("/tool/convert", h.ConvertPlaylist)

	log.Printf("Starting server on %s", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
