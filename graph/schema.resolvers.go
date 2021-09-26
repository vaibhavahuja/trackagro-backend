package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vaibhavahuja/trackagro-backend/graph/generated"
	"github.com/vaibhavahuja/trackagro-backend/graph/model"
)

func (r *queryResolver) GetItemDetails(ctx context.Context, id int) (*model.ItemDetailsResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllProcesses(ctx context.Context, id int) ([]*model.Process, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDetailedContent(ctx context.Context, processID int) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
