package campaign

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/sedit95/gowit-mini-ad-platform/backend/internal/errors"
)

// Service implements the business logic for campaigns.
type Service struct {
	repo *Repository
}

// NewService creates a new Campaign Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// isValidStatus checks if a status is valid.
func isValidStatus(status CampaignStatus) bool {
	return status == StatusActive || status == StatusPaused || status == StatusCompleted
}

// mapToResponse converts a domain Campaign to a CampaignResponse DTO.
func mapToResponse(c *Campaign) CampaignResponse {
	return CampaignResponse{
		ID:              c.ID,
		Title:           c.Title,
		Currency:        c.Currency,
		InitialBudget:   c.InitialBudget,
		RemainingBudget: c.RemainingBudget,
		ImpressionCount: c.ImpressionCount,
		Status:          c.Status,
		StartDate:       c.StartDate,
		EndDate:         c.EndDate,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

// Create validates input and creates a new campaign.
func (s *Service) Create(ctx context.Context, req CreateCampaignRequest) (*CampaignResponse, error) {
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		return nil, errors.Validation("title is required")
	}
	if req.Budget <= 0 {
		return nil, errors.Validation("budget must be greater than 0")
	}
	req.Currency = strings.TrimSpace(req.Currency)
	if req.Currency == "" {
		return nil, errors.Validation("currency is required")
	}
	if req.StartDate.IsZero() {
		return nil, errors.Validation("start_date is required")
	}
	if req.EndDate.IsZero() {
		return nil, errors.Validation("end_date is required")
	}
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.Validation("end_date must be after or equal to start_date")
	}

	status := StatusActive
	if req.Status != "" {
		parsedStatus := CampaignStatus(req.Status)
		if !isValidStatus(parsedStatus) {
			return nil, errors.Validation("status must be active, paused, or completed")
		}
		status = parsedStatus
	}

	c, err := s.repo.Create(ctx, req, status)
	if err != nil {
		return nil, errors.Internal("failed to create campaign")
	}

	res := mapToResponse(c)
	return &res, nil
}

// List retrieves all non-deleted campaigns.
func (s *Service) List(ctx context.Context) ([]CampaignResponse, error) {
	campaigns, err := s.repo.List(ctx)
	if err != nil {
		return nil, errors.Internal("failed to list campaigns")
	}

	responses := make([]CampaignResponse, 0, len(campaigns))
	for _, c := range campaigns {
		responses = append(responses, mapToResponse(&c))
	}

	return responses, nil
}

// GetByID retrieves a single non-deleted campaign.
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*CampaignResponse, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.NotFound("campaign not found")
		}
		return nil, errors.Internal("failed to get campaign")
	}

	res := mapToResponse(c)
	return &res, nil
}

// Update validates input and updates a campaign's allowed fields.
func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateCampaignRequest) (*CampaignResponse, error) {
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		return nil, errors.Validation("title is required")
	}
	req.Currency = strings.TrimSpace(req.Currency)
	if req.Currency == "" {
		return nil, errors.Validation("currency is required")
	}
	if req.StartDate.IsZero() {
		return nil, errors.Validation("start_date is required")
	}
	if req.EndDate.IsZero() {
		return nil, errors.Validation("end_date is required")
	}
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.Validation("end_date must be after or equal to start_date")
	}

	parsedStatus := CampaignStatus(req.Status)
	if !isValidStatus(parsedStatus) {
		return nil, errors.Validation("status must be active, paused, or completed")
	}

	c, err := s.repo.Update(ctx, id, req, parsedStatus)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.NotFound("campaign not found")
		}
		return nil, errors.Internal("failed to update campaign")
	}

	res := mapToResponse(c)
	return &res, nil
}

// SoftDelete marks a campaign as deleted.
func (s *Service) SoftDelete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.SoftDelete(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.NotFound("campaign not found")
		}
		return errors.Internal("failed to delete campaign")
	}
	return nil
}

// GetStats returns the calculated statistics for a campaign.
func (s *Service) GetStats(ctx context.Context, id uuid.UUID) (*StatsResponse, error) {
	stats, err := s.repo.GetStats(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.NotFound("campaign not found")
		}
		return nil, errors.Internal("failed to get campaign stats")
	}
	return stats, nil
}

// RecordImpression attempts to safely deduct budget and record an impression.
func (s *Service) RecordImpression(ctx context.Context, id uuid.UUID) (*ImpressionResponse, error) {
	// 1. Attempt the atomic update via repository.
	res, err := s.repo.RecordImpression(ctx, id)
	
	if err == nil {
		// Atomic update succeeded.
		return &ImpressionResponse{
			CampaignID:      res.CampaignID,
			Accepted:        true,
			RemainingBudget: res.RemainingBudget,
			ImpressionCount: res.ImpressionCount,
			Status:          res.Status,
		}, nil
	}

	// 2. If it's not a "no rows" error, it's a real DB failure.
	if err != pgx.ErrNoRows {
		return nil, errors.Internal("failed to record impression")
	}

	// 3. The atomic update returned no rows. We read the state to determine the business reason.
	state, stateErr := s.repo.GetImpressionState(ctx, id)
	if stateErr != nil {
		if stateErr == pgx.ErrNoRows {
			return nil, errors.NotFound("campaign not found")
		}
		return nil, errors.Internal("failed to retrieve impression state")
	}

	if state.DeletedAt != nil {
		return nil, errors.NotFound("campaign not found")
	}

	if state.RemainingBudget <= 0 {
		return &ImpressionResponse{
			CampaignID:      state.CampaignID,
			Accepted:        false,
			Reason:          "budget_exhausted",
			RemainingBudget: state.RemainingBudget,
			Status:          state.Status,
		}, nil
	}

	if state.Status != StatusActive {
		return &ImpressionResponse{
			CampaignID:      state.CampaignID,
			Accepted:        false,
			Reason:          "campaign_not_active",
			RemainingBudget: state.RemainingBudget,
			Status:          state.Status,
		}, nil
	}

	// Fallback, though logically unreachable given atomic conditions.
	return nil, errors.Internal("impression failed due to unknown state")
}
