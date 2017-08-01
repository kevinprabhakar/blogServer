$(document).ready(function(){
    $("#passwordCheckAlert").hide()

    api.changeLogin()

    if (api.readCookie('accessToken')!=null){
        api.eraseCookie(
            'accessToken'
        )
    }

    $("#alertCloseButton").off('click')
    $("#alertCloseButton").click(function(){
            event.preventDefault()
            window.location.replace("signup.html")
        })
    $("#SignUpButton").off('click')
    $("#SignUpButton").click(function(){
        event.preventDefault()
        SignUpDisplay.getData()
    })
})

var SignUp = function(){
    this.temp = ""
}

SignUp.prototype.getData = function(){
    var self = this

    var email = $("#emailForm").val()
    var password = $("#passwordForm").val()
    var passwordCheck = $("#passwordReenterForm").val()
    var name = $("#nameForm").val()

    if (password != passwordCheck){
        $("#passwordCheckAlert").show()
        setTimeout(function(){
            window.location.replace("signup.html")
        },2000)
    } else {
        var signInParams = {
                "name" : name,
                "email" : email,
                "password" : password
            }

            jsonParams = JSON.stringify(signInParams)

            $.ajax({
                url: '/api/signUp',
                type: 'POST',
                data: {"p" : jsonParams},
                success: function(result){
                      var responseObj = JSON.parse(result)
                      self.loadCookie(responseObj.signedString)
                }
            });
    }


}

SignUp.prototype.loadCookie = function(accessToken){
    api.createCookie('accessToken', accessToken, 7)
    window.location.replace("index.html");

}

var SignUpDisplay = new SignUp()