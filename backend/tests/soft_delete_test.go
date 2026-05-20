package tests

import (
	"context"
	"testing"
)

func TestSoftDelete(t *testing.T) {
	svc, pool := setupTestService(t)
	defer pool.Close()

	ctx := context.Background()

	// Create campaign
	req := validCreateRequest()
	c, err := svc.Create(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create campaign: %v", err)
	}

	// SoftDelete
	err = svc.SoftDelete(ctx, c.ID)
	if err != nil {
		t.Fatalf("Failed to soft delete campaign: %v", err)
	}

	// Assert GetByID returns not_found mapping
	_, err = svc.GetByID(ctx, c.ID)
	if err == nil {
		t.Fatalf("Expected error for GetByID on deleted campaign, got nil")
	}

	// Assert GetStats returns not_found mapping
	_, err = svc.GetStats(ctx, c.ID)
	if err == nil {
		t.Fatalf("Expected error for GetStats on deleted campaign, got nil")
	}

	// Assert RecordImpression behaves correctly (returns error mapping to not found for soft deleted)
	_, err = svc.RecordImpression(ctx, c.ID)
	if err == nil {
		t.Fatalf("Expected error for RecordImpression on deleted campaign, got nil")
	}
}
