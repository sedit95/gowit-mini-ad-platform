package tests

import (
	"context"
	"testing"
)

func TestCampaignStats(t *testing.T) {
	svc, pool := setupTestService(t)
	defer pool.Close()

	ctx := context.Background()

	// Create campaign with budget=10
	req := validCreateRequest()
	req.Budget = 10
	c, err := svc.Create(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create campaign: %v", err)
	}

	// Record 3 impressions
	for i := 0; i < 3; i++ {
		res, err := svc.RecordImpression(ctx, c.ID)
		if err != nil {
			t.Fatalf("Failed to record impression %d: %v", i+1, err)
		}
		if !res.Accepted {
			t.Fatalf("Expected impression %d to be accepted", i+1)
		}
	}

	// GetStats
	stats, err := svc.GetStats(ctx, c.ID)
	if err != nil {
		t.Fatalf("Failed to get stats: %v", err)
	}

	// Assert stats
	if stats.TotalImpressions != 3 {
		t.Errorf("Expected total_impressions to be 3, got %d", stats.TotalImpressions)
	}
	if stats.InitialBudget != 10 {
		t.Errorf("Expected initial_budget to be 10, got %d", stats.InitialBudget)
	}
	if stats.SpentBudget != 3 {
		t.Errorf("Expected spent_budget to be 3, got %d", stats.SpentBudget)
	}
	if stats.RemainingBudget != 7 {
		t.Errorf("Expected remaining_budget to be 7, got %d", stats.RemainingBudget)
	}
	if string(stats.Status) != "active" {
		t.Errorf("Expected status to be active, got %s", stats.Status)
	}
}
