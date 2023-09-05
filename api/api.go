package api

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/PyMarcus/distribuidos/controllers"
	"github.com/PyMarcus/distribuidos/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r)
	if r.Method == "GET" {
		log.Println("Get method")
		handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		log.Println("Post method")
		handleCreateUser(w, r)
	}
	if r.Method == "PUT" {
		log.Println("Put method")
		handleUpdateUser(w, r)
	}
	if r.Method == "DELETE" {
		log.Println("delete method")
		handleDeleteUser(w, r)
	}

}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := controllers.GetAllUsers()
	for _, u := range users {
		log.Println(u)
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Server ERROR", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	controllers.CreateUser(name, email)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateData models.UpdateData
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Fail to encoding JSON", http.StatusBadRequest)
		return
	}
	controllers.UpdateUser(updateData.Name, updateData.Email, updateData.OldName, updateData.OldEmail)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var delete models.Delete
	err := json.NewDecoder(r.Body).Decode(&delete)
	if err != nil {
		http.Error(w, "Fail to encoding JSON", http.StatusBadRequest)
		return
	}
	controllers.DeleteUser(delete.Name, delete.Email)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func routes() {
	http.HandleFunc("/", handler)
}

func Run() {
	log.Println("Running on localhost:9998")
	routes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatalln(http.ListenAndServe(":9998", nil))
}
