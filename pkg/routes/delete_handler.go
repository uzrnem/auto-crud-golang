package routes

import (
	"net/http"

	"autocrud/pkg/config"
	"autocrud/pkg/db"
	"autocrud/pkg/responder"

	"github.com/gorilla/mux"
)

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	routeName := mux.CurrentRoute(req).GetName()
	routeDetails, _ := RegisteredRts.getRouteDetails(DELETE, routeName)

	id := config.Config.Application.Documents[routeDetails.Document].ID

	err := db.DB.DeleteDocumentById(routeDetails.Document, mux.Vars(req)[id.Name])
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	} else {
		responder.Resp.SendJson(w, routeDetails.Api.Response.StatusCode, nil)
		return
	}
}
