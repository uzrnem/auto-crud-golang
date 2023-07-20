package routes

import (
	"net/http"
	"strings"

	"autocrud/pkg/config"
	"autocrud/pkg/responder"
	"github.com/gorilla/mux"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

func getRouteFor(extention string, path ...string) string {
	if extention == "" || extention == "/" {
		return strings.Join(path, "/")
	}
	return extention + strings.Join(path, "")
}

func getSpecificHandler(method string) func(w http.ResponseWriter, req *http.Request) {
	if method == "POST" {
		return postHandler
	} else if method == "PUT" {
		return putHandler
	} else if method == "DELETE" {
		return deleteHandler
	} else {
		return getHandler
	}
}

func Load(r *mux.Router) error {
	RegisteredRts = &RegisteredRoutes{
		routes: map[string]map[string]RouteDetails{},
	}
	for docName, doc := range config.Config.Application.Documents {
		for _, api := range doc.Apis {
			apiPath := api.Path
			if api.Path == "/" {
				apiPath = ""
			}
			route := getRouteFor(config.Config.Application.PathExtention, doc.Path, apiPath)

			r.HandleFunc(route, getSpecificHandler(api.Method)).Methods(api.Method).Name(route)
			RegisteredRts.SetRouteDetails(api.Method, route, docName, api)
		}
	}
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		responder.Resp.SendJson(w, 404, map[string]string{"error": "Route not found in config"})
	})
	return nil
}
