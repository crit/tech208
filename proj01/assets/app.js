(function($){
	var timer = {
		init: 0,
		duration: 0,
		start: function() {
			this.init = new Date().getTime();
		},
		stop: function() {
			this.duration = new Date().getTime() - this.init;
		}
	};

	var restartButton = $("<button>").addClass("btn btn-primary").text("restart").on("click", function(e){
		current();
		$(this).parent().removeClass("bg-danger");
		$(this).remove();
	});

	function current() {
		timer.start()

		$.ajax({
			url: "/current",
			dataType: "json"
		}).complete(list);
	}

	function list(res) {
		var data, content = [];

		timer.stop();
		
		if (res.status != 200) {
			$("#timer").text("-").parent().parent().addClass("bg-danger").append(restartButton);
			return
		}

		data = res.responseJSON;

		if (!data) {
			$("#timer").text("-").parent().parent().addClass("bg-danger").append(res.responseText);
			return
		}
		
		$.each(data.current, function(i, item){
			var el = $("<li>").addClass('list-group-item').text(item)
			content.push(el)
		});

		if (content.length > 0) $("#stage").html(content);
		$("#timer").text(timer.duration);

		setTimeout(current, 2000);
	}

	// on dom load bind stuff and start the loop
	$(function(){
		$("input:first").focus();
		current();
	});
}(jQuery));