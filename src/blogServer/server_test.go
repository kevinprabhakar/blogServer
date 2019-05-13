package main

import(
	"testing"
	"time"
	"encoding/json"
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"blog_server/src/blogServer/blogposts"
	"github.com/PuerkitoBio/goquery"
	"blog_server/src/blogServer/user"
	"bitbucket.org/rajeevpra/itip/mongo"
	"gopkg.in/mgo.v2/bson"
)

func TestCreatePost(t *testing.T){
	urlText:= "http://saganipsum.com/?p=10"
	doc, _ := goquery.NewDocument(urlText)

	var postText string = ""

	doc.Find(".universe").Each(func(j int, selection *goquery.Selection){
		postText = selection.Text()

	})
	newPost := blogposts.BlogPost{
		Title: "Space, the Final Frontier",
		PostDescription: "Why Space is like Super Cool, Man",
		ImageURL: "/img/post-bg.jpg",
		AuthorName:"Kevin Prabhakar",
		Created: time.Now().Unix(),
		Written: time.Now().Unix(),
		PostText: postText,
	}

	params := url.Values{}

	jsonForm, err := json.Marshal(newPost)

	if (err != nil){
		t.Fail()
	}

	params.Set("p", string(jsonForm))

	_, respErr := http.PostForm("http://localhost:3000/api/createPost", params)

	if (respErr != nil){
		fmt.Println(respErr.Error())
		t.Fail()
	}
}

//func TestGetBlogPosts(t *testing.T){
//	params := blogposts.GetPostsParams{time.Now().AddDate(0,0,-7).Unix(), time.Now().Unix()}
//
//	jsonForm, err := json.Marshal(params)
//
//
//	if (err != nil){
//		t.Fail()
//	}
//
//	urlParams := url.Values{}
//	urlParams.Set("p",string(jsonForm))
//
//	response, responseErr  := http.PostForm("http://localhost:3000/api/getBlogPosts", urlParams)
//
//	if (responseErr != nil){
//		fmt.Println(responseErr.Error())
//		t.Fail()
//	}else{
//		respBody, _ := ioutil.ReadAll(response.Body)
//		fmt.Println(string(respBody))
//	}
//}

func TestBaconIpsum(t *testing.T){
	url:= "http://saganipsum.com/?p=10"
	doc, _ := goquery.NewDocument(url)

	doc.Find(".universe").Each(func(j int, selection *goquery.Selection){
		fmt.Println(selection.Text())

	})
}

func TestEmptyStruct(t *testing.T){
	type SignInParams struct {
		Email		string 		`json:"email"`
		Password 	string 		`json:"password"`
	}

	var a SignInParams

	if (a == SignInParams{}){
		fmt.Println("True")
	}

	a = SignInParams{"a","b"}

	if (a == SignInParams{}){
		fmt.Println("True2")
	}
}

func TestSignUp(t *testing.T){
	params := user.SignUpParams{"kevin.sury1a@gmail.com","bottle","Kevin Prabhakar"}
	urlParams := url.Values{}

	jsonParams, _ := json.Marshal(params)

	urlParams.Set("p", string(jsonParams))

	respBody, respErr := http.PostForm("http://localhost:3000/api/SignUp", urlParams)

	if (respErr != nil){
		fmt.Println(respErr.Error())
		t.Fail()
	}else{
		respBodyString, _ := ioutil.ReadAll(respBody.Body)
		fmt.Println(string(respBodyString))
	}
}

func TestSignIn(t *testing.T) {
	params := user.SignInParams{"kevin.surya@gmail.com", "bottle"}
	urlParams := url.Values{}

	jsonParams, _ := json.Marshal(params)

	urlParams.Set("p", string(jsonParams))

	respBody, respErr := http.PostForm("http://localhost:3000/api/SignIn", urlParams)

	if (respErr != nil) {
		fmt.Println(respErr.Error())
		t.Fail()
	} else {
		respBodyString, _ := ioutil.ReadAll(respBody.Body)
		fmt.Println(string(respBodyString))
	}
}

func TestGetUser(t *testing.T){
	params := url.Values{}

	params.Set("accessToken", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDEyODc2NzAsInVpZCI6IjU5NzZiNDUzNzkxZTRiYmE0YTA2YzA4OCJ9.aqKN83QVew8hfaIpFryNBMhTt_Doay5p4tahibUjke0")

	respBody, respErr := http.PostForm("http://localhost/api/user", params)


	if (respErr != nil) {
		fmt.Println(respErr.Error())
		t.Fail()
	} else {
		respBodyString, _ := ioutil.ReadAll(respBody.Body)
		fmt.Println(string(respBodyString))
	}
}

func TestHashing(t *testing.T){
	fmt.Println(user.HashPassword("jprabhakar@gmail.com"))
	fmt.Println(user.HashPassword("bottle1235"))

}

func TestIterations(t *testing.T){
	iter := mongo.GetUserCollection(mongo.GetMongoDB(mongoSession)).Find(nil).Sort("$natural").Tail(5*time.Second)

	var result user.User
	var lastId bson.ObjectId

	for {
		for iter.Next(&result) {
			fmt.Println(result.Id)
			lastId = result.Id
		}
		if iter.Err() != nil {
			iter.Close()
		}
		if iter.Timeout() {
			continue
		}
		query := mongo.GetUserCollection(mongo.GetMongoDB(mongoSession)).Find(bson.M{"_id": bson.M{"$gt": lastId}})
		iter = query.Sort("$natural").Tail(5 * time.Second)
	}
	iter.Close()
}