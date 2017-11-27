package routers

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetTaskRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/tasks", controllers.Tasks.Create).Methods("POST")
	router.HandleFunc("/tasks", controllers.Tasks.Get).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Show).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Update).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Delete).Methods("DELETE")

	return router
}
