package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joko345/goBook/pkg/routes"
	"github.com/rs/cors"
)

func main() {
	// r := mux.NewRouter()
	// routes.RegisterBookRoutes(r) //r akan digunakan untuk var router
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:8080", r)) //jalan di 3306

	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	handler := cors.AllowAll().Handler(r) // Middleware CORS untuk mengizinkan semua permintaan
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
