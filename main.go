package main

import (
	"net/http"

	"github.com/emersongonzal86/go-gorm-restapi/db"
	"github.com/emersongonzal86/go-gorm-restapi/models"
	"github.com/emersongonzal86/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBconnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
