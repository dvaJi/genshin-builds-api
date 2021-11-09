package genshindata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "github.com/dvaJi/genshin-builds-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DBImpl struct {
	DbUserName string
	DbPassword string
	DbHost     string
}

func Init(t *DBImpl) *mongo.Client {
	uriDB := "mongodb://" + t.DbUserName + ":" + t.DbPassword + "@" + t.DbHost + ":27017"
	log.Printf("Connecting to %s", uriDB)
	clientOptions := options.Client().ApplyURI(uriDB)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetArtifacts(db *mongo.Client, language string) ([]*model.Artifact, error) {
	var artifacts []*model.Artifact
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("artifacts").Find(ctx, bson.D{})
	if err != nil {
		return artifacts, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Artifact
		err := cur.Decode(&t)
		if err != nil {
			return artifacts, err
		}

		artifacts = append(artifacts, &t)
	}

	if err := cur.Err(); err != nil {
		return artifacts, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return artifacts, nil
}

func GetArtifactById(db *mongo.Client, language string, id string) (*model.Artifact, error) {
	ctx := context.Background()

	artifact := &model.Artifact{}
	err := db.Database("genshindata_"+language).Collection("artifacts").FindOne(ctx, bson.M{"id": id}).Decode(&artifact)

	return artifact, err
}

func GetCharacters(db *mongo.Client, language string) ([]*model.Character, error) {
	var characters []*model.Character
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("characters").Find(ctx, bson.D{})
	if err != nil {
		return characters, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Character
		err := cur.Decode(&t)
		if err != nil {
			return characters, err
		}

		characters = append(characters, &t)
	}

	if err := cur.Err(); err != nil {
		return characters, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return characters, nil
}

func GetCharacterById(db *mongo.Client, language string, id string) (*model.Character, error) {
	ctx := context.Background()

	character := &model.Character{}
	err := db.Database("genshindata_"+language).Collection("characters").FindOne(ctx, bson.M{"id": id}).Decode(&character)

	return character, err
}

func GetCommonMaterials(db *mongo.Client, language string) ([]*model.CommonMaterial, error) {
	var commonMaterials []*model.CommonMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("common_materials").Find(ctx, bson.D{})
	if err != nil {
		return commonMaterials, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.CommonMaterial
		err := cur.Decode(&t)
		if err != nil {
			return commonMaterials, err
		}

		commonMaterials = append(commonMaterials, &t)
	}

	if err := cur.Err(); err != nil {
		return commonMaterials, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return commonMaterials, nil
}

func GetCommonMaterialById(db *mongo.Client, language string, id string) (*model.CommonMaterial, error) {
	ctx := context.Background()

	commonMaterial := &model.CommonMaterial{}
	err := db.Database("genshindata_"+language).Collection("common_materials").FindOne(ctx, bson.M{"id": id}).Decode(&commonMaterial)

	return commonMaterial, err
}

func GetElementalStoneMaterials(db *mongo.Client, language string) ([]*model.ElementalStoneMaterial, error) {
	var elementalStoneMaterial []*model.ElementalStoneMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("elemental_stone_materials").Find(ctx, bson.D{})
	if err != nil {
		return elementalStoneMaterial, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.ElementalStoneMaterial
		err := cur.Decode(&t)
		if err != nil {
			return elementalStoneMaterial, err
		}

		elementalStoneMaterial = append(elementalStoneMaterial, &t)
	}

	if err := cur.Err(); err != nil {
		return elementalStoneMaterial, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return elementalStoneMaterial, nil
}

func GetElementalStoneMaterialById(db *mongo.Client, language string, id string) (*model.ElementalStoneMaterial, error) {
	ctx := context.Background()

	elementalStoneMaterial := &model.ElementalStoneMaterial{}
	err := db.Database("genshindata_"+language).Collection("elemental_stone_materials").FindOne(ctx, bson.M{"id": id}).Decode(&elementalStoneMaterial)

	return elementalStoneMaterial, err
}

func GetFish(db *mongo.Client, language string) ([]*model.Fish, error) {
	var list []*model.Fish
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("fish").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Fish
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetFishingRods(db *mongo.Client, language string) ([]*model.FishingRod, error) {
	var list []*model.FishingRod
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("fishing_rod").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.FishingRod
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetBaits(db *mongo.Client, language string) ([]*model.Bait, error) {
	var list []*model.Bait
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("bait").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Bait
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetFood(db *mongo.Client, language string) ([]*model.Food, error) {
	var list []*model.Food
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("food").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Food
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetIngredients(db *mongo.Client, language string) ([]*model.Ingredient, error) {
	var list []*model.Ingredient
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("ingredients").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Ingredient
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetJewelMaterials(db *mongo.Client, language string) ([]*model.JewelMaterial, error) {
	var list []*model.JewelMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("jewels_materials").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.JewelMaterial
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetLocalMaterials(db *mongo.Client, language string) ([]*model.LocalMaterial, error) {
	var list []*model.LocalMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("local_materials").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.LocalMaterial
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetPotions(db *mongo.Client, language string) ([]*model.Potion, error) {
	var list []*model.Potion
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("potions").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Potion
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetTalentLvlUpMaterials(db *mongo.Client, language string) ([]*model.TalentLvlUpMaterial, error) {
	var list []*model.TalentLvlUpMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("talent_lvl_up_materials").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.TalentLvlUpMaterial
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetWeaponPrimaryMaterials(db *mongo.Client, language string) ([]*model.WeaponPrimaryMaterial, error) {
	var list []*model.WeaponPrimaryMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("weapon_primary_materials").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.WeaponPrimaryMaterial
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetWeaponSecondaryMaterials(db *mongo.Client, language string) ([]*model.WeaponSecondaryMaterial, error) {
	var list []*model.WeaponSecondaryMaterial
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("weapon_secondary_materials").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.WeaponSecondaryMaterial
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetWeapons(db *mongo.Client, language string) ([]*model.Weapon, error) {
	var list []*model.Weapon
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("weapons").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.Weapon
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}

		list = append(list, &t)
	}

	if err := cur.Err(); err != nil {
		return list, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return list, nil
}

func GetWeaponById(db *mongo.Client, language string, id string) (*model.Weapon, error) {
	ctx := context.Background()

	obj := &model.Weapon{}
	err := db.Database("genshindata_"+language).Collection("weapons").FindOne(ctx, bson.M{"id": id}).Decode(&obj)

	return obj, err
}
