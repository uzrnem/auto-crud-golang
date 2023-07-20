package routes

import (
	"net/http"

	"autocrud/pkg/config"
	"autocrud/pkg/db"
	"autocrud/pkg/integer"
	"autocrud/pkg/responder"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, req *http.Request) {
	routeName := mux.CurrentRoute(req).GetName()
	routeDetails, _ := RegisteredRts.getRouteDetails(POST, routeName)

	id := config.Config.Application.Documents[routeDetails.Document].ID

	data := map[string]any{}
	err := getUpdateData(data, req.Body, routeDetails.Document, true)
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	}
	if id.Type == "uuid" {
		data[id.Name] = uuid.New()
	} else {
		data[id.Name] = integer.GetCounter()
	}
	err = db.DB.SaveDocument(routeDetails.Document, data)
	if err != nil {
		responder.Resp.SendError(w, err)
		return
	} else {
		responder.Resp.SendJson(w, routeDetails.Api.Response.StatusCode, data)
		return
	}
}
