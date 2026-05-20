package campaign

import (
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up the HTTP routing for campaign endpoints.
func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Get("/campaigns", handler.ListCampaigns)
	r.Post("/campaigns", handler.CreateCampaign)
	r.Get("/campaigns/{id}", handler.GetCampaign)
	r.Put("/campaigns/{id}", handler.UpdateCampaign)
	r.Delete("/campaigns/{id}", handler.DeleteCampaign)
	
	r.Post("/impression/{id}", handler.RecordImpression)
	r.Get("/stats/{id}", handler.GetStats)
}
