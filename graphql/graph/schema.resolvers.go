package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math"
	"src/genshindata"

	"github.com/dvaJi/genshin-builds-api/graph/generated"
	"github.com/dvaJi/genshin-builds-api/graph/model"
)

func (r *queryResolver) Artifacts(ctx context.Context, lang string) ([]*model.Artifact, error) {
	return genshindata.GetArtifacts(r.DB, lang)
}

func (r *queryResolver) Artifact(ctx context.Context, lang string, id string) (*model.Artifact, error) {
	art, err := genshindata.GetArtifactById(r.DB, lang, id)
	fmt.Printf("%+v\n", art)
	return art, err
}

func (r *queryResolver) Characters(ctx context.Context, lang string) ([]*model.CharacterInfo, error) {
	return genshindata.GetCharacters(r.DB, lang)
}

func (r *queryResolver) Character(ctx context.Context, lang string, id string) (*model.Character, error) {
	return genshindata.GetCharacterById(r.DB, lang, id)
}

func (r *queryResolver) CharacterExpMaterials(ctx context.Context, lang string) ([]*model.ExpMaterial, error) {
	return genshindata.GetCharacterExpMaterials(r.DB, lang)
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

func (r *queryResolver) WeaponExpMaterials(ctx context.Context, lang string) ([]*model.ExpMaterial, error) {
	return genshindata.GetWeaponExpMaterials(r.DB, lang)
}

func (r *queryResolver) Weapons(ctx context.Context, lang string) ([]*model.WeaponInfo, error) {
	return genshindata.GetWeapons(r.DB, lang)
}

func (r *queryResolver) Weapon(ctx context.Context, lang string, id string) (*model.Weapon, error) {
	return genshindata.GetWeaponById(r.DB, lang, id)
}

func (r *queryResolver) CalculateCharacterLevel(ctx context.Context, characterID string, lang string, params model.CalculateCharacterParams) (*model.CalculationCharacterResult, error) {
	lvlExp := genshindata.CharacterLvlExpList()
	character, err := genshindata.GetCharacterById(r.DB, lang, characterID)

	if err != nil {
		return nil, err
	}

	charExpMaterial, err := genshindata.GetCharacterExpMaterials(r.DB, lang)

	if err != nil {
		return nil, err
	}

	var current float64 = 0
	var moraNeeded = 0
	items := []*model.CalculationItemResult{}

	// Calculate EXP
	// TODO: This should calculate based on ascension, you will lose exp on every ascension level
	if params.IntendedLevel.Lvl > params.CurrentLevel.Lvl {
		var MinIndex = params.IntendedLevel.Lvl - 1
		var MaxIndex = params.CurrentLevel.Lvl - 1
		target := float64(lvlExp[MinIndex] - (lvlExp[MaxIndex] + 0))
		current = target
		var lvlCost float64 = 1000
		moraNeeded = int((math.Floor(target/lvlCost) * lvlCost) / 5)

		// calculate exp materials
		for _, expItem := range charExpMaterial {
			if expItem.ID == "wanderers_advice" {
				var amount = math.Ceil(current / float64(expItem.Exp))

				items = append(items, &model.CalculationItemResult{
					ID:     expItem.ID,
					Name:   expItem.Name,
					Img:    "/materials/" + expItem.ID + ".png",
					Amount: int(amount),
				})

				current = target - math.Ceil(target/float64(expItem.Exp))*float64(expItem.Exp)
			} else if current > 0 && math.Floor(current/float64(expItem.Exp)) > 0 {
				var amount = math.Floor(current / float64(expItem.Exp))
				items = append(items, &model.CalculationItemResult{
					ID:     expItem.ID,
					Name:   expItem.Name,
					Img:    "/materials/" + expItem.ID + ".png",
					Amount: int(amount),
				})

				current = target - math.Floor(target/float64(expItem.Exp))*float64(expItem.Exp)
			}
		}
	}

	var ItemsMap = make(map[string]*model.CalculationItemResult)

	// Calculate Ascension materials
	if params.CurrentLevel.AsclLvl < params.IntendedLevel.AsclLvl {
		for _, item := range character.Ascension {
			if *item.Ascension <= params.CurrentLevel.AsclLvl || *item.Ascension > params.IntendedLevel.AsclLvl {
				continue
			}

			moraNeeded += *item.Cost
			// Mat one
			if ItemsMap[item.MatOne.ID] == nil {
				ItemsMap[item.MatOne.ID] = &model.CalculationItemResult{
					ID:     item.MatOne.ID,
					Name:   item.MatOne.Name,
					Img:    "/jewels_materials/" + item.MatOne.ID + ".png",
					Amount: *item.MatOne.Amount,
				}
			} else {
				ItemsMap[item.MatOne.ID].Amount += *item.MatOne.Amount
			}

			// Mat two
			if item.MatTwo != nil {
				if ItemsMap[item.MatTwo.ID] == nil {
					ItemsMap[item.MatTwo.ID] = &model.CalculationItemResult{
						ID:     item.MatTwo.ID,
						Name:   item.MatTwo.Name,
						Img:    "/elemental_stone_materials/" + item.MatTwo.ID + ".png",
						Amount: *item.MatTwo.Amount,
					}
				} else {
					ItemsMap[item.MatTwo.ID].Amount += *item.MatTwo.Amount
				}
			}

			// Mat three
			if ItemsMap[item.MatThree.ID] == nil {
				ItemsMap[item.MatThree.ID] = &model.CalculationItemResult{
					ID:     item.MatThree.ID,
					Name:   item.MatThree.Name,
					Img:    "/local_materials/" + item.MatThree.ID + ".png",
					Amount: *item.MatThree.Amount,
				}
			} else {
				ItemsMap[item.MatThree.ID].Amount += *item.MatThree.Amount
			}

			// Mat four
			if ItemsMap[item.MatFour.ID] == nil {
				ItemsMap[item.MatFour.ID] = &model.CalculationItemResult{
					ID:     item.MatFour.ID,
					Name:   item.MatFour.Name,
					Img:    "/common_materials/" + item.MatFour.ID + ".png",
					Amount: *item.MatFour.Amount,
				}
			} else {
				ItemsMap[item.MatFour.ID].Amount += *item.MatFour.Amount
			}
		}
	}

	var TalentsMaterialFolder []string = []string{
		"talent_lvl_up_materials",
		"common_materials",
		"talent_lvl_up_materials",
		"talent_lvl_up_materials",
	}

	// Calculate materials for talents
	CalculateTalentMaterials := func(levelMin int, levelMax int) {
		for _, talent := range character.TalentMaterials {
			if (levelMin <= *talent.Level) && (*talent.Level <= levelMax) {
				moraNeeded += *talent.Cost

				for index, item := range talent.Items {
					var currentFolder = TalentsMaterialFolder[index]
					if ItemsMap[item.ID] == nil {
						ItemsMap[item.ID] = &model.CalculationItemResult{
							ID:     item.ID,
							Name:   item.Name,
							Img:    "/" + currentFolder + "/" + item.ID + ".png",
							Amount: *item.Amount,
						}
					} else {
						ItemsMap[item.ID].Amount += *item.Amount
					}
				}
			}
		}
	}

	// Auto attack
	if params.CurrentTalentLvl.Aa < params.IntendedTalentLvl.Aa {
		CalculateTalentMaterials(params.CurrentTalentLvl.Aa, params.IntendedTalentLvl.Aa)
	}

	// Skill
	if params.CurrentTalentLvl.Skill < params.IntendedTalentLvl.Skill {
		CalculateTalentMaterials(params.CurrentTalentLvl.Skill, params.IntendedTalentLvl.Skill)
	}

	// Burst
	if params.CurrentTalentLvl.Burst < params.IntendedTalentLvl.Burst {
		CalculateTalentMaterials(params.CurrentTalentLvl.Burst, params.IntendedTalentLvl.Burst)
	}

	// append ItemsMap to items
	for _, item := range ItemsMap {
		items = append(items, item)
	}

	if moraNeeded > 0 {
		items = append(items, &model.CalculationItemResult{
			ID:     "mora",
			Name:   "Mora",
			Img:    `/materials/mora.png`,
			Amount: int(moraNeeded),
		})
	}

	expWasted := int(current)

	result := model.CalculationCharacterResult{
		ExpWasted: &expWasted,
		Items:     items,
	}

	return &result, nil
}

func (r *queryResolver) CalculateWeaponLevel(ctx context.Context, lang string, weaponID string, params model.CalculateWeaponParams) ([]*model.CalculationItemResult, error) {
	lvlExp := genshindata.WeaponLvlExpList()
	weapon, err := genshindata.GetWeaponById(r.DB, lang, weaponID)

	if err != nil {
		return nil, err
	}

	weaponExpMaterial, err := genshindata.GetWeaponExpMaterials(r.DB, lang)

	if err != nil {
		return nil, err
	}

	var current float64 = 0
	var moraNeeded = 0
	items := []*model.CalculationItemResult{}

	// Calculate EXP
	if params.IntendedLevel.Lvl >= params.CurrentLevel.Lvl {
		var target = float64(lvlExp[*weapon.Rarity-3][params.IntendedLevel.Lvl-1] -
			(lvlExp[*weapon.Rarity-3][params.CurrentLevel.Lvl-1] + 0))
		current = target
		moraNeeded = int(math.Ceil(target/10/20) * 20)

		for _, expItem := range weaponExpMaterial {
			if expItem.ID == "enhancement_ore" {
				var amount = math.Ceil(current / float64(expItem.Exp))

				items = append(items, &model.CalculationItemResult{
					ID:     expItem.ID,
					Name:   expItem.Name,
					Img:    "/materials/" + expItem.ID + ".png",
					Amount: int(amount),
				})

				current = target - math.Ceil(target/float64(expItem.Exp))*float64(expItem.Exp)
			} else if current > 0 && math.Floor(current/float64(expItem.Exp)) > 0 {
				var amount = math.Floor(current / float64(expItem.Exp))
				items = append(items, &model.CalculationItemResult{
					ID:     expItem.ID,
					Name:   expItem.Name,
					Img:    "/materials/" + expItem.ID + ".png",
					Amount: int(amount),
				})

				current = target - math.Floor(target/float64(expItem.Exp))*float64(expItem.Exp)
			}
		}
	}

	var ItemsMap = make(map[string]*model.CalculationItemResult)
	var TalentsMaterialFolder []string = []string{
		"weapon_primary_materials",
		"weapon_secondary_materials",
		"common_materials",
		"weapon_primary_materials",
	}

	// Calculate materials for talents
	for _, ascension := range weapon.Ascensions {
		if (params.CurrentLevel.AsclLvl <= *ascension.Ascension) && (*ascension.Ascension <= params.IntendedLevel.AsclLvl) {
			if ascension.Cost == nil {
				continue
			}

			moraNeeded += *ascension.Cost

			for index, item := range ascension.Materials {
				var currentFolder = TalentsMaterialFolder[index]
				if ItemsMap[item.ID] == nil {
					ItemsMap[item.ID] = &model.CalculationItemResult{
						ID:     item.ID,
						Name:   item.Name,
						Img:    "/" + currentFolder + "/" + item.ID + ".png",
						Amount: *item.Amount,
					}
				} else {
					ItemsMap[item.ID].Amount += *item.Amount
				}
			}
		}
	}

	// append ItemsMap to items
	for _, item := range ItemsMap {
		items = append(items, item)
	}

	if moraNeeded > 0 {
		items = append(items, &model.CalculationItemResult{
			ID:     "mora",
			Name:   "Mora",
			Img:    `/materials/mora.png`,
			Amount: int(moraNeeded),
		})
	}

	return items, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
