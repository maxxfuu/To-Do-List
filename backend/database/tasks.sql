CREATE TABLE tasks (
    -- UserID  INT NOT NULL REFERENCES users(UserID),

    TaskID  SERIAL NOT NULLPRIMARY KEY,
	Title   TEXT NOT NULL,
	Content TEXT,
	Status  TEXT DEFAULT 'pending',
);