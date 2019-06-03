package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = index(router)
	router = visit(router)

	return router

}
