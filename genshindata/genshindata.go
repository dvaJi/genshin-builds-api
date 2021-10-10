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
	log.Printf(language)
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
