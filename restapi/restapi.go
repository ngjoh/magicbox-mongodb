package restapi

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4emb"
)

type ListNotication struct {
	Value []struct {
		SubscriptionID     string    `json:"subscriptionId"`
		ClientState        string    `json:"clientState"`
		ExpirationDateTime time.Time `json:"expirationDateTime"`
		Resource           string    `json:"resource"`
		TenantID           string    `json:"tenantId"`
		SiteURL            string    `json:"siteUrl"`
		WebID              string    `json:"webId"`
	} `json:"value"`
}

func SharePoint() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Koksmat Magicbox for SharePoint"
	s.OpenAPI.Info.WithDescription(`
	
Service  for managing Microsoft 365 resources

## Getting started 

### Authentication
You need a credential key to access the API. The credential is issue by [niels.johansen@nexigroup.com](mailto:niels.johansen@nexigroup.com).

Use the credential key to get an access token through the /v1/authorize end point. The access token is valid for 10 minutes.

Pass the access token in the Authorization header as a Bearer token to access the API.
	`)
	s.OpenAPI.Info.Version = "v1.0.0"

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	s.Post("/authorize", signin())
	s.MethodFunc(http.MethodPost, "/api/v1/subscription/notify", validateSubscription)

	s.Route("/v1/info", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodGet, "/", nethttp.NewHandler(getInfo()))

		})
	})

	s.Docs("/docs", swgui.New)
	s.Mount("/debug", middleware.Profiler())

	log.Println("Server starting")
	if err := http.ListenAndServe(":4321", s); err != nil {
		log.Fatal(err)
	}
}
func Run() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Koksmat Magicbox"
	s.OpenAPI.Info.WithDescription(`
	
Service  for managing Microsoft 365 resources

## Getting started 

### Authentication
You need a credential key to access the API. The credential is issue by [niels.johansen@nexigroup.com](mailto:niels.johansen@nexigroup.com).

Use the credential key to get an access token through the /v1/authorize end point. The access token is valid for 10 minutes.

Pass the access token in the Authorization header as a Bearer token to access the API.
	`)
	s.OpenAPI.Info.Version = "v1.0.0"

	//adminAuth := middleware.BasicAuth("Admin Access", map[string]string{"admin": "admin"})

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	s.Post("/authorize", signin())
	//s.Post("/v1/demo", demo())
	// Endpoints with user access.
	s.Route("/v1/sharedmailboxes", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(adminAuth, nethttp.HTTPBasicSecurityMiddleware(s.OpenAPICollector, "User", "User access"))
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodPost, "/", nethttp.NewHandler(createSharedMailbox()))
			r.Method(http.MethodGet, "/{id}", nethttp.NewHandler(getSharedMailbox()))
			r.Method(http.MethodPatch, "/{id}", nethttp.NewHandler(updateSharedMailbox()))
			r.Method(http.MethodPost, "/{id}/smtp", nethttp.NewHandler(addSharedMailboxEmail()))
			r.Method(http.MethodPatch, "/{id}/primarysmtp", nethttp.NewHandler(updateSharedMailboxPrimaryEmailAddress()))
			r.Method(http.MethodDelete, "/{id}/smtp", nethttp.NewHandler(removeSharedMailboxEmail()))
			r.Method(http.MethodPost, "/{id}/members", nethttp.NewHandler(addSharedMailboxMembers()))
			r.Method(http.MethodDelete, "/{id}/members", nethttp.NewHandler(removeSharedMailboxMembers()))
			r.Method(http.MethodPatch, "/{id}/owners", nethttp.NewHandler(setSharedMailboxOwners()))

			r.Method(http.MethodPost, "/{id}/readers", nethttp.NewHandler(addSharedMailboxReaders()))
			r.Method(http.MethodDelete, "/{id}/readers", nethttp.NewHandler(removeSharedMailboxReaders()))
			r.Method(http.MethodGet, "/", nethttp.NewHandler(listSharedMailbox()))
			r.Method(http.MethodDelete, "/{id}", nethttp.NewHandler(deleteSharedMailbox()))
		})
	})
	s.Route("/v1/addresses", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(adminAuth, nethttp.HTTPBasicSecurityMiddleware(s.OpenAPICollector, "User", "User access"))
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodGet, "/{address}", nethttp.NewHandler(resolveAddress()))

		})
	})
	s.Route("/v1/info", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(adminAuth, nethttp.HTTPBasicSecurityMiddleware(s.OpenAPICollector, "User", "User access"))
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodGet, "/", nethttp.NewHandler(getInfo()))
			r.Method(http.MethodGet, "/domains", nethttp.NewHandler(getDomains()))

		})
	})
	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)
	s.Mount("/debug", middleware.Profiler())

	// Start server.
	log.Println("Server starting")
	if err := http.ListenAndServe(":5001", s); err != nil {
		log.Fatal(err)
	}
}
