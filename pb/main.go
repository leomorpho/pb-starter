package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/stripe/stripe-go/v79"

	stripehandlers "pocketbase/internal/stripe"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := pocketbase.New()

	// Configure Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

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

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}