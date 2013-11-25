{{ template "includes/header" }}
	<div class="content">
		<h1>Welcome to the Purrraceholder</h1>
		<h3>generate image placeholders powered by the Grumpy cat</h3>
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

