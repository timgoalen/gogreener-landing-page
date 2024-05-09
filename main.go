package main

import (
	"context"
	"log"

	"github.com/thejimmyg/greener"
)

func main() {
	homePageContent := `
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

	app := greener.NewDefaultApp(
		greener.NewDefaultServeConfigProviderFromEnvironment(),
		greener.NewDefaultLogger(log.Printf),
		greener.NewDefaultEmptyPageProvider([]greener.Injector{}),
	)
	app.HandleWithServices("/", func(s greener.Services) {
		s.W().Write([]byte(app.Page("GoGreener", greener.HTMLPrintf(homePageContent))))
		// s.W().Write([]byte(app.Page("GoGreener", greener.HTMLPrintf(homepage))))
	})
	app.Serve(context.Background())
}
