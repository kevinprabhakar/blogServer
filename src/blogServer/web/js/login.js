$(document).ready(function(){
    if (api.readCookie('accessToken')!=null){
        api.eraseCookie(
            'accessToken'
        )
    }

    $("#navLogin").hide()

    $("#loginButton").off('click')
    $("#loginButton").click(function(){
        event.preventDefault()
        LoginDisplay.getData()
    })
})

var Login = function(){
    this.temp = ""
}

Login.prototype.getData = function(){
    var self = this

    var email = $("#emailForm").val()
    var password = $("#passwordForm").val()

    var signInParams = {
        "email" : email,
        "password" : password
    }

    jsonParams = JSON.stringify(signInParams)

    $.ajax({
        url: '/api/signIn',
        type: 'POST',
        data: {"p" : jsonParams},
        success: function(result){
              var responseObj = JSON.parse(result)
              self.loadCookie(responseObj.signedString)
        }
    });
}

Login.prototype.loadCookie = function(accessToken){
    api.createCookie('accessToken', accessToken, 7)
    window.location.replace("index.html")
}

var LoginDisplay = new Login()