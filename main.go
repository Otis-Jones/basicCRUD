package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "CRUD/view"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/",view.HomeView)
  r.HandleFunc("/Create",view.CreateViewGET).Methods("GET")
  r.HandleFunc("/Create",view.CreateViewPOST).Methods("POST")
  r.HandleFunc("/Read",view.ReadView)
  r.HandleFunc("/displayPage",view.DisplayPageView)
  r.HandleFunc("/Update",view.UpdateViewGET).Methods("GET")
  r.HandleFunc("/Update",view.UpdateViewPOST).Methods("POST")
  r.HandleFunc("/Delete",view.DeleteViewGET).Methods("GET")
  r.HandleFunc("/Delete",view.DeleteViewPOST).Methods("POST")
  http.ListenAndServe("localhost:3000",r)
}
