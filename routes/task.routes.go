package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Eduardo403/go_bakend/db"
	"github.com/Eduardo403/go_bakend/models"
	"github.com/gorilla/mux"
)
//get all tasks 
func GetTasksHandel(w http.ResponseWriter , r *http.Request) {
var tasks []models.Taks
db.DB.Find(&tasks)
json.NewEncoder(w).Encode(&tasks)
}
//get onli tasks
func GetTaskHandel(w http.ResponseWriter , r *http.Request) {
var task models.Taks
parames:= mux.Vars(r)

db.DB.First(&task, parames["id"])
if task.ID==0 {
    w.WriteHeader(http.StatusBadRequest) //404
    w.Write([]byte("tasks not found"))
    return
}

json.NewEncoder(w).Encode(&task)
w.WriteHeader(http.StatusNoContent) //204
}

//insetr tasks
func PostTaksHandel(w http.ResponseWriter , r *http.Request) {
var task models.Taks
json.NewDecoder(r.Body).Decode(&task)
createTask := db.DB.Create(&task)
err:= createTask.Error
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
}
json.NewEncoder(w).Encode(&task)
}
//delete tasks 
func DeleteTaksHandel(w http.ResponseWriter , r *http.Request) {
var task models.Taks
json.NewDecoder(r.Body).Decode(&task)
createTask := db.DB.Create(&task)
err:= createTask.Error
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
}
db.DB.Delete(&task)
}