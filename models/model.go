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

type Page struct {
  Title string
  Author string
  Article string
}

func LookUpByTitle(title string) (Page,error) {
  db := Connect()
  defer db.Close()
  row := db.QueryRow("SELECT * FROM Pages WHERE Title = $1", title)
  if row.Err() != nil {
    return Page{},row.Err()
  }
  page := new(Page)
  err := row.Scan(&page.Title,&page.Author,&page.Article)
  if err != nil {
    return Page{},err
  }
  return *page, nil
}

func CreatePage(title string, author string, article string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("SELECT * FROM Pages WHERE Title = $1;",title)
  rowsaffected, err := result.RowsAffected()
  if rowsaffected > 0 {
    return errors.New("Page already exits")
  }
  if err != nil {
    panic(err)
  }
  _, err = db.Exec("INSERT INTO Pages (Title,Author,Article) VALUES($1,$2,$3);", title,author,article)
  if err != nil {
    panic(err)
  }
	return nil
}

func UpdatePage(title string, article string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("UPDATE Pages SET Article = $1 WHERE title = $2", article,title)
  if err != nil {
    panic(err)
  }
  num, _ := result.RowsAffected()
  if num == 0 {
    return errors.New("No rows were affected, page does not exit")
  }
  return nil
}

func DeletePage(title string) error {
  db := Connect()
  defer db.Close()
  result, err := db.Exec("DELETE FROM Pages WHERE title = $1;", title)
  if err != nil {
    panic(err)
  }
  num, _ := result.RowsAffected()
  if num == 0 {
    return errors.New("No rows were affected, page does not exit")
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
