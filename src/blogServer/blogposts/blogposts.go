package blogposts

import (
	"encoding/json"
	"strings"
	"blog_server/src/blogServer/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type BlogPostController struct {
	mongoSession 	*mgo.Session
}

type returnBlogPost struct {
	Id 		string		`json:"id" bson:"_id"`
	Owner 		string 		`json:"owner" bson:"owner"`
	Title 		string 		`json:"title" bson:"title"`
	PostDescription	string		`json:"postDescription" bson:"postDescription"`
	ImageURL	string		`json:"imageURL" bson:"imageURL"`
	AuthorName	string		`json:"authorName" bson:"authorName"`
	Created 	int64		`json:"created" bson:"created"`
	Written 	int64		`json:"written" bson:"written"`
	PostText 	string		`json:"postText" bson:"postText"`
}

type createBlogPostParams struct {
	Owner 		string 		`json:"owner" bson:"owner"`
	Title 		string 		`json:"title" bson:"title"`
	PostDescription	string		`json:"postDescription" bson:"postDescription"`
	ImageURL	string		`json:"imageURL" bson:"imageURL"`
	AuthorName	string		`json:"authorName" bson:"authorName"`
	Created 	int64		`json:"created" bson:"created"`
	Written 	int64		`json:"written" bson:"written"`
	PostText 	string		`json:"postText" bson:"postText"`
}

func NewBlogPostController(session *mgo.Session) *BlogPostController {
	return &BlogPostController{session}
}

//params
//{"owner": <bson string id>, "title":<string for title>, "postDescription":<string>, "imageURL":<string>, "authorName":<string>, "created":<unix time>, "written":<unix time>, "postText":<string>}
func (self *BlogPostController) CreateBlogPost(params string) (error){
	dec := json.NewDecoder(strings.NewReader(params))

	var newPostParams createBlogPostParams
	err := dec.Decode(&newPostParams)

	if (err != nil){
		return err
	}

	if (!bson.IsObjectIdHex(newPostParams.Owner)){
		return errors.New("Invalid Owner BSON ID")
	}

	newPost := BlogPost{
		Id: bson.NewObjectId(),
		Owner: bson.ObjectIdHex(newPostParams.Owner),
		Title: newPostParams.Title,
		PostDescription: newPostParams.PostDescription,
		ImageURL: newPostParams.ImageURL,
		AuthorName:newPostParams.AuthorName,
		Created:newPostParams.Created,
		Written:newPostParams.Written,
		PostText:newPostParams.PostText,
	}

	blogPostCollection := mongo.GetBlogPostCollection(mongo.GetDataBase(self.mongoSession))

	insertErr := blogPostCollection.Insert(newPost)
	if (insertErr != nil){
		return insertErr
	}

	userCollection := mongo.GetUserCollection(mongo.GetDataBase(self.mongoSession))

	updateErr := userCollection.Update(
		bson.M{"_id":bson.ObjectIdHex(newPostParams.Owner)},
		bson.M{"$push":bson.M{
			"postIDList" : newPost.Id,
		}})

	if (updateErr != nil){
		return updateErr
	}
	return nil
}

//params
//{"owner":<bson uid as string>, "startDate":<int64 unix time>, "endDate":<int64 unix time>}
func (self *BlogPostController) GetBlogPosts(params string) ([]BlogPost, error){
	dec := json.NewDecoder(strings.NewReader(params))
	var postParams GetPostsParams
	decodeErr := dec.Decode(&postParams)

	if (decodeErr != nil){
		return []BlogPost{}, decodeErr
	}

	if (!bson.IsObjectIdHex(postParams.Owner)){
		return []BlogPost{}, errors.New("Invalid Owner BSON ID")
	}

	query := bson.M{
		"owner" : bson.ObjectIdHex(postParams.Owner),
		"written" : bson.M{
			"$gte": postParams.StartDate,
			"$lte": postParams.EndDate,

		},
	}

	blogPostCollection := mongo.GetBlogPostCollection(mongo.GetDataBase(self.mongoSession))

	blogPostList := make([]BlogPost,0)

	findErr := blogPostCollection.Find(query).All(&blogPostList)

	if (findErr != nil){
		return []BlogPost{}, findErr
	}

	return blogPostList, nil
}

func (self *BlogPostController) GetBlogPostFromID(postID string) (returnBlogPost, error){
	var returnPost returnBlogPost
	var receiverPost BlogPost

	if (!bson.IsObjectIdHex(postID)){
		return returnPost, errors.New("Post ID given is not valid BSON ID")
	}

	query := bson.M{
		"_id" : bson.ObjectIdHex(postID),
	}

	blogPostCollection := mongo.GetBlogPostCollection(mongo.GetDataBase(self.mongoSession))

	findErr := blogPostCollection.Find(query).One(&receiverPost)

	if (findErr != nil){
		return returnPost, findErr
	}

	newPost := returnBlogPost{
		receiverPost.Id.Hex(),
		receiverPost.Owner.Hex(),
		receiverPost.Title,
		receiverPost.PostDescription,
		receiverPost.ImageURL,
		receiverPost.AuthorName,
		receiverPost.Created,
		receiverPost.Written,
		receiverPost.PostText,
	}

	return newPost, nil
}

