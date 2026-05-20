CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE campaigns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    initial_budget INTEGER NOT NULL,
    remaining_budget INTEGER NOT NULL,
    impression_count INTEGER NOT NULL DEFAULT 0,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT chk_initial_budget CHECK (initial_budget > 0),
    CONSTRAINT chk_remaining_budget CHECK (remaining_budget >= 0),
    CONSTRAINT chk_impression_count CHECK (impression_count >= 0),
    CONSTRAINT chk_status CHECK (status IN ('active', 'paused', 'completed')),
    CONSTRAINT chk_end_date CHECK (end_date >= start_date)
);

CREATE INDEX idx_campaigns_active_listing ON campaigns (created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_campaigns_status ON campaigns (status) WHERE deleted_at IS NULL;
