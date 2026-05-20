package campaign

import (
	"time"

	"github.com/google/uuid"
)

// CampaignStatus defines the valid states for a campaign.
type CampaignStatus string

const (
	StatusActive    CampaignStatus = "active"
	StatusPaused    CampaignStatus = "paused"
	StatusCompleted CampaignStatus = "completed"
)

// Campaign represents the core domain model for an advertising campaign.
type Campaign struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	Title           string         `json:"title" db:"title"`
	Currency        string         `json:"currency" db:"currency"`
	InitialBudget   int            `json:"initial_budget" db:"initial_budget"`
	RemainingBudget int            `json:"remaining_budget" db:"remaining_budget"`
	ImpressionCount int            `json:"impression_count" db:"impression_count"`
	Status          CampaignStatus `json:"status" db:"status"`
	StartDate       time.Time      `json:"start_date" db:"start_date"`
	EndDate         time.Time      `json:"end_date" db:"end_date"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at" db:"deleted_at"`
}
