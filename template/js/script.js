$(function(){

	doLogin()


})

function doLogin() {
	var user = { "name" : "mario", "password" : "1234"}
	$.post("http://localhost:4000/api/v1/login", 
		JSON.stringify(user), 
		function(data, err){
			localStorage.setItem("token", data.token)
		},
		"json"
	)
}

function getToken() {
	return localStorage.getItem("token")
}


function users() {
	var token = getToken()
	$.ajax({
	url: 'http://localhost:4000/api/v1/users',
	type: 'GET',
	headers: {
	    "Authorization": "Bearer " + token  //for object property name, use quoted notation shown in second
	},
	dataType: 'json',
	success: function (data) {
	    
	}
	});
}

