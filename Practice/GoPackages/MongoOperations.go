package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	MongoConnection()
}
func MongoConnection() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	//list the databases
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	fmt.Println(databases, err)

	databaseName := "practice"
	collectionName := "sample"

	//Inside database level operations like
	// client.Database(databaseName).CreateCollection()
	//working in particular collection
	//Create,delete and List collections
	//create index
	//Aggregation on database level

	//Here in particular collection operations like CRUD, aggregation operations
	// mongoCursor, err := client.Database(databaseName).Collection(collectionName).

	// InsertMongoDocuments(ctx, client, databaseName, collectionName)
	// ReadMongoDocuments(ctx, client, databaseName, collectionName)
	// DeleteMongoDocument(ctx, client, databaseName, collectionName)

}

func DeleteMongoDocument(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) {
	deleteFilter := bson.D{{"name", "sathish"}}
	DeleteResult, err := client.Database(databaseName).Collection(collectionName).DeleteOne(ctx, deleteFilter)
	fmt.Println("DeleteResult", DeleteResult, err)
}

func InsertMongoDocuments(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) {
	var insertDocument = mongoDocument{1, "sathish", "india"}
	//inserting one document
	insertResult, err := client.Database(databaseName).Collection(collectionName).InsertOne(ctx, insertDocument)
	fmt.Println("insertResult", insertResult, err)

}

func ReadMongoDocuments(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) {

	//passing different filter conditions format
	// filterCondition := bson.D{}  //this empty read filter reads all documents from colelction
	// filterCondition := bson.D{{"name", "sathish"}}  //two sets of paranthesis
	filterCondition := bson.D{{"id", bson.M{"$gt": 0, "$lt": 2}}}

	results := []mongoDocument{}
	mongoCursor, err := client.Database(databaseName).Collection(collectionName).Find(ctx, filterCondition)
	if err != nil {
		panic(err)
	}
	for mongoCursor.Next(ctx) {
		var document mongoDocument
		err := mongoCursor.Decode(&document)
		if err != nil {
			panic(err)
		}
		results = append(results, document)
	}

	//closing this read cursor
	mongoCursor.Close(ctx)

	fmt.Println("mongo documents find() --", results) //[{0  } {0  } {0  } {1 sathish india} {1 sathish india}]   //{0  } {0  } {0  }  -- these documents are not suited with our reading mongoDocument struct.
}

// For mongo both write and read queries - Should use structs, So we can converts structs fields into mongo bson's fields.
type mongoDocument struct {
	Id      int    `bson:"id,omitempty`
	Name    string `bson:"name,omitempty`
	Country string `bson:"country,omitempty`
}
