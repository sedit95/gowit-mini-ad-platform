package tests

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

func TestImpressionConcurrency(t *testing.T) {
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

	const concurrentAttempts = 100
	var wg sync.WaitGroup
	wg.Add(concurrentAttempts)

	var (
		acceptedCount   int32
		rejectedCount   int32
		unexpectedError int32
	)

	// Launch 100 concurrent goroutines trying to deduct from a budget of 10
	for i := 0; i < concurrentAttempts; i++ {
		go func() {
			defer wg.Done()

			res, err := svc.RecordImpression(ctx, c.ID)
			if err != nil {
				atomic.AddInt32(&unexpectedError, 1)
				return
			}

			if res.Accepted {
				atomic.AddInt32(&acceptedCount, 1)
			} else {
				atomic.AddInt32(&rejectedCount, 1)
				// Ensure rejected reasons are correct
				if res.Reason != "budget_exhausted" && res.Reason != "campaign_not_active" {
					t.Errorf("Unexpected rejection reason: %s", res.Reason)
				}
			}
		}()
	}

	// Wait for all concurrent attempts to finish
	wg.Wait()

	// Assertions on the goroutine execution counts
	if unexpectedError != 0 {
		t.Errorf("Expected 0 unexpected errors, got %d", unexpectedError)
	}
	if acceptedCount != 10 {
		t.Errorf("Expected exactly 10 accepted impressions, got %d", acceptedCount)
	}
	if rejectedCount != 90 {
		t.Errorf("Expected exactly 90 rejected impressions, got %d", rejectedCount)
	}

	// Final verification of DB state using GetStats as the source of truth
	stats, err := svc.GetStats(ctx, c.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve final stats: %v", err)
	}

	if stats.RemainingBudget != 0 {
		t.Errorf("Expected final remaining_budget to be 0, got %d", stats.RemainingBudget)
	}
	if stats.RemainingBudget < 0 {
		t.Errorf("FATAL: remaining_budget dropped below zero: %d", stats.RemainingBudget)
	}
	if stats.TotalImpressions != 10 {
		t.Errorf("Expected final total_impressions to be 10, got %d", stats.TotalImpressions)
	}
	if stats.SpentBudget != 10 {
		t.Errorf("Expected final spent_budget to be 10, got %d", stats.SpentBudget)
	}
	if string(stats.Status) != "paused" {
		t.Errorf("Expected final status to be paused, got %s", stats.Status)
	}
}
