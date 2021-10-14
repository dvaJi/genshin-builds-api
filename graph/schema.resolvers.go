package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"src/genshindata"

	"github.com/dvaJi/genshin-builds-api/graph/generated"
	"github.com/dvaJi/genshin-builds-api/graph/model"
)

func (r *queryResolver) Artifacts(ctx context.Context, lang string) ([]*model.Artifact, error) {
	return genshindata.GetArtifacts(r.DB, lang)
}

func (r *queryResolver) Artifact(ctx context.Context, lang string, id string) (*model.Artifact, error) {
	return genshindata.GetArtifactById(r.DB, lang, id)
}

func (r *queryResolver) Characters(ctx context.Context, lang string) ([]*model.Character, error) {
	return genshindata.GetCharacters(r.DB, lang)
}

func (r *queryResolver) Character(ctx context.Context, lang string, id string) (*model.Character, error) {
	return genshindata.GetCharacterById(r.DB, lang, id)
}

func (r *queryResolver) CommonMaterials(ctx context.Context, lang string) ([]*model.CommonMaterial, error) {
	return genshindata.GetCommonMaterials(r.DB, lang)
}

func (r *queryResolver) CommonMaterial(ctx context.Context, lang string, id string) (*model.CommonMaterial, error) {
	return genshindata.GetCommonMaterialById(r.DB, lang, id)
}

func (r *queryResolver) ElementalStoneMaterials(ctx context.Context, lang string) ([]*model.ElementalStoneMaterial, error) {
	return genshindata.GetElementalStoneMaterials(r.DB, lang)
}

func (r *queryResolver) ElementalStoneMaterial(ctx context.Context, lang string, id string) (*model.ElementalStoneMaterial, error) {
	return genshindata.GetElementalStoneMaterialById(r.DB, lang, id)
}

func (r *queryResolver) Fish(ctx context.Context, lang string) ([]*model.Fish, error) {
	return genshindata.GetFish(r.DB, lang)
}

func (r *queryResolver) FishingRods(ctx context.Context, lang string) ([]*model.FishingRod, error) {
	return genshindata.GetFishingRods(r.DB, lang)
}

func (r *queryResolver) Baits(ctx context.Context, lang string) ([]*model.Bait, error) {
	return genshindata.GetBaits(r.DB, lang)
}

func (r *queryResolver) Food(ctx context.Context, lang string) ([]*model.Food, error) {
	return genshindata.GetFood(r.DB, lang)
}

func (r *queryResolver) Ingredients(ctx context.Context, lang string) ([]*model.Ingredient, error) {
	return genshindata.GetIngredients(r.DB, lang)
}

func (r *queryResolver) JewelMaterials(ctx context.Context, lang string) ([]*model.JewelMaterial, error) {
	return genshindata.GetJewelMaterials(r.DB, lang)
}

func (r *queryResolver) LocalMaterials(ctx context.Context, lang string) ([]*model.LocalMaterial, error) {
	return genshindata.GetLocalMaterials(r.DB, lang)
}

func (r *queryResolver) Potions(ctx context.Context, lang string) ([]*model.Potion, error) {
	return genshindata.GetPotions(r.DB, lang)
}

func (r *queryResolver) TalentLvlUpMaterials(ctx context.Context, lang string) ([]*model.TalentLvlUpMaterial, error) {
	return genshindata.GetTalentLvlUpMaterials(r.DB, lang)
}

func (r *queryResolver) WeaponPrimaryMaterials(ctx context.Context, lang string) ([]*model.WeaponPrimaryMaterial, error) {
	return genshindata.GetWeaponPrimaryMaterials(r.DB, lang)
}

func (r *queryResolver) WeaponSecondaryMaterials(ctx context.Context, lang string) ([]*model.WeaponSecondaryMaterial, error) {
	return genshindata.GetWeaponSecondaryMaterials(r.DB, lang)
}

func (r *queryResolver) Weapons(ctx context.Context, lang string) ([]*model.Weapon, error) {
	return genshindata.GetWeapons(r.DB, lang)
}

func (r *queryResolver) Weapon(ctx context.Context, lang string, id string) (*model.Weapon, error) {
	return genshindata.GetWeaponById(r.DB, lang, id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
