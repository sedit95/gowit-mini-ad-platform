package campaign

import (
	"time"

	"github.com/google/uuid"
)

// CreateCampaignRequest is the incoming payload to create a new campaign.
// Note: It uses the general field "budget".
type CreateCampaignRequest struct {
	Title     string    `json:"title"`
	Budget    int       `json:"budget"`
	Currency  string    `json:"currency"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status,omitempty"`
}

// UpdateCampaignRequest is the incoming payload to update a campaign.
// Note: Budget is explicitly excluded from MVP update capabilities.
type UpdateCampaignRequest struct {
	Title     string    `json:"title"`
	Currency  string    `json:"currency"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}

// CampaignResponse is the outgoing payload representing a campaign's state.
type CampaignResponse struct {
	ID              uuid.UUID      `json:"id"`
	Title           string         `json:"title"`
	Currency        string         `json:"currency"`
	InitialBudget   int            `json:"initial_budget"`
	RemainingBudget int            `json:"remaining_budget"`
	ImpressionCount int            `json:"impression_count"`
	Status          CampaignStatus `json:"status"`
	StartDate       time.Time      `json:"start_date"`
	EndDate         time.Time      `json:"end_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ImpressionResponse is the outgoing payload confirming an impression attempt.
type ImpressionResponse struct {
	CampaignID      uuid.UUID      `json:"campaign_id"`
	Accepted        bool           `json:"accepted"`
	Reason          string         `json:"reason,omitempty"`
	RemainingBudget int            `json:"remaining_budget"`
	ImpressionCount int            `json:"impression_count,omitempty"`
	Status          CampaignStatus `json:"status"`
}

// StatsResponse is the outgoing payload containing campaign statistics.
type StatsResponse struct {
	CampaignID       uuid.UUID      `json:"campaign_id"`
	Title            string         `json:"title"`
	Currency         string         `json:"currency"`
	TotalImpressions int            `json:"total_impressions"`
	InitialBudget    int            `json:"initial_budget"`
	SpentBudget      int            `json:"spent_budget"`
	RemainingBudget  int            `json:"remaining_budget"`
	Status           CampaignStatus `json:"status"`
}
