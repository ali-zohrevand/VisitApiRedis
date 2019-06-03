package routers

import (
	"github.com/ali-zohrevand/VisitApi/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func visit(router *mux.Router) *mux.Router {
	HandleFunc := negroni.New(
		negroni.HandlerFunc(controllers.Visit),
	)
	router.Handle("/visit", HandleFunc).Methods("GET")

	return router
}
