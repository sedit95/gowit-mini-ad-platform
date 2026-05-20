package campaign

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ImpressionRecordResult holds the output of a successful atomic decrement.
type ImpressionRecordResult struct {
	CampaignID      uuid.UUID
	RemainingBudget int
	ImpressionCount int
	Status          CampaignStatus
}

// ImpressionState holds minimal campaign state to diagnose failed decrements.
type ImpressionState struct {
	CampaignID      uuid.UUID
	RemainingBudget int
	Status          CampaignStatus
	DeletedAt       *time.Time
}

// Repository handles database operations for campaigns.
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository creates a new Campaign Repository.
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

// Create inserts a new campaign into the database.
func (r *Repository) Create(ctx context.Context, req CreateCampaignRequest, status CampaignStatus) (*Campaign, error) {
	query := `
		INSERT INTO campaigns (title, currency, initial_budget, remaining_budget, impression_count, status, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, currency, initial_budget, remaining_budget, impression_count, status, start_date, end_date, created_at, updated_at, deleted_at
	`
	
	row := r.pool.QueryRow(ctx, query,
		req.Title,
		req.Currency,
		req.Budget,
		req.Budget,
		0,
		status,
		req.StartDate,
		req.EndDate,
	)

	var c Campaign
	err := row.Scan(
		&c.ID, &c.Title, &c.Currency, &c.InitialBudget, &c.RemainingBudget,
		&c.ImpressionCount, &c.Status, &c.StartDate, &c.EndDate,
		&c.CreatedAt, &c.UpdatedAt, &c.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// List retrieves all active (non-deleted) campaigns.
func (r *Repository) List(ctx context.Context) ([]Campaign, error) {
	query := `
		SELECT id, title, currency, initial_budget, remaining_budget, impression_count, status, start_date, end_date, created_at, updated_at, deleted_at
		FROM campaigns
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var c Campaign
		err := rows.Scan(
			&c.ID, &c.Title, &c.Currency, &c.InitialBudget, &c.RemainingBudget,
			&c.ImpressionCount, &c.Status, &c.StartDate, &c.EndDate,
			&c.CreatedAt, &c.UpdatedAt, &c.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return campaigns, nil
}

// GetByID retrieves a single non-deleted campaign.
func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*Campaign, error) {
	query := `
		SELECT id, title, currency, initial_budget, remaining_budget, impression_count, status, start_date, end_date, created_at, updated_at, deleted_at
		FROM campaigns
		WHERE id = $1 AND deleted_at IS NULL
	`
	var c Campaign
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.Title, &c.Currency, &c.InitialBudget, &c.RemainingBudget,
		&c.ImpressionCount, &c.Status, &c.StartDate, &c.EndDate,
		&c.CreatedAt, &c.UpdatedAt, &c.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Update modifies allowed fields of a non-deleted campaign.
func (r *Repository) Update(ctx context.Context, id uuid.UUID, req UpdateCampaignRequest, status CampaignStatus) (*Campaign, error) {
	query := `
		UPDATE campaigns
		SET title = $1, currency = $2, start_date = $3, end_date = $4, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
		RETURNING id, title, currency, initial_budget, remaining_budget, impression_count, status, start_date, end_date, created_at, updated_at, deleted_at
	`
	var c Campaign
	err := r.pool.QueryRow(ctx, query,
		req.Title, req.Currency, req.StartDate, req.EndDate, status, id,
	).Scan(
		&c.ID, &c.Title, &c.Currency, &c.InitialBudget, &c.RemainingBudget,
		&c.ImpressionCount, &c.Status, &c.StartDate, &c.EndDate,
		&c.CreatedAt, &c.UpdatedAt, &c.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// SoftDelete marks a campaign as deleted.
func (r *Repository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE campaigns
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	commandTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

// GetStats returns calculated statistics for a non-deleted campaign.
func (r *Repository) GetStats(ctx context.Context, id uuid.UUID) (*StatsResponse, error) {
	query := `
		SELECT id, title, currency, impression_count, initial_budget, initial_budget - remaining_budget as spent_budget, remaining_budget, status
		FROM campaigns
		WHERE id = $1 AND deleted_at IS NULL
	`
	var s StatsResponse
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&s.CampaignID, &s.Title, &s.Currency, &s.TotalImpressions,
		&s.InitialBudget, &s.SpentBudget, &s.RemainingBudget, &s.Status,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// RecordImpression performs a safe, atomic conditional update to deduct budget and record an impression.
func (r *Repository) RecordImpression(ctx context.Context, id uuid.UUID) (*ImpressionRecordResult, error) {
	query := `
		UPDATE campaigns
		SET remaining_budget = remaining_budget - 1,
		    impression_count = impression_count + 1,
		    status = CASE WHEN remaining_budget - 1 = 0 THEN 'paused' ELSE status END,
		    updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
		  AND status = 'active'
		  AND remaining_budget > 0
		RETURNING id, remaining_budget, impression_count, status
	`
	var res ImpressionRecordResult
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&res.CampaignID, &res.RemainingBudget, &res.ImpressionCount, &res.Status,
	)
	if err != nil {
		return nil, err // Returns pgx.ErrNoRows if conditions aren't met
	}
	return &res, nil
}

// GetImpressionState retrieves minimal campaign state to diagnose failed impression records.
func (r *Repository) GetImpressionState(ctx context.Context, id uuid.UUID) (*ImpressionState, error) {
	query := `
		SELECT id, remaining_budget, status, deleted_at
		FROM campaigns
		WHERE id = $1
	`
	var state ImpressionState
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&state.CampaignID, &state.RemainingBudget, &state.Status, &state.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &state, nil
}
