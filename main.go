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

var homePageContent string = `
	<h2>Save carbon by building faster web apps in Go.</h2>
	<p>Edit the <strong>main.go</strong> file to get started.</p>

	<h1>GoGreener</h1>

	<h3>Docs</h3>
	<p>Find in-depth information about GoGreener.</p>
	<h3>Features</h3>
	<p>Explore the unique features of this Go framework.</p>
	<h3>Deploy</h3>
	<p>Follow these deployment instructions.</p>
	`

var css string = `
	body {
		height: 100vh;
		height: 100dvh;
		background: #000;
		background: linear-gradient(0deg, rgba(255,255,255,1) 0%, rgba(214,219,220,1) 100%);
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
			s.WritePage("Hello", greener.HTMLPrintf(homePageContent))
		}
	})

	// Serve
	app.Serve(context.Background())
}
