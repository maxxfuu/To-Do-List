CREATE TABLE tasks (
	TaskID  SERIAL NOT NULL PRIMARY KEY,
    UserID  INT NOT NULL REFERENCES users(userID),
	Title   TEXT NOT NULL,
	Content TEXT,
	Status  TEXT DEFAULT 'pending',
);