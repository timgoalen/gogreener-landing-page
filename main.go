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
	<div class="get-started">Edit the <strong>main.go</strong> file to get started.</div>

	<div class="circle-container">
		<div class="circle">
		</div>
	</div>

	<div class="title-container">
		<h1 class="title">GoGreener</h1>
	</div>
	
	<div class="links">
		<div class="link">
			<h3>Docs</h3>
			<p>Find in-depth information about GoGreener.</p>
		</div>
		<div class="link">
			<h3>Features</h3>
			<p>Explore the unique features of this Go framework.</p>
		</div>
		<div class="link">
			<h3>Deploy</h3>
			<p>Follow these deployment instructions.</p>
		</div>
	</div>
	`

var css string = `
	body {
		height: 100vh;
		height: 100dvh;
		background: #000;
		background: linear-gradient(0deg, rgba(255,255,255,1) 0%, rgba(214,219,220,1) 100%);
		font-family: "Inter";
		text-align: center;
	}

	.get-started {
		font-size: 1.25rem;
		font-family: "Roboto Mono";
	}

	.headline {
		font-size: 1.5rem;
		font-family: "Roboto Mono";
		font-weight: 400;
	}

	.circle-container,
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

	.circle {
		height: 400px;
		width: 400px;
		border-radius: 50%;
		display: grid;
		place-items: center;
		background: #FFF;
		background: radial-gradient(circle, #D6DBDC 0%, #FFF 80%);
		filter: blur(4px);
	}

	.title {
		font-size: 3rem;
	}

	.links {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 100%;
		display: grid;
		grid-template-columns: repeat(3, 1fr)
	}

	.link {
		display: flex;
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
