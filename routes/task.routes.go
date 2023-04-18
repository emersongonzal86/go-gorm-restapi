package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emersongonzal86/go-gorm-restapi/db"
	"github.com/emersongonzal86/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
	
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //error400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID ==0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)

}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID ==0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	//de la siguiente manera hace el softdelete del registro
	//db.DB.Delete(&user)
	//si quiero borrar el dato fisicamente de la base de datos uso lo siguiente
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) //204

}
