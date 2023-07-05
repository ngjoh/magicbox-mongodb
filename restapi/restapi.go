package restapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

const description = `
	
Service  for managing Microsoft 365 resources

## Getting started 

### Authentication
You need a credential key to access the API. The credential is issue by [niels.johansen@nexigroup.com](mailto:niels.johansen@nexigroup.com).

Use the credential key to get an access token through the /v1/authorize end point. The access token is valid for 10 minutes.

Pass the access token in the Authorization header as a Bearer token to access the API.
	`

func Core() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Koksmat Magicbox CORE"
	s.OpenAPI.Info.WithDescription(description)
	s.OpenAPI.Info.Version = "v1.0.0"

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	s.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	s.Post("/authorize", signin())
	//s.Get("/blob/{tag}", getBlob())
	s.Method(http.MethodGet, "/blob/{tag}", nethttp.NewHandler(getBlob()))

	s.MethodFunc(http.MethodPost, "/api/v1/subscription/notify", validateSubscription)
	s.Method(http.MethodGet, "/v1/business/countries", nethttp.NewHandler(getCountries()))
	s.Method(http.MethodGet, "/v1/business/units", nethttp.NewHandler(getBusinessAndGroupUnits()))

	s.Route("/v1/info", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))
			r.Method(http.MethodGet, "/", nethttp.NewHandler(getInfo()))
		})
	})

	s.Docs("/docs/core", swgui.New)
	s.Mount("/debug/core", middleware.Profiler())

	log.Println("Server starting")
	if err := http.ListenAndServe(":4321", s); err != nil {
		log.Fatal(err)
	}
}
func Run() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Koksmat Magicbox"
	s.OpenAPI.Info.WithDescription(fmt.Sprintf("%s %s", description, `
## Version History
### V2.0.0 - Parameter name changed
Changed parameter names from id to exchangeObjectId in relevant endpoints, breaking compatibility with previous versions hence relasing as a new major version

### V1.0.0 - Initial version

	`))
	s.OpenAPI.Info.Version = "v2.0.0"

	jwtAuth := Authenticator

	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	s.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	s.Post("/authorize", signin())
	//s.Post("/v1/demo", demo())
	// Endpoints with user access.
	s.Route("/v1/sharedmailboxes", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(adminAuth, nethttp.HTTPBasicSecurityMiddleware(s.OpenAPICollector, "User", "User access"))
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodPost, "/", nethttp.NewHandler(createSharedMailbox()))
			r.Method(http.MethodGet, "/{exchangeObjectId}", nethttp.NewHandler(getSharedMailbox()))
			r.Method(http.MethodPatch, "/{exchangeObjectId}", nethttp.NewHandler(updateSharedMailbox()))
			r.Method(http.MethodPost, "/{exchangeObjectId}/smtp", nethttp.NewHandler(addSharedMailboxEmail()))
			r.Method(http.MethodPatch, "/{exchangeObjectId}/primarysmtp", nethttp.NewHandler(updateSharedMailboxPrimaryEmailAddress()))
			r.Method(http.MethodDelete, "/{exchangeObjectId}/smtp", nethttp.NewHandler(removeSharedMailboxEmail()))
			r.Method(http.MethodPost, "/{exchangeObjectId}/members", nethttp.NewHandler(addSharedMailboxMembers()))
			r.Method(http.MethodDelete, "/{exchangeObjectId}/members", nethttp.NewHandler(removeSharedMailboxMembers()))
			r.Method(http.MethodPatch, "/{exchangeObjectId}/owners", nethttp.NewHandler(setSharedMailboxOwners()))

			r.Method(http.MethodPost, "/{exchangeObjectId}/readers", nethttp.NewHandler(addSharedMailboxReaders()))
			r.Method(http.MethodDelete, "/{exchangeObjectId}/readers", nethttp.NewHandler(removeSharedMailboxReaders()))
			r.Method(http.MethodGet, "/", nethttp.NewHandler(listSharedMailbox()))
			r.Method(http.MethodDelete, "/{exchangeObjectId}", nethttp.NewHandler(deleteSharedMailbox()))
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
func Admin() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Koksmat Magicbox ADMIN"
	s.OpenAPI.Info.WithDescription(description)
	s.OpenAPI.Info.Version = "v0.0.1"

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	s.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	s.Post("/authorize", signin())

	s.Route("/v1/admin", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))
			r.Method(http.MethodGet, "/auditlogs/date/{date}", nethttp.NewHandler(getAuditLogs()))
			r.Method(http.MethodGet, "/auditlogs/powershell/{objectId}", nethttp.NewHandler(getAuditLogPowershell()))
		})
	})

	s.Docs("/docs/admin", swgui.New)
	s.Mount("/debug/admin", middleware.Profiler())

	log.Println("Server starting")
	if err := http.ListenAndServe(":4322", s); err != nil {
		log.Fatal(err)
	}
}
