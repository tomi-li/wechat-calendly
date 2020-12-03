package models

import (
	"calendly/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID(id string) (*User, error) {
	_db := db.GetDB()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := _db.Collection("post").InsertOne(ctx, User{ID: id})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(res)
	}

	user := User{
		ID: id,
	}
	return &user, nil
}

func (h User) ListAll() []User {
	_db := db.GetDB()
	filter := bson.D{}
	cur, err := _db.Collection("post").Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	var results []User
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	return results
}
