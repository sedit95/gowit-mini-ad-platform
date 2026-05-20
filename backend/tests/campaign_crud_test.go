package tests

import (
	"context"
	"testing"
	"time"

	"github.com/sedit95/gowit-mini-ad-platform/backend/internal/campaign"
)

func TestCampaignCRUD(t *testing.T) {
	svc, pool := setupTestService(t)
	defer pool.Close()

	ctx := context.Background()

	// 1. Create Campaign
	req := validCreateRequest()
	req.Budget = 50
	c, err := svc.Create(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create campaign: %v", err)
	}
	if c.Title != req.Title || c.InitialBudget != 50 || c.RemainingBudget != 50 {
		t.Fatalf("Created campaign values are incorrect: %+v", c)
	}

	// 2. List Campaigns
	list, err := svc.List(ctx)
	if err != nil {
		t.Fatalf("Failed to list campaigns: %v", err)
	}
	if len(list) != 1 || list[0].ID != c.ID {
		t.Fatalf("List did not return the expected campaign")
	}

	// 3. Get Campaign By ID
	got, err := svc.GetByID(ctx, c.ID)
	if err != nil {
		t.Fatalf("Failed to get campaign by ID: %v", err)
	}
	if got.ID != c.ID {
		t.Fatalf("Retrieved wrong campaign")
	}

	// 4. Update Campaign (Allowed fields)
	updateReq := campaign.UpdateCampaignRequest{
		Title:     "Updated Campaign",
		Currency:  "EUR",
		StartDate: req.StartDate,
		EndDate:   req.EndDate.Add(48 * time.Hour),
		Status:    "paused",
	}
	updated, err := svc.Update(ctx, c.ID, updateReq)
	if err != nil {
		t.Fatalf("Failed to update campaign: %v", err)
	}
	
	// Ensure allowed fields changed
	if updated.Title != "Updated Campaign" || updated.Currency != "EUR" || updated.Status != campaign.StatusPaused {
		t.Fatalf("Campaign was not updated correctly: %+v", updated)
	}
	
	// Ensure budget fields remain unchanged
	if updated.InitialBudget != 50 || updated.RemainingBudget != 50 {
		t.Fatalf("Campaign budget was illegally modified during update: %+v", updated)
	}
}
