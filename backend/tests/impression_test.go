package tests

import (
	"context"
	"testing"
)

func TestImpressionLifecycle(t *testing.T) {
	svc, pool := setupTestService(t)
	defer pool.Close()

	ctx := context.Background()

	// Create campaign with budget=2
	req := validCreateRequest()
	req.Budget = 2
	c, err := svc.Create(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create campaign: %v", err)
	}

	// 1. Record first impression (accepted)
	res1, err := svc.RecordImpression(ctx, c.ID)
	if err != nil {
		t.Fatalf("First impression failed: %v", err)
	}
	if !res1.Accepted || res1.RemainingBudget != 1 || res1.ImpressionCount != 1 || string(res1.Status) != "active" {
		t.Fatalf("First impression state incorrect: %+v", res1)
	}

	// 2. Record second impression (accepted, triggers pause)
	res2, err := svc.RecordImpression(ctx, c.ID)
	if err != nil {
		t.Fatalf("Second impression failed: %v", err)
	}
	if !res2.Accepted || res2.RemainingBudget != 0 || res2.ImpressionCount != 2 || string(res2.Status) != "paused" {
		t.Fatalf("Second impression state incorrect: %+v", res2)
	}

	// 3. Record third impression (rejected, budget_exhausted)
	res3, err := svc.RecordImpression(ctx, c.ID)
	if err != nil {
		t.Fatalf("Third impression failed unexpectedly with system error: %v", err)
	}
	if res3.Accepted || res3.Reason != "budget_exhausted" || res3.RemainingBudget != 0 || string(res3.Status) != "paused" {
		t.Fatalf("Third impression state incorrect: %+v", res3)
	}

	// 4. Assert stats do not exceed 2 impressions
	stats, err := svc.GetStats(ctx, c.ID)
	if err != nil {
		t.Fatalf("Failed to get stats: %v", err)
	}
	if stats.TotalImpressions != 2 || stats.SpentBudget != 2 || stats.RemainingBudget != 0 {
		t.Fatalf("Final stats incorrect after exhausted attempts: %+v", stats)
	}
}
