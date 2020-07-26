package dbutil

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//TODO So I will have several database, then I need to think about how to reuse functionality.
func ConnectMongoDB(){
   clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

   client, err := mongo.Connect(context.TODO(), clientOptions)

   if err != nil {
	 	log.Fatal(err)
   }

   err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

   collection := client.Database("test").Collection("quest")
   fmt.Println(collection)

   err = client.Disconnect(context.TODO())

   if err != nil {
   		log.Fatal(err)
   }

   fmt.Println("Connection to MongoDB closed.")
}

func CloseMongoDB(){
}