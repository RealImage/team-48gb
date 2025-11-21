package router

import (
	"net/http"

	"github.com/RealImage/team-48gb/internal/handler"
	"github.com/RealImage/team-48gb/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates and configures a new chi router
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize services
	campaignService := services.NewCampaignService()

	// Initialize handlers
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// Routes
	r.Route("/v1/campaign", func(r chi.Router) {
		r.Post("/", campaignHandler.CreateCampaign)
	})

	return r
}
