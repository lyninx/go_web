console.log("hello!")

$('#create-form').submit(function () {
    // custom handling here
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