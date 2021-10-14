package view

import (
  "net/http"
  "CRUD/models"
  "html/template"
)

// func renderTemplate(temp string, data {}interface) error {
//
// }

func HomeView(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/home.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func CreateViewGET(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/create.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func CreateViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  name := r.FormValue("name")
  password := r.FormValue("password")
  models.CreateUser(name,email,password)
    //http.Redirect(w,r,"http://localhost:3000/",200)
}

func ReadView(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/read.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func DisplayUserView(w http.ResponseWriter,r *http.Request) {
  email := r.FormValue("email")
  foundUser, err := models.LookUpByEmail(email)
  if err != nil {
    userNotFoundView(w,r,email)
  } else {
    temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/displayUser.gohtml")
    if err != nil {
      panic(err)
    }
    temp.Execute(w,foundUser)
}
}

func UpdateViewGET(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/Update.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func UpdateViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  password := r.FormValue("password")
  models.UpdateUser(email,password)
}

func DeleteViewGET(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/delete.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func DeleteViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  models.DeleteUser(email)
}

func userNotFoundView(w http.ResponseWriter,r *http.Request, email string) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/userNotFound.gohtml")
  if err != nil {
    panic(err)
  }
  data := struct {
    Email string
  }{
    Email: email,
  }
  temp.Execute(w,data)
}
