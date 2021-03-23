package todo

import gonanoid "github.com/matoous/go-nanoid/v2"

type Todo struct {
	ID     string
	Text   string
	Done   bool
	IDUser string
}

var todoDatas = []*Todo{
	&Todo{ID: "1", Done: false, Text: "Coding", IDUser: "1"},
	&Todo{ID: "2", Done: true, Text: "Monitor database", IDUser: "1"},
}

func newTodoID() string {
	id, _ := gonanoid.New()
	return id
}

func AddTodo(todo Todo) *Todo {
	todo.ID = newTodoID()
	todoDatas = append(todoDatas, &todo)
	return &todo
}

func GetAllTodo() []*Todo {
	return todoDatas
}
