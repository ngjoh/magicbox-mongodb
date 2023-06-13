package restapi

import (
	"log"
	"net/http"

	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4emb"
)

func Run() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Magicbox"
	s.OpenAPI.Info.WithDescription("MagicBox integration for managing Microsoft 365 resources")
	s.OpenAPI.Info.Version = "v0.0.1"

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	s.Post("/v1/sharedmailboxes", createSharedMailbox())
	s.Get("/v1/sharedmailboxes/{id}", getSharedMailbox())
	s.Patch("/v1/sharedmailboxes/{id}", updateSharedMailbox())
	s.Post("/v1/sharedmailboxes/{id}/smtp", addSharedMailboxEmail())
	s.Delete("/v1/sharedmailboxes/{id}/smtp", removeSharedMailboxEmail())
	s.Post("/v1/sharedmailboxes/{id}/members", addSharedMailboxMembers())
	s.Delete("/v1/sharedmailboxes/{id}/members", removeSharedMailboxMembers())
	s.Post("/v1/sharedmailboxes/{id}/owners", addSharedMailboxOwners())
	s.Delete("/v1/sharedmailboxes/{id}/owners", removeSharedMailboxOwners())
	s.Get("/v1/sharedmailboxes", listSharedMailbox())
	s.Delete("/v1/sharedmailboxes/{id}", deleteSharedMailbox())

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	// Start server.
	log.Println("Server starting")
	if err := http.ListenAndServe(":5001", s); err != nil {
		log.Fatal(err)
	}
}
