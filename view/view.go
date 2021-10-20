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
  title := r.FormValue("title")
  author := r.FormValue("author")
  article := r.FormValue("article")
  err = models.CreatePage(title,author,article)
  if err != nil {
    renderTemplate(w,"templates/pageAlreadyExists.gohtml",struct{Title string}{Title: title,})
  } else {
    renderTemplate(w,"templates/success.gohtml",nil)
  }
}

func ReadView(w http.ResponseWriter,r *http.Request) {
  renderTemplate(w,"templates/read.gohtml",nil)
}

func DisplayPageView(w http.ResponseWriter,r *http.Request) {
  title := r.FormValue("title")
  page, err := models.LookUpByTitle(title)
  if err != nil {
    renderTemplate(w,"templates/pageNotFound.gohtml",struct{Title string}{Title: title,})
  } else {
    renderTemplate(w,"templates/displayPage.gohtml",page)
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
  title := r.FormValue("title")
  article := r.FormValue("article")
  err = models.UpdatePage(title,article)
  if err != nil {
    renderTemplate(w,"templates/pageNotFound.gohtml",struct{Title string}{Title: title,})
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
  title := r.FormValue("title")
  err = models.DeletePage(title)
  if err != nil {
    renderTemplate(w,"templates/pageNotFound.gohtml",struct{Title string}{Title: title,})
  } else{
  renderTemplate(w,"templates/success.gohtml",nil)
}
}
