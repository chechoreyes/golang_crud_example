package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chechoreyes/go-react-crud/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// DotENV file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoDBURI := os.Getenv("MONGODB_URI")

	app := fiber.New()

	app.Use(cors.New())

	// MongoDB connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBURI))

	app.Post("/users", func(c *fiber.Ctx) error {

		var user models.User
		//Parsea los valores que viene del frontend
		c.BodyParser(&user)

		// insert
		coll := client.Database("goFiberExample").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key: "name", Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	if err != nil {
		panic(err)
	}

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {

		var users []models.User

		coll := client.Database("goFiberExample").Collection("users")

		results, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			panic(err)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
		})
	})

	fmt.Println("Server on port 3000")
	app.Listen(":" + port)

}
