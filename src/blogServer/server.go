package main

import(
	"net/http"
	"blog_server/src/blogServer/blogposts"
	"blog_server/src/blogServer/mongo"
	"fmt"
	"blog_server/src/blogServer/util"
	"encoding/json"
	"blog_server/src/blogServer/api_wrapper"
	"blog_server/src/blogServer/user"
	"io/ioutil"
	"os"
)

var mongoSession = mongo.GetMongoSession()

var BlogPostController = blogposts.NewBlogPostController(mongoSession)
var UserController = user.NewUserController(mongoSession)

var ServerLogger = util.NewLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

func main(){
	http.HandleFunc("/api/createPost", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		ServerLogger.Debug("Create Post Request Received")

		params := r.PostForm.Get("p")

		ServerLogger.Debug("params : " + params)

		createErr := BlogPostController.CreateBlogPost(params)
		if (createErr != nil){
			ServerLogger.ErrorMsg(createErr.Error())
			fmt.Fprintf(w, createErr.Error())
			return
		}

		fmt.Fprintf(w, util.GetNoDataSuccessResponse())
	})
	http.HandleFunc("/api/getBlogPosts", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		ServerLogger.Debug("Get Blog Posts Request Received")

		params := r.PostForm.Get("p")

		ServerLogger.Debug("params : " + params)

		blogPostList, getBlogPostListError := BlogPostController.GetBlogPosts(params)

		if (getBlogPostListError != nil){
			ServerLogger.ErrorMsg(getBlogPostListError.Error())
			fmt.Fprintf(w, getBlogPostListError.Error())
			return
		}

		jsonBlogList, jsonBlogListErr := json.Marshal(blogPostList)
		if (jsonBlogListErr != nil){
			ServerLogger.ErrorMsg(jsonBlogListErr.Error())
			fmt.Fprintf(w, jsonBlogListErr.Error())
			return
		}

		fmt.Fprintf(w, string(jsonBlogList))
	})
	http.HandleFunc("/api/getBlogPostWithID", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		ServerLogger.Debug("Get Singular Blog Post Request Received")
		postID := r.Form.Get("postID")

		blogPost, blogPostError := BlogPostController.GetBlogPostFromID(postID)

		if (blogPostError != nil){
			ServerLogger.ErrorMsg(blogPostError.Error())
			fmt.Fprintf(w, blogPostError.Error())
			return
		}

		jsonForm, jsonErr := json.Marshal(blogPost)
		if (jsonErr != nil){
			ServerLogger.ErrorMsg(jsonErr.Error())
			fmt.Fprintf(w,jsonErr.Error())
			return
		}

		ServerLogger.Debug(string(jsonForm))
		fmt.Fprintf(w, string(jsonForm))
	})

	http.HandleFunc("/api/signUp", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		params := r.Form.Get("p")

		ServerLogger.Debug("Sign Up Request Received")

		newUser, newUserErr := UserController.SignUp(params)

		ServerLogger.Debug("params : " + params)

		if (newUserErr != nil){
			ServerLogger.ErrorMsg(newUserErr.Error())
			fmt.Fprintf(w, newUserErr.Error())
			return
		}

		accessToken, accessTokenErr := user.GetAccessToken(newUser.Id.Hex())

		if (accessTokenErr != nil){
			ServerLogger.ErrorMsg(accessTokenErr.Error())
			fmt.Fprintf(w, accessTokenErr.Error())
			return
		}

		accessTokenJson, accessTokenJsonErr := json.Marshal(user.SignedString{accessToken})

		if (accessTokenJsonErr != nil){
			ServerLogger.ErrorMsg(accessTokenJsonErr.Error())
			fmt.Fprintf(w,accessTokenJsonErr.Error())
			return
		}

		ServerLogger.Debug(string(accessTokenJson))
		fmt.Fprintf(w, string(accessTokenJson))
	})

	http.HandleFunc("/api/signIn", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		params := r.Form.Get("p")

		ServerLogger.Debug("Sign In Request Received")

		newUser, newUserErr := UserController.SignIn(params)

		ServerLogger.Debug("params : " + params)

		if (newUserErr != nil){
			ServerLogger.ErrorMsg(newUserErr.Error())
			fmt.Fprintf(w, newUserErr.Error())
			return
		}

		accessToken, accessTokenErr := user.GetAccessToken(newUser.Id.Hex())

		if (accessTokenErr != nil){
			ServerLogger.ErrorMsg(accessTokenErr.Error())
			fmt.Fprintf(w, accessTokenErr.Error())
			return
		}

		accessTokenJson, accessTokenJsonErr := json.Marshal(user.SignedString{accessToken})

		if (accessTokenJsonErr != nil){
			ServerLogger.ErrorMsg(accessTokenJsonErr.Error())
			fmt.Fprintf(w,accessTokenJsonErr.Error())
			return
		}

		ServerLogger.Debug(string(accessTokenJson))
		fmt.Fprintf(w, string(accessTokenJson))
	})

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		ServerLogger.Debug("User Request Received")

		accessToken := r.Form.Get("accessToken")

		findUser, findUserErr := UserController.GetUser(accessToken)

		if (findUserErr != nil){
			ServerLogger.Debug(findUserErr.Error())
			fmt.Fprintf(w, findUserErr.Error())
			return
		}

		jsonForm, jsonFormErr := json.Marshal(findUser)
		if (jsonFormErr != nil){
			ServerLogger.ErrorMsg(jsonFormErr.Error())
			fmt.Fprintf(w, jsonFormErr.Error())
			return
		}

		ServerLogger.Debug(string(jsonForm))
		fmt.Fprintf(w, string(jsonForm))
	})

	http.HandleFunc("/",api_wrapper.ServeHelperFile)

	http.ListenAndServe(":3000",nil)
}
