package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// InitRouter initialize a new chi router instance.
func (s *App) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		// Basic set of handler routes
		r.Get("/status", s.statusHandler)
		r.Post("/buy_pizza", s.orderPizzaHandler)
		r.Put("/order", s.updateOrderHandler)
		r.Delete("order/{order_id}", s.deleteStatusHandler)
		r.Get("/order-status/{user_id}", s.orderStatusHandler)
		r.Put("/new_pizza", s.newPizzaHandler)
		r.Get("show_details", s.showPizzaDetailsHandler)
		r.Get("/list_orders", s.ListAllOrdersHandler)

	})

	return r
}
