package view

import (
  "net/http"
  "fmt"
  "CRUD/models"
  "html/template"
  "os"
)

func HomeView(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/home.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func CreateView(w http.ResponseWriter,r *http.Request) {
  switch r.Method {
  case "GET":
    temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/create.gohtml")
    if err != nil {
      panic(err)
    }
    temp.Execute(w,nil)
  case "POST":
    err := r.ParseForm()
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(os.Stdout, "Post from website! r.PostFrom = %v\n", r.PostForm)
    email := r.FormValue("email")
    name := r.FormValue("name")
    password := r.FormValue("password")
    fmt.Fprintln(os.Stdout, name)
    fmt.Fprintln(os.Stdout, email)
    fmt.Fprintln(os.Stdout, password)

    models.CreateUser(name,email,password)
    //http.Redirect(w,r,"http://localhost:3000/",200)
  }
}

func ReadView(w http.ResponseWriter,r *http.Request) {
  switch r.Method {
  case "GET":
    temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/read.gohtml")
    if err != nil {
      panic(err)
    }
    temp.Execute(w,nil)
  case "POST":
    fmt.Println("Should not be a POST")
  }
}

func DisplayUserView(w http.ResponseWriter,r *http.Request) {
  email := r.FormValue("email")
  foundUser := models.LookUpByEmail(email)
  fmt.Fprintf(os.Stdout,"FOUND USER")
  fmt.Fprintf(os.Stdout, "FOUND THE USER = %v\n", foundUser)
  fmt.Fprintf(os.Stdout,foundUser.Name)
  fmt.Fprintf(os.Stdout,foundUser.Email)
  fmt.Fprintf(os.Stdout,foundUser.Password)
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/displayUser.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,foundUser)
}

func UpdateView(w http.ResponseWriter,r *http.Request) {
fmt.Fprintf(os.Stdout,r.Method)
switch r.Method {
case "GET":
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/Update.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
case "POST":
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  email := r.FormValue("email")
  password := r.FormValue("password")
  fmt.Fprintf(os.Stdout,email)
  fmt.Fprintf(os.Stdout,password)
  models.UpdateUser(email,password)
}
}


func DeleteView(w http.ResponseWriter,r *http.Request) {
  switch r.Method {
  case "GET":
    temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/delete.gohtml")
    if err != nil {
      panic(err)
    }
    temp.Execute(w,nil)
  case "POST":
    err := r.ParseForm()
    if err != nil {
      panic(err)
    }
    email := r.FormValue("email")
    models.DeleteUser(email)
  }
}
