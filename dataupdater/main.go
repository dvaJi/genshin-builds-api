package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "3000"

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Printf("Running in development mode")
		r.Use(cors.Default())
	}

	Token := os.Getenv("TOKEN")

	r.POST("/updateData", func(c *gin.Context) {

		// Check Authorization header
		authHeader := c.Request.Header.Get("Authorization")
		log.Printf("authHeader: %s", authHeader)
		log.Printf("Token: %s", Token)

		if authHeader != "Bearer "+Token {
			log.Printf("Authorization header is not valid")
			c.JSON(200, gin.H{
				"message": "Hello World!",
			})
			return
		}

		// Init DB
		ctx := context.Background()
		client := StartDb(ctx)

		// Get files from request
		form, _ := c.MultipartForm()
		files := form.File["data"]

		log.Println("Files: ", files)

		for _, file := range files {
			log.Println(file.Filename)
			language := strings.ReplaceAll(strings.ReplaceAll(file.Filename, "data_", ""), ".min.json", "")

			log.Println("Connecting to DB", language)
			dbo := client.Database("genshindata_" + language)

			f, err := file.Open()
			if err != nil {
				log.Println("Error opening file")
				log.Println(err)
				return
			}

			c.SaveUploadedFile(file, "./tmp/"+file.Filename)
			log.Println("File saved")

			defer f.Close()

			// open saved file
			jsonFile, err := ioutil.ReadFile("./tmp/" + file.Filename)
			if err != nil {
				log.Println("Error opening file")
				log.Println(err)
				return
			}

			mp := make(map[string]interface{})
			// Decode JSON into our map
			errJson := json.Unmarshal([]byte(jsonFile), &mp)
			if errJson != nil {
				log.Println("Error decoding JSON")
				println(err)
				return
			}

			// iterate over map
			for key, value := range mp {
				// Drop and create collection
				dbo.Collection(key).Drop(ctx)
				collection := dbo.Collection(key)

				collection.InsertMany(ctx, value.([]interface{}))
				log.Printf("[%s] Collection %s created", language, key)
			}

			// Delete file
			os.Remove("./tmp/" + file.Filename)
		}

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})

		client.Disconnect(ctx)
		log.Printf("Disconnected")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r.Run(":" + port)
}

func StartDb(ctx context.Context) *mongo.Client {
	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	uriDB := "mongodb://" + username + ":" + password + "@" + host + ":27017"
	log.Printf("Connecting to %s", uriDB)
	clientOptions := options.Client().ApplyURI(uriDB)

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
