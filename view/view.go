package view

import (
  "net/http"
  "CRUD/models"
  "html/template"
)

func renderTemplate(w http.ResponseWriter,filePath string, dataToRender interface{}) {
  temp, err := template.ParseFiles(filePath)
  if err != nil {
    panic(err)
  }
  temp.Execute(w,dataToRender)
}

func HomeView(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/home.gohtml",nil)
}

func CreateViewGET(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/create.gohtml",nil)
}

func CreateViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  name := r.FormValue("name")
  password := r.FormValue("password")
  err = models.CreateUser(name,email,password)
  if err != nil {
    renderTemplate(w,"templates/userAlreadyExists.gohtml",struct{Email string}{Email: email,})
  } else {
    renderTemplate(w,"templates/success.gohtml",nil)
  }
}

func ReadView(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/read.gohtml",nil)
}

func DisplayUserView(w http.ResponseWriter,r *http.Request) {
  email := r.FormValue("email")
  user, err := models.LookUpByEmail(email)
  if err != nil {
    renderTemplate(w,"templates/userNotFound.gohtml",struct{Email string}{Email: email,})
  } else {
    renderTemplate(w,"templates/displayUser.gohtml",user)
}
}

func UpdateViewGET(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/Update.gohtml",nil)
}

func UpdateViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  password := r.FormValue("password")
  err = models.UpdateUser(email,password)
  if err != nil {
    renderTemplate(w,"templates/userNotFound.gohtml",struct{Email string}{Email: email,})
  } else {
    renderTemplate(w,"templates/success.gohtml",nil)
  }
}

func DeleteViewGET(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/delete.gohtml",nil)
}

func DeleteViewPOST(w http.ResponseWriter,r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  err = models.DeleteUser(email)
  if err != nil {
    renderTemplate(w,"templates/userNotFound.gohtml",struct{Email string}{Email: email,})
  } else{
  renderTemplate(w,"templates/success.gohtml",nil)
}
}
