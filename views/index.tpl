{{ template "includes/header" }}
	<div class="content">
		<a href="https://github.com/wstucco/purrraceholder"><img style="position: relative; top: -32px; left: -32px; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_left_darkblue_121621.png" alt="Fork me on GitHub"></a>	
		<h1>Welcome to the Purrraceholder</h1>
		<h3>Image placeholder service powered by the Grumpy cat</h3>
		<br/>

		<h3>How to use the Purrraceholder</h3>
		<p>Just put your image size after the URL and get the placeholder.</p>
		<p>for example</p>
		<p><a href="/700">/700</a> generate a square image 700px wide and tall</p>
		<p><a href="/700/300">/700/300</a> generate an image 700px wide and 300px tall</p>
		<p>max size is 2560px for both dimensions</p>

		<p>View the <a href="{{ .ImageUrl }}"> latest purrraceholder</a></p>
  </div>
{{ template "includes/footer" }}

