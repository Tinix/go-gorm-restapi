//
// users.routes.go
// Copyright (C) 2022 tinix <tinix@archlinux>
//
// Distributed under terms of the MIT license.
//

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tinix/go-gorm-restapi/db"
	"github.com/Tinix/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

// GET /users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// POST /users
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

// GET /users/1
func GetUserIdHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found."))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// DELETE /users/1
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found."))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
