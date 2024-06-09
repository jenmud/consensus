package ui

type Card struct {
	ID      int64
	Title   string
	Content string
}

type Project struct {
	ID          int64
	Title       string
	Description string
	Backlog     []Card
	InProgress  []Card
	CodeReview  []Card
	Testing     []Card
	Done        []Card
	Owner       User
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Role      string
}
