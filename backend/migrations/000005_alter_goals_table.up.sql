ALTER TABLE goals ADD COLUMN category_id UUID REFERENCES goal_categories(id) ON DELETE SET NULL;
ALTER TABLE goals ADD CONSTRAINT check_goals_parent_not_self CHECK (parent_goal_id IS DISTINCT FROM id);

CREATE INDEX idx_goals_category ON goals (user_id, category_id);
