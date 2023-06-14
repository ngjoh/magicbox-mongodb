package restapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4emb"
)

func ParseIdToken(tokenString string) (appName string, appSecret string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})
	if err != nil {

		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		app := fmt.Sprint(claims["app"])
		key := fmt.Sprint(claims["key"])
		return app, key, nil

	} else {

		return "", "", errors.New("Not implemented")
	}

}

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through. It's just fine
// until you decide to write something similar and customize your client response.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := jwtauth.TokenFromHeader(r)
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
			app := fmt.Sprint(claims["app"])
			// Token is authenticated, pass it through
			ctx := context.WithValue(r.Context(), "app", app)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

	})
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
	s.OpenAPI.Info.Version = "v0.0.2"

	//adminAuth := middleware.BasicAuth("Admin Access", map[string]string{"admin": "admin"})

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	s.Post("/v1/authorize", signin())
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
			r.Method(http.MethodDelete, "/{id}/smtp", nethttp.NewHandler(removeSharedMailboxEmail()))
			r.Method(http.MethodPost, "/{id}/members", nethttp.NewHandler(addSharedMailboxMembers()))
			r.Method(http.MethodDelete, "/{id}/members", nethttp.NewHandler(removeSharedMailboxMembers()))
			r.Method(http.MethodPost, "/{id}/owners", nethttp.NewHandler(addSharedMailboxOwners()))
			r.Method(http.MethodDelete, "/{id}/owners", nethttp.NewHandler(removeSharedMailboxOwners()))
			r.Method(http.MethodGet, "/", nethttp.NewHandler(listSharedMailbox()))
			r.Method(http.MethodDelete, "/{id}", nethttp.NewHandler(deleteSharedMailbox()))
		})
	})
	s.Route("/v1/info", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			//r.Use(adminAuth, nethttp.HTTPBasicSecurityMiddleware(s.OpenAPICollector, "User", "User access"))
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))

			r.Method(http.MethodPost, "/", nethttp.NewHandler(getInfo()))

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
