package restapi

import (
	"context"
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
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

var mySigningKey = []byte("AllYourBase")

func issueToken() usecase.Interactor {
	type JWT struct {
		Token string
	}
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *string) error {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": 234,
		})

		tokenString, err := token.SignedString(mySigningKey)
		*output = tokenString
		return err

	})

	u.SetTitle("Issue a token")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Authentication")
	return u
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
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
func Run() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Magicbox"
	s.OpenAPI.Info.WithDescription("MagicBox integration for managing Microsoft 365 resources")
	s.OpenAPI.Info.Version = "v0.0.1"

	//adminAuth := middleware.BasicAuth("Admin Access", map[string]string{"admin": "admin"})

	jwtAuth := Authenticator

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	s.Post("/v1/login", issueToken())
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

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)
	s.Mount("/debug", middleware.Profiler())
	// Start server.
	log.Println("Server starting")
	if err := http.ListenAndServe(":5001", s); err != nil {
		log.Fatal(err)
	}
}
