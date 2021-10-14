package models

import (
  _ "github.com/lib/pq"
	"database/sql"
  "fmt"
  "os"
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

func LookUpByEmail(email string) (User,error) {
  db := Connect()
  defer db.Close()
  row := db.QueryRow("SELECT * FROM USERDATA WHERE email = $1", email)
  if row.Err() != nil {
    return User{},row.Err()
  }
  user := new(User)
  err := row.Scan(&user.Userid,&user.Name,&user.Email,&user.Password)
  if err != nil {
    return User{},err
  }
  fmt.Fprintf(os.Stdout,user.Name)
  return *user, nil
}

func CreateUser(name string, email string, password string) {
  db := Connect()
  defer db.Close()
  _, err := db.Exec("INSERT INTO USERDATA (name,email,password) VALUES($1,$2,$3);", name, email, password)
	if err != nil {
		panic(err)
	}
  db.Close()
}

func UpdateUser(email string, password string) {
  db := Connect()
  defer db.Close()
  _, err := db.Exec("UPDATE USERDATA SET password = $1 WHERE email = $2", password,email)
	if err != nil {
		panic(err)
}
}

func DeleteUser(email string) {
  db := Connect()
  defer db.Close()
  _, err := db.Exec("DELETE FROM USERDATA WHERE email = $1;", email)
	if err != nil {
		panic(err)
}
}

func Connect() *sql.DB {
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
