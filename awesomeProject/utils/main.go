//package utils
//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"go.mongodb.org/mongo-driver/mongo/readpref"
//	"log"
//	"time"
//)
//
//func main() {
//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://zarinatoleukhanova:Madina2005@zarina.vb4wvvt.mongodb.net/?retryWrites=true&w=majority&appName=Zarina"))
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	err = client.Connect(ctx)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Disconnect(ctx)
//	err = client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal(err)
//	}
//	databases, err := client.ListDatabaseNames(ctx, bson.D{})
//	fmt.Printf("%v", databases)
//}
