package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sedit95/gowit-mini-ad-platform/backend/internal/campaign"
	"github.com/sedit95/gowit-mini-ad-platform/backend/internal/db"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL is missing. Skipping integration tests.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.NewPostgresPool(ctx, dbURL)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	return pool
}

func cleanCampaignsTable(t *testing.T, pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(), "DELETE FROM campaigns")
	if err != nil {
		t.Fatalf("Failed to clean campaigns table: %v", err)
	}
}

func setupTestService(t *testing.T) (*campaign.Service, *pgxpool.Pool) {
	pool := setupTestDB(t)
	cleanCampaignsTable(t, pool)

	repo := campaign.NewRepository(pool)
	service := campaign.NewService(repo)

	return service, pool
}

func validCreateRequest() campaign.CreateCampaignRequest {
	return campaign.CreateCampaignRequest{
		Title:     "Test Campaign",
		Budget:    10,
		Currency:  "USD",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		Status:    "active",
	}
}
