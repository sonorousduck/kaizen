CREATE TABLE goal_entries (
    id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    goal_id                     UUID NOT NULL REFERENCES goals(id) ON DELETE CASCADE,
    value                       NUMERIC NOT NULL,
    note                        TEXT,
    date                        TIMESTAMPTZ NOT NULL,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_goal_entries_goal_idx ON goal_entries (goal_id, date);
CREATE INDEX idx_goal_entries_date ON goal_entries (date);

