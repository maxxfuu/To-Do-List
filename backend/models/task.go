package models

type Task struct {
	TaskID  int
	UserID  int
	Title   string
	Content string
	Status  string
}
