package ui

type Card struct {
	ID      int64
	Title   string
	Content string
}

type Project struct {
	ID         int64
	Title      string
	Backlog    []Card
	InProgress []Card
	CodeReview []Card
	Testing    []Card
	Done       []Card
}
