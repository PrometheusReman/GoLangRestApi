package MongoConnection 

/* import (
	"context"
	"fmt"
	"log"
	"time"
 
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
 )
 
 func DbConnMongo() {
	 client, err := mongo.NewClient(options.Client().ApplyURI(“<<MongoDB Connection URI>>))
	 if err != nil {
		 log.Fatal(err)
	 }
	 ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	 err = client.Connect(ctx)
	 if err != nil {
			 log.Fatal(err)
	 }
	 defer client.Disconnect(ctx)
	 } */