package controller

import (
	"encoding/json"
	"net/http"

	"task/model"
	"task/service"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask model.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call
	task, err := service.NewTask(newTask.ID, newTask.UserID, newTask.DueDate, newTask.Title, newTask.Description, newTask.Priority, newTask.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var updateTask model.Task

	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call
	err = service.UpdateTaskTitle(&updateTask, updateTask.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(updateTask)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var deleteTask model.Task

	err := json.NewDecoder(r.Body).Decode(&deleteTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call
	err = service.DeleteTask(&deleteTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(deleteTask)
}

func GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	var getalltask []*model.Task

	err := json.NewDecoder(r.Body).Decode(&getalltask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alltask, err := service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(alltask)
}
