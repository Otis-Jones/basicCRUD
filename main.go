package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "CRUD/view"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/",view.HomeView)
  r.HandleFunc("/Create",view.CreateView)
  r.HandleFunc("/Read",view.ReadView)
  r.HandleFunc("/displayUser",view.DisplayUserView)
  r.HandleFunc("/Update",view.UpdateView)
  r.HandleFunc("/Delete",view.DeleteView)
  http.ListenAndServe("localhost:3000",r)
}
