$(document).ready(function(){
    if (api.readCookie('accessToken')==null){
        window.location.replace("login.html");
    }

    api.changeLogin()
    PostDisplayActual.loadPost()
})

var PostDisplay = function() {
    this.ID = ""
    this.Title = ""
    this.PostDescription = ""
    this.ImageURL = ""
    this.AuthorName = ""
    this.Created = 0
    this.Written = 0
    this.PostText = ""
}

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

PostDisplay.prototype.loadPost = function() {
    var self = this

    var postID = getParameterByName('postID')

    $.ajax({
            url: '/api/getBlogPostWithID',
            type: 'POST',
            data: {"postID" : postID},
            success: function(result){
                console.log("Success")

                var responseObj = JSON.parse(result)
                self.displayPost(responseObj)
            },
            error: function(xhr,status,error) {
                console.log("Failure")
            }
        });
}

PostDisplay.prototype.displayPost = function(responseObj) {
    var self = this

    var finalHtml = ""

    var BlogPost = responseObj
    
    $("#blogPost").show()

    var BlogPostTemplate = '<header class="intro-header" style="background-image: url(%imageURL%)">\
                                                              <div class="container">\
                                                                  <div class="row">\
                                                                      <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">\
                                                                          <div class="post-heading">\
                                                                              <h1>%title%</h1>\
                                                                              <h2 class="subheading">%postDescription%</h2>\
                                                                              <span class="meta">Posted by <a href="about.html">%authorName%</a> on %written%</span>\
                                                                          </div>\
                                                                      </div>\
                                                                  </div>\
                                                              </div>\
                                                          </header>\
                                                          <article>\
                                                              <div class="container">\
                                                                  <div class="row">\
                                                                      <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1 text-justify">%postText%</div>\
                                                                  </div>\
                                                              </div>\
                                                          </article>\
                                                          <hr>'

    var realImageURL = "'" + BlogPost.imageURL + "'"
    
    var blogPostData = {
                "%imageURL%" : realImageURL,
                "%title%" : BlogPost.title,
                "%postDescription%" : BlogPost.postDescription,
                "%authorName%" : BlogPost.authorName,
                "%written%": moment.unix(BlogPost.written).format('MM/DD/YYYY'),
                "%postText%" : BlogPost.postText,
            }
    
    finalHtml += BlogPostTemplate.replace(/%\w+%/g, function(all) {
                                                if(typeof(blogPostData[all]) != 'undefined') return blogPostData[all];
                                                else return all;
                                            });



    
    $('#blogPost').html(finalHtml)

}

var PostDisplayActual = new PostDisplay()
