package routes

import (
	"github.com/gorilla/mux"
	"github.com/joko345/goBook/pkg/control"
)

var RegisterBookRoutes = func(router *mux.Router) { //route untuk controler gunakan
	router.HandleFunc("/book/", control.CreateBook).Methods("POST")
	router.HandleFunc("/book/", control.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", control.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", control.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", control.DeleteBook).Methods("DELETE")
}
