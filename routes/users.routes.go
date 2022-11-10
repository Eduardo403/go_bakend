package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Eduardo403/go_bakend/db"
	"github.com/Eduardo403/go_bakend/models"
	"github.com/gorilla/mux"
)
// get all users 
func GetUsersHandel(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
     w.Write([]byte("get users"))
}
//get onli users 
func GetUserHandel(w http.ResponseWriter, r *http.Request) {
	var users models.User
	parans:= mux.Vars(r)
	db.DB.First(&users, parans["id"] )

	if users.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not fout"))
		return
	}

	db.DB.Model(&users).Association("Taks").Find(&users)
json.NewEncoder(w).Encode(parans)
}
//add new user
func PostUsersHandel(w http.ResponseWriter, r *http.Request) {
var user models.User
json.NewDecoder(r.Body).Decode(&user)
createUser :=  db.DB.Create(&user)

err:= createUser.Error
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
json.NewEncoder(w).Encode(&user)
}
//update onli user
func PathUsersHandel(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("update users"))
}

//delete user
func DeleteUsersHandel(w http.ResponseWriter, r *http.Request) {
	var users models.User
	parans:= mux.Vars(r)
	db.DB.First(&users, parans["id"] )

	if users.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not fout"))
		return
	}
	db.DB.Delete(&users)
w.WriteHeader(http.StatusOK)
}