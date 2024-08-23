package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"task/controller"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/getUser/{ID}", controller.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/deletedUser/{ID}", controller.DeleteUserIDHandler).Methods("DELETE")

	r.HandleFunc("/create-user", controller.CreateUserHandler).Methods("POST")
	r.HandleFunc("/update-user", controller.UpdateUsernameHandler).Methods("PUT")
	r.HandleFunc("/delete-user", controller.DeleteUserHandler).Methods("DELETE")
	http.HandleFunc("/getUsers", controller.GetUserHandler)

	http.HandleFunc("/create-task", controller.CreateTaskHandler)
	http.HandleFunc("/update-task", controller.UpdateTaskHandler)
	http.HandleFunc("/delete-task", controller.DeleteTaskHandler)
	http.HandleFunc("/get-task", controller.GetAllTaskHandler)
}
