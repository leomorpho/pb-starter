package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/stripe/stripe-go/v79"

	stripehandlers "pocketbase/internal/stripe"
	tushandlers "pocketbase/internal/tus"
	"pocketbase/webauthn"
	_ "pocketbase/migrations"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := pocketbase.New()

	// Check if we're running with `go run`
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	// Register the migrate command
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun, // Auto-migrate during development
	})

	// Configure Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Register WebAuthn
	webauthn.Register(app)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Initialize TUS handler
		tusHandler, err := tushandlers.NewTUSHandler(app)
		if err != nil {
			log.Fatal("Failed to initialize TUS handler:", err)
		}

		// Stripe routes
		se.Router.POST("/create-checkout-session", func(e *core.RequestEvent) error {
			return stripehandlers.CreateCheckoutSession(e, app)
		})

		se.Router.POST("/create-portal-link", func(e *core.RequestEvent) error {
			return stripehandlers.CreatePortalLink(e, app)
		})

		se.Router.POST("/stripe", func(e *core.RequestEvent) error {
			return stripehandlers.HandleWebhook(e, app)
		})

		// TUS upload routes - simple specific routes
		tusHandlerFunc := func(e *core.RequestEvent) error {
			tusHandler.ServeHTTP(e.Response, e.Request)
			return nil
		}
		
		// TUS protocol requires these specific endpoints
		se.Router.POST("/tus", tusHandlerFunc)      // Create upload
		se.Router.HEAD("/tus/{id}", tusHandlerFunc) // Get upload info
		se.Router.PATCH("/tus/{id}", tusHandlerFunc) // Upload chunks
		se.Router.DELETE("/tus/{id}", tusHandlerFunc) // Cancel upload
		se.Router.Any("OPTIONS /tus", tusHandlerFunc) // CORS preflight for base
		se.Router.Any("OPTIONS /tus/{id}", tusHandlerFunc) // CORS preflight for upload

		// Serve static files from the provided public dir (if exists)
		// This must be registered last as it's a catch-all route
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}