package main

// TODO: import the google fonts properly.
var HomePageHTML string = `
	<style>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100..900&family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap');
	</style>

	<header>
		<h2 class="headline">Save carbon by building faster web apps in Go.</h2>

		<div class="get-started-container">
			<div class="get-started">Edit the <strong>main.go</strong> file to get started.</div>
		</div>
	</header>

	<main>
		<div class="circle-container">
			<div class="circle"></div>
		</div>

		<div class="circle-glow-container">
				<div class="circle-glow"></div>
				<div class="circle-blob"></div>
		</div>

		<div class="title-container">
			<h1 class="title">GoGreener</h1>
		</div>
	</main>
	
	<footer class="links">
		<a class="link-container" href="https://github.com/thejimmyg/greener" rel="noopener" target="_blank">
			<div class="link-bullet-point"></div>
			<div class="link-content">
				<h3 class="link-title">Features</h3>
				<p class="link-body">Explore the unique features of this Go web framework.</p>
			</div>
		</a>

		<a class="link-container" href="https://github.com/thejimmyg/greener" rel="noopener" target="_blank">
			<div class="link-bullet-point"></div>
			<div class="link-content">
				<h3 class="link-title">Learn</h3>
				<p class="link-body">Follow a  quick-start guide or read the in-depth docs.</p>
			</div>
		</a>
		
		<a class="link-container" href="https://github.com/thejimmyg/greener" rel="noopener" target="_blank">
			<div class="link-bullet-point"></div>
			<div class="link-content">
				<h3 class="link-title">Deploy</h3>
				<p class="link-body">Launch your site with the help of these instructions.</p>
			</div>
		</a>
	</footer>
`
