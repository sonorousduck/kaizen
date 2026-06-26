CREATE TABLE goals (
    id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id                     UUID NOT NULL REFERENCES users(id),
    parent_goal_id              UUID REFERENCES goals(id),
    title                       VARCHAR(255) NOT NULL,
    description                 TEXT,
    starting_value              NUMERIC,
    target_value                NUMERIC,
    unit                        VARCHAR(255),
    frequency_interval          INTEGER NOT NULL DEFAULT 1,
    frequency                   VARCHAR(20) NOT NULL CHECK (frequency IN ('once', 'daily', 'weekly', 'monthly', 'yearly')),
    goal_type                   VARCHAR(20) NOT NULL CHECK (goal_type IN ('habit', 'numeric')),

    due_date                    TIMESTAMPTZ,
    deleted_at                  TIMESTAMPTZ,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
    

CREATE INDEX idx_goals_user_parent_goal ON goals (user_id, parent_goal_id);