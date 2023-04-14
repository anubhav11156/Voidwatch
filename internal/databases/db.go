package database

import (
	"context"
	"fmt"
	// "os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection

var client *mongo.Client

const atlasConnectionUri = "mongodb+srv://anubhav11697:myMongo123@myfirstcluster.hfdwigv.mongodb.net/?retryWrites=true&w=majority"

func establisConnection() {

	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	
	opts := options.Client().ApplyURI(atlasConnectionUri)

	client, err := mongo.Connect(context.TODO(), opts)

	if err !=nil {
		// panic(err)
		fmt.Println(err)
	}


	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
		// fmt.Println(err)
	  }
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// fmt.Println("------Options------")
	// fmt.Println("1. Disconnect")
	// fmt.Println("2. Check connection status")
	// fmt.Println("3. Show Databases")
	// fmt.Println("4. Show Collections")
	// fmt.Println("5. Exit")
	



	// for true {
	// 	fmt.Print("Enter option : ")
	// 	var choice int
	// 	fmt.Scan(&choice)
	// 	switch choice {
	// 	case 1:
	// 		myMongoDisconnect(client)
	// 		break
	// 	case 2:
	// 		checkConnectionStatus(client)
	// 		break
	// 	case 3:
	// 		showDatabases(client)
	// 		break
	// 	case 4:
	// 		showCollections(client)
	// 		break
	// 	case 5:
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("Invalid input!")
	// 	}
	
	// }

}

// func myMongoDisconnect(client *mongo.Client) {
	

// 	err := client.Disconnect(context.TODO())

// 	if(err != nil ){
// 		fmt.Println(err)
// 	}else {
// 		fmt.Println("Disconnected")
// 	}
// }

// func showDatabases(client *mongo.Client) {
// 	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Print("Databases: ",)

// 	for _, db := range databases {
// 		fmt.Print(db)
// 		fmt.Print(" ")
// 	}
// 	fmt.Println("")
// 	}	

	
// }

// func checkConnectionStatus(client *mongo.Client) {
// 	err := client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Connected to MongoDB!")
// 	}
// }

// func showCollections(client *mongo.Client) {
// 	var dbName string
// 	fmt.Print("Enter database : ")
// 	fmt.Scan(&dbName)

// 	db := client.Database(dbName)

// 	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Print("Collections : ",)

// 	for _, coll := range collections {
// 		fmt.Print(coll)
// 		fmt.Print(" ")
// 	}
// 	fmt.Println("")
// 	}	
// }

func getMongoClient() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(atlasConnectionUri)

	client, err := mongo.Connect(context.TODO(), opts)

	if err !=nil {
		// panic(err)
		fmt.Println(err)
	}

	return client, nil
}

