package data

import (
	"context"
	"errors"
	"time"

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
	ID            string `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Image         string `json:"image"`
	Date          string `json:"date"`
	Active        bool   `json:"active"`
	Authorization int    `json:"authorization"`
}

func generateID() string {
	return uuid.NewString()
}

func CreateUser(ctx context.Context, newUser *User) error {
	newUser.Date = time.Now().UTC().String()
	newUser.ID = generateID()
	newUser.Image = ""
	newUser.Authorization = 0
	newUser.Active = false
	_, err := collection.InsertOne(ctx, *newUser)

	return err
}

func UpdateUser(ctx context.Context, updatedUser User) error {
	filter := bson.M{"id": updatedUser.ID}

	updateUser := bson.M{"$set": bson.M{
		"password":     updatedUser.Password,
		"firstname":    updatedUser.FirstName,
		"lastname":     updatedUser.LastName,
		"image":        updatedUser.Image,
		"active":       updatedUser.Active,
		"authorizatio": updatedUser.Authorization,
	}}

	res, err := collection.UpdateOne(ctx, filter, updateUser)
	if err != nil {
		return err
	}

	if res.ModifiedCount <= 0 {
		return errors.New("image change failed")
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
		return errors.New("image change failed")
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
