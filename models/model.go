package models

import (
  _ "github.com/lib/pq"
	"database/sql"
  "fmt"
  "errors"
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
  return *user, nil
}

func CreateUser(name string, email string, password string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("SELECT * FROM USERDATA WHERE email = $1;",email)
  rowsaffected, err := result.RowsAffected()
  if rowsaffected > 0 {
    return errors.New("User already exits")
  }
  if err != nil {
    panic(err)
  }
  _, err = db.Exec("INSERT INTO USERDATA (name,email,password) VALUES($1,$2,$3);", name, email, password)
  if err != nil {
    panic(err)
  }
	return nil
}

func UpdateUser(email string, password string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("UPDATE USERDATA SET password = $1 WHERE email = $2", password,email)
  if err != nil {
    panic(err)
  }
  num, _ := result.RowsAffected()
  if num == 0 {
    return errors.New("No rows were affected, user does not exit")
  }
  return nil
}

func DeleteUser(email string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("DELETE FROM USERDATA WHERE email = $1;", email)
  if err != nil {
    panic(err)
  }
  num, _ := result.RowsAffected()
  if num == 0 {
    return errors.New("No rows were affected, user does not exit")
  }
  return nil
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
  //fmt.Println(os.Stdout,"Succesfully connected")
  return db
}
