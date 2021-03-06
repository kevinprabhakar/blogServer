var BlogServer = function(){
    this.accessTokenTimerId = -1
}

BlogServer.prototype.createCookie = function(name,value,days) {
	if (days) {
		var date = new Date();
		date.setTime(date.getTime()+(days*24*60*60*1000));
		var expires = "; expires="+date.toGMTString();
	}
	else var expires = "";
	document.cookie = name+"="+value+expires+"; path=/";
	console.log("success")
}

BlogServer.prototype.readCookie = function(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for(var i=0;i < ca.length;i++) {
		var c = ca[i];
		while (c.charAt(0)==' ') c = c.substring(1,c.length);
		if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
	}
	return null;
}

BlogServer.prototype.eraseCookie = function(name) {
    var self = this
	self.createCookie(name,"",-1);
}

BlogServer.prototype.changeLogin = function() {
    var self = this

    if (self.readCookie('accessToken')==null){
        $('#navLogin').text('Login')
    }else{
        $('#navLogin').text('Logout')
    }
}

var api = new BlogServer()