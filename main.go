package main

import (
	"log"
	"net/http"

	"github.com/elasticsearch_go_impl/middleware"
	"github.com/elasticsearch_go_impl/src"
)

// main is the entry point of the application
func main() {
    // create a new http router
    router := http.NewServeMux()

    // register the API routes
    router.HandleFunc("POST /api/create/{index}", middleware.ResponseHandler(src.CreateIndex))
    router.HandleFunc("POST /api/insertMany/{index}", middleware.ResponseHandler(src.InsertMany))
    router.HandleFunc("POST /api/insertOne/{index}", middleware.ResponseHandler(src.InsertOne))
    router.HandleFunc("POST /api/search/{index}", middleware.ResponseHandler(src.Search))
    router.HandleFunc("PUT /api/update/{index}", middleware.ResponseHandler(src.UpdateOne))
    router.HandleFunc("DELETE /api/delete/{index}", middleware.ResponseHandler(src.DeleteOne))

    // create a new server and start listening
    server := &http.Server{
        Addr:    ":8080",
        Handler: middleware.Logger(router),
    }
    log.Fatal(server.ListenAndServe())
}
