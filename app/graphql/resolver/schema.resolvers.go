package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/entities"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/graphql/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/graphql/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	fmt.Print("aaaaaa")
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []entities.Todo
	if err := r.DB.Joins("User").Find(&todos).Error; err != nil {
		return nil, err
	}

	var results []*model.Todo
	for _, todo := range todos {
		var newTodo model.Todo
		var newUser model.User
		newTodo.ID = strconv.Itoa(int(todo.BaseModel.ID))
		newTodo.Title = todo.Title
		newTodo.Comment = todo.Comment
		newUser.ID = strconv.Itoa(int(todo.User.BaseModel.ID))
		newUser.Name = todo.User.Name
		newUser.Email = todo.User.Email
		newTodo.User = &newUser
		results = append(results, &newTodo)
	}

	return results, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
