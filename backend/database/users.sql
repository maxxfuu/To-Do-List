CREATE TABLE users (
	userID SERIAL PRIMARY KEY, 
	username TEXT NOT NULL UNIQUE, 
	email    TEXT NOT NULL, 
	password TEXT NOT NULL 
); 