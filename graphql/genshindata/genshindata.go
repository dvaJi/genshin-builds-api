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

func GetCharacters(db *mongo.Client, language string) ([]*model.CharacterInfo, error) {
	var characters []*model.CharacterInfo
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("characters").Find(ctx, bson.D{})
	if err != nil {
		return characters, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.CharacterInfo
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

func GetWeapons(db *mongo.Client, language string) ([]*model.WeaponInfo, error) {
	var list []*model.WeaponInfo
	ctx := context.Background()

	cur, err := db.Database("genshindata_"+language).Collection("weapons").Find(ctx, bson.D{})
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.WeaponInfo
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

func GetWeaponExpMaterials(db *mongo.Client, language string) ([]*model.ExpMaterial, error) {
	var list []*model.ExpMaterial
	ctx := context.Background()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"exp", -1}})
	cur, err := db.Database("genshindata_"+language).Collection("weapon_enhancement_material").Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.ExpMaterial
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

func GetCharacterExpMaterials(db *mongo.Client, language string) ([]*model.ExpMaterial, error) {
	var list []*model.ExpMaterial
	ctx := context.Background()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"exp", -1}})
	cur, err := db.Database("genshindata_"+language).Collection("character_exp_material").Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return list, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t model.ExpMaterial
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

func CharacterLvlExpList() [90]int {
	return [90]int{
		0, 1000, 2325, 4025, 6175, 8800, 11950, 15675, 20025, 25025, 30725, 37175,
		44400, 52450, 61375, 71200, 81950, 93675, 106400, 120175, 135050, 151850,
		169850, 189100, 209650, 231525, 254775, 279425, 305525, 333100, 362200,
		392850, 425100, 458975, 494525, 531775, 570750, 611500, 654075, 698500,
		744800, 795425, 848125, 902900, 959800, 1018875, 1080150, 1143675, 1209475,
		1277600, 1348075, 1424575, 1503625, 1585275, 1669550, 1756500, 1846150,
		1938550, 2033725, 2131725, 2232600, 2341550, 2453600, 2568775, 2687100,
		2808625, 2933400, 3061475, 3192875, 3327650, 3465825, 3614525, 3766900,
		3922975, 4082800, 4246400, 4413825, 4585125, 4760350, 4939525, 5122700,
		5338925, 5581950, 5855050, 6161850, 6506450, 6893400, 7327825, 7815450,
		8362650,
	}
}

func WeaponLvlExpList() [3][90]int {
	return [3][90]int{
		{
			0, 275, 700, 1300, 2100, 3125, 4400, 5950, 7800, 9975, 12475, 15350, 18600,
			22250, 26300, 30800, 35750, 41150, 47050, 53475, 60400, 68250, 76675, 85725,
			95400, 105725, 116700, 128350, 140700, 153750, 167550, 182075, 197375,
			213475, 230375, 248075, 266625, 286025, 306300, 327475, 349525, 373675,
			398800, 424925, 452075, 480275, 509525, 539850, 571275, 603825, 637475,
			674025, 711800, 750800, 791075, 832625, 875475, 919625, 965125, 1011975,
			1060200, 1112275, 1165825, 1220875, 1277425, 1335525, 1395175, 1456400,
			1519200, 1583600, 1649625, 1720700, 1793525, 1868100, 1944450, 2022600,
			2102600, 2184450, 2268150, 2353725, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3988200,
		},
		{
			0, 400, 1025, 1925, 3125, 4675, 6625, 8975, 11775, 15075, 18875, 23225,
			28150, 33675, 39825, 46625, 54125, 62325, 71275, 81000, 91500, 103400,
			116175, 129875, 144525, 160150, 176775, 194425, 213125, 232900, 253800,
			275825, 299025, 323400, 349000, 375825, 403925, 433325, 464050, 496125,
			529550, 566125, 604200, 643800, 684950, 727675, 772000, 817950, 865550,
			914850, 965850, 1021225, 1078450, 1137550, 1198575, 1261525, 1326450,
			1393350, 1462275, 1533250, 1606300, 1685200, 1766325, 1849725, 1935425,
			2023450, 2113825, 2206575, 2301725, 2399300, 2499350, 2607025, 2717350,
			2830350, 2946050, 3064475, 3185675, 3309675, 3436500, 3566175, 3698750,
			3855225, 4031100, 4228700, 4450675, 4699975, 4979925, 5294175, 5646875,
			6042650,
		},
		{
			0, 600, 1550, 2900, 4700, 7025, 9950, 13475, 17675, 22625, 28325, 34850,
			42250, 50550, 59775, 69975, 81225, 93525, 106950, 121550, 137300, 155150,
			174325, 194875, 216850, 240300, 265250, 291725, 319775, 349450, 380800,
			413850, 448650, 485225, 523625, 563875, 606025, 650125, 696225, 744350,
			794500, 849375, 906500, 965900, 1027625, 1091725, 1158225, 1227150, 1298550,
			1372500, 1449000, 1532075, 1617925, 1706575, 1798125, 1892550, 1989950,
			2090300, 2193700, 2300175, 2409750, 2528100, 2649800, 2774900, 2903450,
			3035500, 3171075, 3310200, 3452925, 3599300, 3749375, 3910900, 4076400,
			4245900, 4419450, 4597100, 4778900, 4964900, 5155150, 5349675, 5548550,
			5783275, 6047100, 6343500, 6676475, 7050425, 7470350, 7941725, 8470775,
			9064450,
		},
	}
}
