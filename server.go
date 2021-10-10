package main

import (
	"log"
	"net/http"
	"os"

	genshindata "src/genshindata"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dvaJi/genshin-builds-api/graph"
	"github.com/dvaJi/genshin-builds-api/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// dbConfig := &genshindata.DBImpl{
	// 	DbUserName: os.Getenv("DATABASE_USERNAME"),
	// 	DbPassword: os.Getenv("DATABASE_PASSWORD"),
	// 	DbHost:     os.Getenv("DATABASE_HOST"),
	// }
	dbConfig := &genshindata.DBImpl{
		DbUserName: "root",
		DbPassword: "example",
		DbHost:     "localhost",
	}

	dbClient := genshindata.Init(dbConfig)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: dbClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
