package campaign

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/sedit95/gowit-mini-ad-platform/backend/internal/errors"
	internalHttp "github.com/sedit95/gowit-mini-ad-platform/backend/internal/http"
)

// Handler handles HTTP requests for campaigns.
type Handler struct {
	service *Service
}

// NewHandler creates a new Campaign Handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// handleError maps domain errors to HTTP responses.
func handleError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		switch appErr.Code {
		case errors.CodeValidationError, errors.CodeInvalidStatus:
			internalHttp.WriteError(w, http.StatusBadRequest, appErr.Code, appErr.Message)
		case errors.CodeNotFound:
			internalHttp.WriteError(w, http.StatusNotFound, appErr.Code, appErr.Message)
		default:
			internalHttp.WriteError(w, http.StatusInternalServerError, errors.CodeInternalError, "An internal error occurred")
		}
		return
	}
	internalHttp.WriteError(w, http.StatusInternalServerError, errors.CodeInternalError, "An internal error occurred")
}

// parseUUID helper extracts and validates the UUID from the URL parameter.
func parseUUID(r *http.Request) (uuid.UUID, error) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, errors.Validation("invalid campaign ID format")
	}
	return id, nil
}

// ListCampaigns handles GET /campaigns
func (h *Handler) ListCampaigns(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.service.List(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusOK, campaigns)
}

// CreateCampaign handles POST /campaigns
func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req CreateCampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, errors.Validation("invalid request body"))
		return
	}

	res, err := h.service.Create(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusCreated, res)
}

// GetCampaign handles GET /campaigns/{id}
func (h *Handler) GetCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	res, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusOK, res)
}

// UpdateCampaign handles PUT /campaigns/{id}
func (h *Handler) UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	var req UpdateCampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, errors.Validation("invalid request body"))
		return
	}

	res, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusOK, res)
}

// DeleteCampaign handles DELETE /campaigns/{id}
func (h *Handler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := h.service.SoftDelete(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusNoContent, nil)
}

// RecordImpression handles POST /impression/{id}
func (h *Handler) RecordImpression(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	res, err := h.service.RecordImpression(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusOK, res)
}

// GetStats handles GET /stats/{id}
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	res, err := h.service.GetStats(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	internalHttp.WriteJSON(w, http.StatusOK, res)
}
