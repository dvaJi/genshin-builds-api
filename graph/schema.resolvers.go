package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"src/genshindata"

	"github.com/dvaJi/genshin-builds-api/graph/generated"
	"github.com/dvaJi/genshin-builds-api/graph/model"
)

func (r *queryResolver) Artifacts(ctx context.Context, lang string) ([]*model.Artifact, error) {
	return genshindata.GetArtifacts(r.DB, lang)
}

func (r *queryResolver) Artifact(ctx context.Context, id string) (*model.Artifact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CommonMaterials(ctx context.Context) ([]*model.CommonMaterial, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CommonMaterial(ctx context.Context, id string) (*model.CommonMaterial, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
