package todo

var lastID uint = 0
var todos []TodoItem

type TodoItem struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}