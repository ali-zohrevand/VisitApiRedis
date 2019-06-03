package routers

import (
	"github.com/ali-zohrevand/VisitApi/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func index(router *mux.Router) *mux.Router {
	HandleFunc := negroni.New(
		negroni.HandlerFunc(controllers.Index),
	)
	UpFunc := negroni.New(
		negroni.HandlerFunc(controllers.Status),
	)
	router.Handle("/", HandleFunc).Methods("GET")
	router.Handle("/status", UpFunc).Methods("GET")

	return router
}
