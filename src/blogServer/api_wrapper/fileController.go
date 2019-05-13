package api_wrapper

import (
	"net/http"
	"strings"
)

func ServeHelperFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if (path == "/" || path == "/index.html"){
		http.ServeFile(w,r,"web/index.html")
	}

	if (path == "/post.html"){
		http.ServeFile(w,r,"web/post.html")
	}

	if (path == "/about.html"){
		http.ServeFile(w,r,"web/about.html")
	}

	if (path == "/contact.html"){
		http.ServeFile(w,r,"web/contact.html")
	}

	if (path == "/createPost.html"){
		http.ServeFile(w,r,"web/createPost.html")
	}

	if (path == "/login.html"){
		http.ServeFile(w,r,"web/login.html")
	}

	if (path == "/signup.html"){
		http.ServeFile(w,r,"web/signup.html")
	}

	if (path == "/socket.html"){
		http.ServeFile(w,r,"web/socket.html")
	}



	pathSplits := strings.Split(r.URL.Path[1:], "/")

	if (pathSplits[0]=="js")||
		(pathSplits[0]=="img")||
		(pathSplits[0]=="vendor")||
		(pathSplits[0]=="css"){
		http.ServeFile(w,r,"web/"+r.URL.Path)
	}
}