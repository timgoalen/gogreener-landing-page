package main

import (
	"context"
	"log"

	"github.com/thejimmyg/greener"
)

func main() {
	// Create UI Support (args: CSS, JavaScript, Service Worker)
	uiSupport := greener.NewDefaultUISupport(GlobalCSS, "", "")

	// Set up logging
	logger := greener.NewDefaultLogger(log.Printf)

	// Set up the application
	app := greener.NewDefaultApp(
		greener.NewDefaultServeConfigProviderFromEnvironment(),
		logger,
		greener.NewDefaultEmptyPageProvider([]greener.Injector{
			// Inject CSS
			greener.NewDefaultStyleInjector(logger, []greener.UISupport{uiSupport}),
		}),
	)

	// Define a simple route handler
	app.HandleWithServices("/", func(s greener.Services) {
		s.W().Write([]byte(app.Page("GoGreener", greener.HTMLPrintf(HomePageHTML))))
	})

	// Start the server
	app.Serve(context.Background())
}
