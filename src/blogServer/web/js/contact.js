$(document).ready(function(){

    if (api.readCookie('accessToken')==null){
        window.location.replace("login.html");

    }

    api.changeLogin()

})