package featuretoggle

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout            time.Duration = 5
	connString         string        = "mongodb://localhost:27017/"
	dbName             string        = "toggleWorld"
	mainCollectionName string        = "featureToggles"
	archCollectionName string        = mainCollectionName + "_archive"
)

var mainCollection *mongo.Collection
var archiveCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connString)
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mainCollection = client.Database(dbName).Collection(mainCollectionName)
	archiveCollection = client.Database(dbName).Collection(archCollectionName)

	fmt.Println("Connected to MongoDB!")
}

func getFeatureToggle(ID primitive.ObjectID) *FeatureToggle {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	var foundFeatureToggle FeatureToggle
	findResult := mainCollection.FindOne(ctx, bson.D{{"_id", ID}})
	findResult.Decode(&foundFeatureToggle)

	return &foundFeatureToggle
}

// not actually remove - move to a different collection.
func removeFeatureToggle(ID primitive.ObjectID) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	_, err := archiveCollection.InsertOne(ctx, getFeatureToggle(ID))
	if err != nil {
		return 0
	}
	deleteResult, err := mainCollection.DeleteOne(ctx, bson.D{{"_id", ID}})
	if err != nil {
		return 0
	}
	return deleteResult.DeletedCount
}

func getFeatureToggleList() []FeatureToggle {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	cursor, err := mainCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var featureToggles []FeatureToggle
	if err = cursor.All(context.TODO(), &featureToggles); err != nil {
		log.Fatal(err)
	}
	return featureToggles
}

func addOrUpdateFeatureToggle(ft FeatureToggle) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	addOrUpdateID := primitive.NilObjectID
	if ft.ID != primitive.NilObjectID {
		// an update
		_, err := mainCollection.ReplaceOne(ctx, bson.D{{"_id", ft.ID}}, ft)
		if err != nil {
			return addOrUpdateID, err
		}
		addOrUpdateID = ft.ID
	} else {
		// an insert
		ft.ID = primitive.NewObjectID()
		insertResult, err := mainCollection.InsertOne(ctx, ft)
		if err != nil {
			return addOrUpdateID, err
		}
		addOrUpdateID = insertResult.InsertedID.(primitive.ObjectID)
	}

	return addOrUpdateID, nil
}
