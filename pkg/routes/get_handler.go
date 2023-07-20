package routes

import (
	"net/http"

	"autocrud/pkg/config"
	"autocrud/pkg/db"
	"autocrud/pkg/responder"

	"github.com/gorilla/mux"
)

func getHandler(w http.ResponseWriter, req *http.Request) {
	routeName := mux.CurrentRoute(req).GetName()
	routeDetails, _ := RegisteredRts.getRouteDetails(GET, routeName)
	if routeDetails.Api.Type == "" || routeDetails.Api.Type == "STATIC" {
		responder.Resp.SendJson(w,
			routeDetails.Api.Response.StatusCode,
			routeDetails.Api.Response.Body,
		)
	} else if routeDetails.Api.Type == "LIST" {
		data, err := db.DB.GetDocuments(routeDetails.Document)
		if err != nil {
			responder.Resp.SendError(w, err)
			return
		} else {
			list := []map[string]any{}
			for _, object := range data {
				list = append(list, object)
			}
			responder.Resp.SendJson(w, routeDetails.Api.Response.StatusCode, list)
			return
		}
	} else {

		id := config.Config.Application.Documents[routeDetails.Document].ID

		data, err := db.DB.GetDocumentById(routeDetails.Document, mux.Vars(req)[id.Name])
		if err != nil {
			responder.Resp.SendError(w, err)
			return
		} else {
			responder.Resp.SendJson(w, routeDetails.Api.Response.StatusCode, data)
			return
		}
	}
}
