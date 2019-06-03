package main

import (
	"fmt"
	"github.com/ali-zohrevand/VisitApi/routers"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	router := routers.InitRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{"GET", "UPDATE", "PUT", "POST", "DELETE"},
		Debug:          true,
	})
	n := negroni.Classic()

	n.Use(c)
	n.UseHandler(router)
	fmt.Println("start server")

	err := http.ListenAndServe(":5000", n)
	fmt.Println(err)

}
