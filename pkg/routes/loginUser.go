package routes

import (
	"github.com/gorilla/mux"
	"github.com/joko345/goBook/pkg/control"
)

var RegisterLoginRoute = func(router *mux.Router) { //route untuk controler gunakan
	router.HandleFunc("/login", control.LoginUser).Methods("POST")
}
