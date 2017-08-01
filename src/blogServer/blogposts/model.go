package blogposts

import "gopkg.in/mgo.v2/bson"

type BlogPost struct {
	Id 		bson.ObjectId	`json:"id" bson:"_id"`
	Owner 		bson.ObjectId	`json:"owner" bson:"owner"`
	Title 		string 		`json:"title" bson:"title"`
	PostDescription	string		`json:"postDescription" bson:"postDescription"`
	ImageURL	string		`json:"imageURL" bson:"imageURL"`
	AuthorName	string		`json:"authorName" bson:"authorName"`
	Created 	int64		`json:"created" bson:"created"`
	Written 	int64		`json:"written" bson:"written"`
	PostText 	string		`json:"postText" bson:"postText"`
}

type GetPostsParams struct {
	Owner  		string 	`json:"owner"`
	StartDate 	int64 	`json:"startDate"`
	EndDate		int64 	`json:"endDate"`
}

