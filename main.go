package main

import (
	"log"
	"net/http"

	"github.com/Adictes/food-chooser/handlers"
	"github.com/Adictes/food-chooser/middleware"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.ServeFiles("/vendor/bootstrap/*filepath", http.Dir("vendor/bootstrap/"))

	router.GET("/", middleware.AccessLog(handlers.Index))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
