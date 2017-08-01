package user

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id 		bson.ObjectId 	`json:"id" bson:"_id"`
	Email		string 		`json:"email" bson:"email"`
	PHash 		string 		`json:"pHash" bson:"pHash"`
	Name 		string 		`json:"name" bson:"name"`
	PostIDList	[]bson.ObjectId	`json:"postIDList" bson:"postIDList"`
	FriendsList 	[]bson.ObjectId	`json:"friendsList" bson:"friendsList"`
}

type SignUpParams struct {
	Email		string 		`json:"email"`
	Password 	string 		`json:"password"`
	Name 		string 		`json:"name"`
}

type SignInParams struct {
	Email		string 		`json:"email"`
	Password 	string 		`json:"password"`
}

type SignedString struct {
	AccessToken 	string 		`json:"signedString"`
}