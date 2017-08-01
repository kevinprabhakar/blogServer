$(document).ready(function() {
    if (api.readCookie('accessToken')==null){
        window.location.replace("login.html");

    }

    api.changeLogin()

    BlogPostList.getBlogPosts()
});

var BlogEntriesList = function() {
    this.userID = ""
    this.entries = []
}

BlogEntriesList.prototype.getBlogPosts = function() {
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
                self.userId = responseObj.id

                var params = {
                        "owner"     : self.userId,
                        "startDate" : parseInt(moment().startOf('year').valueOf()/1000,10),
                        "endDate"   : parseInt(moment().valueOf().valueOf()/1000,10)
                    }

                    $.ajax({
                        url: '/api/getBlogPosts',
                        type: 'POST',
                        data: {"p" : JSON.stringify(params)},
                        success: function(result){
                            var responseObj = JSON.parse(result)
                            self.loadPosts(responseObj)
                        },
                        error: function(xhr,status,error) {}
                    });
            },
            error: function(xhr,status,error) {}
        });
}

BlogEntriesList.prototype.loadPosts = function (blogPosts) {
    var self = this

    $('#blog_entries_list').show()
    var BlogPostTemplate =  "<div class='post-preview'>\
                                   <a href=%postURL%>\
                                       <h2 class='post-title'>\
                                           %title%\
                                       </h2>\
                                       <h3 class='post-subtitle'>\
                                           %titleDescription%\
                                       </h3>\
                                   </a>\
                                   <p class='post-meta'>Written by <a href='about.html'>%authorName%</a> on %written%</p>\
                               </div>\
                               <hr>"

    self.entries = blogPosts

    var finalHtml = ""

    if (blogPosts.length == 0){
        finalHtml += "<h1>No Posts To Display...Try Writing Some</h1>"
    }else{
        for (var blogPostCount = 0 ; blogPostCount < blogPosts.length ; blogPostCount++){
                var postURL = "post.html?" + "postID=" + blogPosts[blogPostCount].id

                var blogPostData = {
                    "%postURL%" : postURL,
                    "%title%" : blogPosts[blogPostCount].title,
                    "%titleDescription%" : blogPosts[blogPostCount].postDescription,
                    "%authorName%" : blogPosts[blogPostCount].authorName,
                    "%written%": moment.unix(blogPosts[blogPostCount].written).format('MM/DD/YYYY'),

                }

                finalHtml += BlogPostTemplate.replace(/%\w+%/g, function(all) {
                                                            if(typeof(blogPostData[all]) != 'undefined') return blogPostData[all];
                                                            else return all;
                                                        });


            }
    }

    $('#blog_entries_list').html(finalHtml)




}


var BlogPostList = new BlogEntriesList()