package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"setuBack/graph/generated"
	"setuBack/graph/model"
)

func (r *mutationResolver) AddUser(ctx context.Context, input []*model.AddUserInput) (*model.AddUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, filter model.UserFilter) (*model.DeleteUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddProject(ctx context.Context, input []*model.AddProjectInput) (*model.AddProjectPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectInput) (*model.UpdateProjectPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProject(ctx context.Context, filter model.ProjectFilter) (*model.DeleteProjectPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddTicket(ctx context.Context, input []*model.AddTicketInput) (*model.AddTicketPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTicket(ctx context.Context, input model.UpdateTicketInput) (*model.UpdateTicketPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTicket(ctx context.Context, filter model.TicketFilter) (*model.DeleteTicketPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddComment(ctx context.Context, input []*model.AddCommentInput) (*model.AddCommentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, input model.UpdateCommentInput) (*model.UpdateCommentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, filter model.CommentFilter) (*model.DeleteCommentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {

	var pass = "Flutter"
	dummyUser := model.User{
		ID:       "1",
		Username: "Dummy",
		Position: &pass,
		Password: "password",
	}

	return &dummyUser, nil
}

func (r *queryResolver) QueryUser(ctx context.Context, filter *model.UserFilter, order *model.UserOrder, first *int, offset *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateUser(ctx context.Context, filter *model.UserFilter) (*model.UserAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProject(ctx context.Context, id string) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryProject(ctx context.Context, filter *model.ProjectFilter, order *model.ProjectOrder, first *int, offset *int) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateProject(ctx context.Context, filter *model.ProjectFilter) (*model.ProjectAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTicket(ctx context.Context, id string) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryTicket(ctx context.Context, filter *model.TicketFilter, order *model.TicketOrder, first *int, offset *int) ([]*model.Ticket, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateTicket(ctx context.Context, filter *model.TicketFilter) (*model.TicketAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryComment(ctx context.Context, filter *model.CommentFilter, order *model.CommentOrder, first *int, offset *int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateComment(ctx context.Context, filter *model.CommentFilter) (*model.CommentAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
