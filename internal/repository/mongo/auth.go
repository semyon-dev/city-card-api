package mongo

import (
	"city-card-api/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepo) CreateUser(user models.UserWithPassword) (newUser models.UserProfile, err error) {
	ctx := context.TODO()
	user.ID = primitive.NewObjectID()
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return models.UserProfile{}, err
	}
	newUser.ID, _ = res.InsertedID.(primitive.ObjectID)
	newUser.Email = user.Email
	newUser.Name = user.Name
	newUser.MName = user.MName
	newUser.Telephone = user.Telephone
	newUser.Role = user.Role
	newUser.Surname = user.Surname
	return newUser, nil
}

func (r *mongoRepo) ReadByLoginAndPass(email, pass string) (user models.UserProfile, err error) {
	filter := bson.M{"$and": []bson.M{
		{"email": email},
		{"password": pass},
	}}
	err = r.collection.FindOne(context.TODO(), filter).Decode(&user)
	return
}

func (r *mongoRepo) ChangePassword(userID int, currentPass, newPass string) error {
	return nil
}

func (r *mongoRepo) UpdateProfile(name, surname, mname string) error {
	return nil
}

func (r *mongoRepo) DeleteUser(userID int) error {
	return nil
}
