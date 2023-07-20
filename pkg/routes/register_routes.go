package routes

import (
	"autocrud/pkg/config"
)

type RouteDetails struct {
	Document string     `json:"document"`
	Path     string     `json:"path"`
	Api      config.Api `json:"api"`
}

type RegisteredRoutes struct {
	routes map[string]map[string]RouteDetails
}

var (
	RegisteredRts *RegisteredRoutes
)

func (r *RegisteredRoutes) SetRouteDetails(method, route, document string, api config.Api) {
	if r.routes[method] == nil {
		r.routes[method] = map[string]RouteDetails{}
	}
	r.routes[method][route] = RouteDetails{
		Document: document,
		Path:     route,
		Api:      api,
	}
}

func (r *RegisteredRoutes) getRouteDetails(method, route string) (RouteDetails, error) {
	return r.routes[method][route], nil
}
