package main

import (
	"log"
	"net/http"
	"os"

	genshindata "src/genshindata"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dvaJi/genshin-builds-api/graph"
	"github.com/dvaJi/genshin-builds-api/graph/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CorsRequest struct {
	Url string `json:"url"`
}

const defaultPort = "8080"

func graphqlHandler(dbClient *mongo.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: dbClient}}))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.POST{})

	if os.Getenv("ENVIRONMENT") == "development" {
		h.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func simpleCorsProxy() gin.HandlerFunc {
	client := resty.New()

	return func(c *gin.Context) {
		var corsReq CorsRequest

		if errBind := c.BindJSON(&corsReq); errBind != nil {
			return
		}

		if corsReq.Url == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "url is missing",
			})
			return
		}

		resp, err := client.R().
			EnableTrace().
			Post(corsReq.Url)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Data(http.StatusOK, resp.Header().Get("Content-Type"), resp.Body())
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConfig := &genshindata.DBImpl{
		DbUserName: os.Getenv("DATABASE_USER"),
		DbPassword: os.Getenv("DATABASE_PASSWORD"),
		DbHost:     os.Getenv("DATABASE_HOST"),
	}
	// dbConfig := &genshindata.DBImpl{
	// 	DbUserName: "root",
	// 	DbPassword: "example",
	// 	DbHost:     "localhost",
	// }

	dbClient := genshindata.Init(dbConfig)

	r := gin.Default()

	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("Running in development mode")
		r.Use(cors.Default())
	}

	r.POST("/query", graphqlHandler(dbClient))

	if os.Getenv("ENVIRONMENT") == "development" {
		r.GET("/", playgroundHandler())
	}

	r.POST("/corsproxy", simpleCorsProxy())

	http.ListenAndServe(":"+port, r)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
