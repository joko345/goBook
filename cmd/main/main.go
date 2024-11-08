package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joko345/goBook/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r) //r akan digunakan untuk var router
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
