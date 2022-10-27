package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"im/models"
	"testing"
	"time"
)

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	fmt.Println("db =======>", db)
	ub := new(models.UserBasic)
	fmt.Println("ub1 =======>", ub)

	err = db.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ub2 =======>", ub)
}

func TestFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	cursor, err := db.Collection("user_room").Find(context.Background(), bson.D{})
	urs := make([]*models.UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(models.UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}
	for _, v := range urs {
		fmt.Println(" UserRoom ===>", v)

	}
}
