package routes

import (
	"net/http"

	"autocrud/pkg/config"
	"autocrud/pkg/db"
	"autocrud/pkg/responder"

	"github.com/gorilla/mux"
)

func putHandler(w http.ResponseWriter, req *http.Request) {
	routeName := mux.CurrentRoute(req).GetName()
	routeDetails, _ := RegisteredRts.getRouteDetails(PUT, routeName)

	id := config.Config.Application.Documents[routeDetails.Document].ID

	data, err := db.DB.GetDocumentById(routeDetails.Document, mux.Vars(req)[id.Name])
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	}
	err = getUpdateData(data, req.Body, routeDetails.Document, false)
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	}
	data[id.Name] = mux.Vars(req)[id.Name]
	err = db.DB.SaveDocument(routeDetails.Document, data)
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	} else {
		responder.Resp.SendJson(w, routeDetails.Api.Response.StatusCode, data)
		return
	}
}
