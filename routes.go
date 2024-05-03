package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"Health", "GET", "/health", Health},
		Route{"Metrics", "GET", "/metrics", Metrics(promhttp.Handler())},
	}
	return routes
}
