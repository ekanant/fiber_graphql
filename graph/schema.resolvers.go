package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fiber_graphql/graph/generated"
	"fiber_graphql/graph/model"
	"fiber_graphql/todo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	todoToAdd := todo.Todo{
		Text: input.Text,
		Done: false,
	}

	newTodo := todo.AddTodo(todoToAdd)

	todoGraphModel := &model.Todo{
		ID:   newTodo.ID,
		Text: newTodo.Text,
		Done: newTodo.Done,
	}

	return todoGraphModel, nil

}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todoList := todo.GetAllTodo()
	var todoGraphs []*model.Todo
	for _, v := range todoList {
		todoGraphs = append(todoGraphs, &model.Todo{
			ID:   v.ID,
			Text: v.Text,
			Done: v.Done,
		})
	}
	return todoGraphs, nil
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	text := "Hello world"
	return &text, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
