console.log("hello!")

$('#create-form').submit(function () {
	var obj = {}
	$( "#create-form input, #create-form textarea" ).each(function( index ) {
		var input = $(this)
		switch(input.attr("name")){
			case "title":
				obj.title = input.val()
				break
			case "url":
				obj.url = input.val()
				break
			case "body":
				obj.content = input.val()
				break
		}
	});
	console.log(obj)
	var formData = JSON.stringify(obj)
	$.ajax({
		type: "POST",
		url: "api/create",
		data: formData,
		success: function(){ console.log("created!")},
		dataType: "json",
		contentType : "application/json"
	});
	return false
});

//if($('#index').length) {
	// get all documents and append to page
	$.get( "api/", function( data ) {
		var pages = JSON.parse(data)
		pages.forEach(function(elem, index, array){
			console.log(elem)
			$('#index').append("<a href=/"+elem.url+"><li><span class='title'>"+elem.title+"</span><span class='modified'>"+new Date(elem.modified).toDateString()+"</span></li></a>")
		})
	})
//}