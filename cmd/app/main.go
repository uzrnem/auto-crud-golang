package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"autocrud/pkg/config"
	"autocrud/pkg/db"
	"autocrud/pkg/responder"
	"autocrud/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {

	err := db.Load()
	if err != nil {
		fmt.Printf("db load error: %+v\n", err)
		return
	}

	err = config.Load()
	if err != nil {
		fmt.Printf("config load error: %+v\n", err)
		return
	}

	err = responder.Load()
	if err != nil {
		fmt.Printf("responder load error: %+v\n", err)
		return
	}

	r := mux.NewRouter()

	err = routes.Load(r)
	if err != nil {
		fmt.Printf("router load error: %+v\n", err)
		return
	}
	r.Use(TrimRightSlash)
	r.Use(loggerRequest)

	port := fmt.Sprintf(":%d", config.Config.Application.Port)

	fmt.Printf("Server Started: %s\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}

func TrimRightSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.TrimRight(req.URL.Path, "/")
		next.ServeHTTP(res, req)
	})
}
func loggerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
