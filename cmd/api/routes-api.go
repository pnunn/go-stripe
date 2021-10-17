package main

import (
  "net/http"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
  mux := chi.NewRouter()

  mux.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"http://*", "https://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Accept-Language", "Authorization", "Content-Type", "X-CSRF-Token"},
    AllowCredentials: true,
    MaxAge:           300,
    Debug:            true,
  }))

  mux.Post("/api/payment-intent", app.GetPaymentIntent)

  return mux
}
