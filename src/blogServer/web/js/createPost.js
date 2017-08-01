$(document).ready(function(){
    if (api.readCookie('accessToken')==null){
        window.location.replace("login.html");

    }

    api.changeLogin()


    createPostDisplay.clearData()

    $("#createPostButton").off("click")
    $("#createPostButton").click(function(){
        event.preventDefault()
        createPostDisplay.getData()
        createPostDisplay.sendData()
    })
})

var createPost = function () {
    this.title = ""
    this.postDescription = ""
    this.postImageUrl = ""
    this.postText = ""
}

createPost.prototype.clearData = function(){
    this.title = ""
    this.postDescription = ""
    this.postImageUrl = ""
    this.postText = ""
}

createPost.prototype.getData = function(){
    var self = this

    var postTitle = $("#postTitle").val()
    var postDescription = $("#postDescription").val()
    var postImageUrl = $("#postImageURL").val()
    var postText = $("#postBody").val()


    self.title = postTitle
    self.postDescription = postDescription
    self.postImageUrl = postImageUrl
    self.postText = postText
}

createPost.prototype.sendData = function(){
    var self = this

    userParams = {
        "accessToken" : api.readCookie('accessToken'),
    }

    getURL = '/api/user?accessToken=' + userParams.accessToken

    $.ajax({
            url: getURL,
            type: 'GET',
            success: function(result){
                var responseObj = JSON.parse(result)
                var authorName = responseObj.name


                var params = {
                        "owner" : responseObj.id,
                        "title" : self.title,
                        "postDescription" : self.postDescription,
                        "imageURL": self.postImageUrl,
                        "authorName": authorName,
                        "written": moment().unix(),
                        "created": moment().unix(),
                        "postText": self.postText,
                    }

                    $.ajax({
                            url: '/api/createPost',
                            type: 'POST',
                            data: {"p" : JSON.stringify(params)},
                            success: function(result){
                                console.log("Success")
                                window.location.replace("index.html");
                            },
                            error: function(xhr,status,error) {
                                console.log("Failure")
                            }
                        });
            },
            error: function(xhr,status,error) {}
        });




}

var createPostDisplay = new createPost()