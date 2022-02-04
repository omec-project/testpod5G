// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package nrf

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
type RoutesDisc []Route

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
	AddServiceDisc(router)
	return router
}

func AddServiceDisc(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nnrf-disc/v1")

	for _, route := range routesDisc {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}

	return group
}

var routesDisc = RoutesDisc{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"SearchNFInstances",
		strings.ToUpper("Get"),
		"/nf-instances",
		HTTPSearchNFInstances,
	},
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nnrf-nfm/v1")

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
		"DeregisterNFInstance",
		strings.ToUpper("Delete"),
		"/nf-instances/:nfInstanceID",
		HTTPDeregisterNFInstance,
	},

	{
		"GetNFInstance",
		strings.ToUpper("Get"),
		"/nf-instances/:nfInstanceID",
		HTTPGetNFInstance,
	},

	{
		"RegisterNFInstance",
		strings.ToUpper("Put"),
		"/nf-instances/:nfInstanceID",
		HTTPRegisterNFInstance,
	},

	{
		"UpdateNFInstance",
		strings.ToUpper("Patch"),
		"/nf-instances/:nfInstanceID",
		HTTPUpdateNFInstance,
	},

	{
		"GetNFInstances",
		strings.ToUpper("Get"),
		"/nf-instances",
		HTTPGetNFInstances,
	},
	/*
		{
			"RemoveSubscription",
			strings.ToUpper("Delete"),
			"/subscriptions/:subscriptionID",
			HTTPRemoveSubscription,
		},

		{
			"UpdateSubscription",
			strings.ToUpper("Patch"),
			"/subscriptions/:subscriptionID",
			HTTPUpdateSubscription,
		},

		{
			"CreateSubscription",
			strings.ToUpper("Post"),
			"/subscriptions",
			HTTPCreateSubscription,
		},
	*/
}
