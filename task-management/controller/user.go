package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"task/model"
	"task/service"

)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser model.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call
	user, err := service.NewUser(newUser.ID, newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// update name
func UpdateUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var updateUser model.User

	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// service call
	updatedUser, err := service.UpdateUserName(&updateUser, updateUser.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var deleteUser model.User

	err := json.NewDecoder(r.Body).Decode(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call
	deleteduser, err := service.DeleteUser(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(deleteduser)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// var getUser model.User

	// err := json.NewDecoder(r.Body).Decode(&getUser)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	users, err := service.GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("it has reached controller")

	vars := mux.Vars(r)
	userIDStr := vars["ID"]
	log.Printf("Received request for user ID: %s", userIDStr)

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserById(uint(userID))
	if err != nil {
		log.Printf("Error fetching user from service: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Printf("Successfully fetched user: %+v", user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("entered delete by id")
	vars := mux.Vars(r)
	userIDStr := vars["ID"]

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteUserID(uint(userID))
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
