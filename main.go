package main

import (
	"net/http"

	"github.com/Tinix/go-gorm-restapi/db"
	"github.com/Tinix/go-gorm-restapi/models"
	"github.com/Tinix/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserIdHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
