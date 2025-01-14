package main

import (
	"context"
	"log"
	"os"

	"github.com/msyrus/gmdhaka2025/repository"
	"github.com/msyrus/gmdhaka2025/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().
		SetBSONOptions(&options.BSONOptions{
			UseJSONStructTags: true,
		}).
		ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalln(err)
	}

	db := dbClient.Database(os.Getenv("MONGO_DB"))

	// users := repository.NewUserWithHelper(db)
	// users := repository.NewUserWithHelper(db)
	users := repository.NewUserGeneric(db)

	userSvc := service.NewUser(users)

	userSvc.Get(ctx, "5f7b3b3b7f7b7b7b7b7b7b7b")
}
