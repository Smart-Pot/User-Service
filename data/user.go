package data

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type UserPublicData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Image     string `json:"image"`
}

type User struct {
	ID            string   `json:"id"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	Image         string   `json:"image"`
	Date          string   `json:"date"`
	Active        bool     `json:"active"`
	Authorization int      `json:"authorization"`
	Devices       []string `json:"devices"`
	OAuth 		  bool     `json:"oauth"`	
}

func generateID() string {
	return uuid.NewString()
}



func UpdateUser(ctx context.Context, user User) error {
	filter := bson.M{"id": user.ID}

	updateUser := bson.M{"$set": bson.M{
		"password":     user.Password,
		"firstname":    user.FirstName,
		"lastname":     user.LastName,
		"image":        user.Image,
		"active":       user.Active,
		"authorization": user.Authorization,
	}}

	res, err := collection.UpdateOne(ctx, filter, updateUser)
	if err != nil {
		return err
	}

	if res.ModifiedCount <= 0 {
		return errors.New("user can not updated")
	}

	return nil
}

func UpdateUserRecord(ctx context.Context, id, key string, value interface{}) error {
	filter := bson.M{"id": id}

	updateUser := bson.M{"$set": bson.M{
		key: value,
	}}

	res, err := collection.UpdateOne(ctx, filter, updateUser)

	if res.ModifiedCount <= 0 {
		return errors.New("record can not updated")
	}

	if err != nil {
		return err
	}

	return nil
}

func GetUsersPublicData(ctx context.Context, userIDList []string) ([]*UserPublicData, error) {
	var results []*UserPublicData
	cur, err := collection.Find(ctx, bson.M{"id": bson.M{"$in": userIDList}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var upd UserPublicData
		err := cur.Decode(&upd)
		if err != nil {
			return nil, err
		}

		results = append(results, &upd)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return results, err
}

func GetUserByID(ctx context.Context, id string) (*UserPublicData, error) {
	res := collection.FindOne(ctx, bson.M{"id": id})
	var u UserPublicData
	if err := res.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
