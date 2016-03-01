<!-- Add the stylesheet to the head -->
<link rel="stylesheet" type="text/css" href="js/clippy.js/build/clippy.css" media="all">
<link rel="stylesheet" type="text/css" href="css/clippy-demo.css" media="all">

<p>
Welcome to the Icinga 2 API Clippy Demo :)
</p>

<p>

<table border=0>
<tr>
<th>Host Problems</th>
<th>Service Problems</th>
</tr>
<tr>
<td><div id="logger-host-problems"></td>
<td><div id="logger-service-problems"></td>
</tr>
</table>
</p>

<!-- Add these scripts to  the bottom of the page -->
<script src="js/jquery/jquery-2.2.1.min.js"></script>

<!-- Clippy.js -->
<script src="js/clippy.js/build/clippy.min.js"></script>

<!-- Init script -->
<script type="text/javascript">
	clippy.load('Clippy', function(agent){
		document.getElementById("logger-host-problems").innerHTML=0;
		document.getElementById("logger-service-problems").innerHTML=0;
		// do anything with the loaded agent
		agent.show();
		agent.moveTo(150,150);

		(function() {
			setInterval(function poll() {
			    setTimeout(function() {
				$.ajax({
				    url: "query.php",
				    type: "GET",
				    success: function(data) {
					//'data' is already a json parsed object providing the icinga2 results response

					var num_hosts_down = data.results[0].status.num_hosts_down;
					var num_services_critical = data.results[0].status.num_services_critical;
					var num_services_warning = data.results[0].status.num_services_warning;

					var num_service_problems = (num_services_critical + num_services_warning);
					var num_host_problems = num_hosts_down;

					//show details to the reader, disable later
					document.getElementById("logger-host-problems").innerHTML=num_host_problems;
					document.getElementById("logger-service-problems").innerHTML=num_service_problems;

					//update clippy based on the current values
					//the agent should alert us on problems with an animation first.
					var problem = false;
					var speak_str = "";

					if (num_host_problems > 0) {
						problem = true;
						speak_str += "Hosts: " + num_host_problems + " ";
					}
					if (num_service_problems > 0) {
						problem = true;
						speak_str += "Services: " + num_service_problems + " ";
					}

					if (problem) {
						agent.speak("PROBLEM: " + speak_str);
						agent.animate();
					} else {
						agent.speak("Everything OK.");
					}
				    },
				    dataType: "json",
				    complete: poll,
				    timeout: 2000
				})
			    }, 1000); //timeout for poll
			}, 5000) //poll icinga2 query in given internal
		})();	

	});


</script>
