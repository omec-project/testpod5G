package smf

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

func HostServices(addr string) {
	router := NewRouter()
	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(router, h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nsmf-pdusession/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"ReleasePduSession",
		strings.ToUpper("Post"),
		"/pdu-sessions/:pduSessionRef/release",
		ReleasePduSession,
	},

	{
		"UpdatePduSession",
		strings.ToUpper("Post"),
		"/pdu-sessions/:pduSessionRef/modify",
		UpdatePduSession,
	},

	{
		"ReleaseSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/release",
		HTTPReleaseSmContext,
	},

	{
		"RetrieveSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/retrieve",
		RetrieveSmContext,
	},

	{
		"UpdateSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/modify",
		HTTPUpdateSmContext,
	},

	{
		"PostPduSessions",
		strings.ToUpper("Post"),
		"/pdu-sessions",
		PostPduSessions,
	},

	{
		"PostSmContexts",
		strings.ToUpper("Post"),
		"/sm-contexts",
		HTTPPostSmContexts,
	},
}
