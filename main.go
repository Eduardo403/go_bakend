package main

import (
	"net/http"

	"github.com/Eduardo403/go_bakend/db"
	"github.com/Eduardo403/go_bakend/models"
	"github.com/Eduardo403/go_bakend/routes"
	"github.com/gorilla/mux"
)


func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Taks{})
	db.DB.AutoMigrate(models.User{})
r:=	mux.NewRouter();

//users routes 
r.HandleFunc("/",routes.HomeHandele)
r.HandleFunc("/users",routes.GetUsersHandel).Methods("GET")
r.HandleFunc("/users/{id}",routes.GetUserHandel).Methods("GET")
r.HandleFunc("/users",routes.PostUsersHandel).Methods("POST")
r.HandleFunc("/users/{id}",routes.PathUsersHandel).Methods("PATCH")
r.HandleFunc("/users/{id}",routes.DeleteUsersHandel).Methods("DELETE")

//tasks routes 

r.HandleFunc("/tasks",routes.GetTasksHandel).Methods("GET")
r.HandleFunc("/tasks/{id}",routes.GetTaskHandel).Methods("GET")
r.HandleFunc("/tasks",routes.PostTaksHandel).Methods("POST")
r.HandleFunc("/tasks/{id}",routes.DeleteTaksHandel).Methods("DELETE")
http.ListenAndServe(":3000",r)
}
