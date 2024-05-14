package main

import (
	"context"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/thejimmyg/greener"
)

var wwwFiles embed.FS

type SimpleConfig struct{}

func NewSimpleConfig() *SimpleConfig {
	return &SimpleConfig{}
}

// An example of injecting a component which needs both the SimpleApp and SimpleServices
type WritePageProvider interface {
	WritePage(title string, body template.HTML)
}

type DefaultWritePageProvider struct {
	greener.ResponseWriterProvider
	greener.EmptyPageProvider
}

func (d *DefaultWritePageProvider) WritePage(title string, body template.HTML) {
	d.W().Write([]byte(d.Page(title, body)))
}

func NewDefaultWritePageProvider(emptyPageProvider greener.EmptyPageProvider, responseWriterProvider greener.ResponseWriterProvider) *DefaultWritePageProvider {
	return &DefaultWritePageProvider{EmptyPageProvider: emptyPageProvider, ResponseWriterProvider: responseWriterProvider}
}

type SimpleServices struct {
	greener.Services
	WritePageProvider // Here is the interface we are extending the serivces with
}

type SimpleApp struct {
	greener.App
	*SimpleConfig
}

func NewSimpleApp(app greener.App, simpleConfig *SimpleConfig) *SimpleApp {
	return &SimpleApp{
		App:          app,
		SimpleConfig: simpleConfig,
	}
}

func (app *SimpleApp) HandleWithSimpleServices(path string, handler func(*SimpleServices)) {
	app.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		services := app.NewServices(w, r)
		s := &SimpleServices{Services: services} // We have to leave WritePageProvider nil temporarily
		writePageProvider := NewDefaultWritePageProvider(app, s)
		s.WritePageProvider = writePageProvider // Now we set it here.
		handler(s)
	})
}

// TODO: import the google font properly.
var homePageContent string = `
	<style>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100..900&family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap');
	</style>

	<h2 class="headline">Save carbon by building faster web apps in Go.</h2>

	<div class="get-started-container">
		<div class="get-started">Edit the <strong>main.go</strong> file to get started.</div>
	</div>

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
	
	<section class="links">
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
				<p class="link-body">Share your site with the help of these instructions.</p>
			</div>
		</a>
	</section>
	`

var css string = `

	* {
		margin: 0;
		padding: 0;
	}

	a {
  		text-decoration: none;
  		color: inherit;
	}

	body {
		height: 100vh;
		height: 100dvh;
		background: linear-gradient(0deg, #000 0%, #2E2E2E 100%);
		font-family: "Inter";
		text-align: center;
	}

	.headline {
		font-size: 1.5rem;
		font-family: "Roboto Mono";
		font-weight: 400;
		margin: 2rem;
		color: #FFF;
	}

	.get-started-container {
		display: grid;
		place-items: center;
	}

	.get-started {
		font-size: 1.1rem;
		font-family: "Roboto Mono";
		border: 1px solid #545454;
		border-radius: 8px;
		padding: 0.5rem;
		color: #D2D2D2;
	}

	.circle-container,
	.circle-glow-container,
	.title-container {
		position: absolute;
		top: 0;
		left: 0;
		display: grid;
		place-items: center;
		height: 100vh;
		height: 100dvh;
		width: 100%;
	}

	.circle-glow-container {
		position: relative;
	}

	.circle {
		height: 400px;
		width: 400px;
		border-radius: 50%;
		display: grid;
		place-items: center;
		background: radial-gradient(circle, #000 42%, rgb(61 61 61 / 85%) 100%);
	}

	.circle-glow {
		height: 550px;
		width: 550px;
		border-radius: 50%;
		background-color: #6df36a;
		opacity: 4%;
		filter: blur(91px);
		position: absolute;
		top: -7%;
		left: 27%;
		z-index: -1;
	}
	
	.circle-blob {
		height: 130px;
		width: 130px;
		border-radius: 50%;
		background-color: #6df36a;
		opacity: 24%;
		filter: blur(72px);
		position: absolute;
		top: 34%;
		left: 51%;
	}

	.title {
		font-size: 3rem;
		cursor: default;
		background: linear-gradient(to right, rgba(248, 119, 0, 0.65), #FFF, rgb(55, 153, 107, 0.75));
		-webkit-background-clip: text;
  		-webkit-text-fill-color: transparent;
	}

	.links {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 100%;
		display: flex;
		justify-content: space-evenly;
		flex-direction: row;
		padding-bottom: 2rem;
	}

	.link-container {
		display: flex;
	}

	.link-bullet-point {
		height: 24px;
		width: 24px;
		border-radius: 50%;
		background-color: #F87700;
		opacity: 22%;
		filter: blur(8px);
		margin-top: 1.1rem;
		transition: height 0.2s ease-in-out;
	}

	.link-container:hover .link-bullet-point {
		height: 48px
	}

	.link-content {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		max-width: 200px;
		margin: 1rem 0 1rem 1rem;
		cursor: pointer;
	}

	.link-title {
		font-size: 1.5rem;
		font-weight: 500;
		opacity: 85%;
		transition: opacity 0.2s ease-in-out;
		color: #FFF;
	}

	.link-body {
		text-align: left;
		font-size: 0.9rem;
		padding-top: 0.3rem;
		color: #929292;
		opacity: 85%;
		transition: opacity 0.2s ease-in-out;
	}

	.link-container:hover .link-body,
	.link-container:hover .link-title {
		opacity: 100%;
	}
	`

func main() {
	// Setup
	wwwFS, _ := fs.Sub(wwwFiles, "www") // Used for the icon and the static file serving
	uiSupport := []greener.UISupport{greener.NewDefaultUISupport(
		css,
		`console.log("Hello from script");`,
		`console.log("Hello from service worker");`,
	)}
	themeColor := "#000000"
	appShortName := "Simple"
	config := greener.NewDefaultServeConfigProviderFromEnvironment()
	logger := greener.NewDefaultLogger(log.Printf)
	injectors := []greener.Injector{
		greener.NewDefaultStyleInjector(logger, uiSupport),
		greener.NewDefaultScriptInjector(logger, uiSupport),
		greener.NewDefaultServiceWorkerInjector(logger, uiSupport),
		greener.NewDefaultThemeColorInjector(logger, themeColor),
		greener.NewDefaultIconsInjector(logger, wwwFS),
		greener.NewDefaultManifestInjector(logger, appShortName),
	}
	emptyPageProvider := greener.NewDefaultEmptyPageProvider(injectors)
	static := greener.NewCompressedFileHandler(http.FS(wwwFS))

	// Routes
	app := NewSimpleApp(greener.NewDefaultApp(config, logger, emptyPageProvider), NewSimpleConfig())
	app.HandleWithSimpleServices("/", func(s *SimpleServices) {
		if s.R().URL.Path != "/" {
			// If no other route is matched and the request is not for / then serve a static file
			static.ServeHTTP(s.W(), s.R())
		} else {
			// Let's use our new WritePageProvider, instead of this version that uses app and s separately
			// app.Page("Hello", greener.Text("Hello <>!")).WriteHTMLTo(s.W())
			// s.WritePage("Hello", greener.Text(homePageContent))
			s.WritePage("GoGreener", greener.HTMLPrintf(homePageContent))
		}
	})

	// Serve
	app.Serve(context.Background())
}
