package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Adictes/food-chooser/handlers"
	"github.com/Adictes/food-chooser/middleware"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/", middleware.AccessLog(handlers.Index))
	router.GET("/frws", handlers.FoodRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
