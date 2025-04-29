CREATE TABLE habit_logs (
                            id SERIAL PRIMARY KEY,
                            habit_id INTEGER NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
                            date DATE NOT NULL,
                            created_at TIMESTAMP DEFAULT now(),
                            updated_at TIMESTAMP DEFAULT now(),
                            deleted_at TIMESTAMP
);
