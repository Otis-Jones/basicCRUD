package main

import (
  "net/http"
  "html/template"
  "fmt"
  "os"
  _ "github.com/lib/pq"
	"database/sql"
  "github.com/gorilla/mux"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = ""
  dbname   = "CRUDdb"
)

type User struct {
  Userid int
  Name string
  Email string
  Password string
}

func homeView(w http.ResponseWriter,r *http.Request) {
  temp, err := template.ParseFiles("/Users/otisjones/Desktop/CRUD/templates/home.gohtml")
  if err != nil {
    panic(err)
  }
  temp.Execute(w,nil)
}

func createUser(name string, email string, password string) {
  db := connect()
  defer db.Close()
  _, err := db.Exec("INSERT INTO USERDATA (name,email,password) VALUES($1,$2,$3);", name, email, password)
	if err != nil {
		panic(err)
	}
  db.Close()
}

func createView(w http.ResponseWriter,r *http.Request) {
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

    createUser(name,email,password)
    //http.Redirect(w,r,"http://localhost:3000/",200)
  }
}

func readView(w http.ResponseWriter,r *http.Request) {
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



func lookUpByEmail(email string) User {
  db := connect()
  defer db.Close()
  row := db.QueryRow("SELECT * FROM USERDATA WHERE email = $1", email)
  if row.Err() != nil {
    fmt.Println("ERROR 1")
    panic(row.Err())
  }
  user := new(User)
  err := row.Scan(&user.Userid,&user.Name,&user.Email,&user.Password)
  if err != nil {
    fmt.Println("ERROR 2")
    panic(err)
  }
  fmt.Fprintf(os.Stdout,user.Name)
  return *user
}

func displayUserView(w http.ResponseWriter,r *http.Request) {

  email := r.FormValue("email")
  foundUser := lookUpByEmail(email)
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

func connect() *sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "dbname=%s sslmode=disable",
  host, port, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println(os.Stdout,"Succesfully connected")
  return db
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/",homeView)
  r.HandleFunc("/Create",createView)
  r.HandleFunc("/Read",readView)
  r.HandleFunc("/displayUser",displayUserView)
  http.ListenAndServe("localhost:3000",r)
}
