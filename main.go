package main

import (
	"log"
	"net/http"

	"github.com/elasticsearch_go_impl/middleware"
	"github.com/elasticsearch_go_impl/src"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /api/create/{index}", middleware.ResponseHandler(src.CreateIndex))
	router.HandleFunc("POST /api/insertMany/{index}", middleware.ResponseHandler(src.InsertMany))
	router.HandleFunc("POST /api/insertOne/{index}", middleware.ResponseHandler(src.InsertOne))
	router.HandleFunc("POST /api/search/{index}", middleware.ResponseHandler(src.Search))
	router.HandleFunc("PUT /api/update/{index}", middleware.ResponseHandler(src.UpdateOne))
	router.HandleFunc("DELETE /api/delete/{index}", middleware.ResponseHandler(src.DeleteOne))

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	log.Fatal(server.ListenAndServe())
}
