package controllers

import (
	"github.com/ali-zohrevand/VisitApi/DB/cash"
	"net/http"
)

func Visit(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	cash.VisitHappened()
	w.Write([]byte("Hello, World!"))
}
