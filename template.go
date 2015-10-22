package main

import (
	"net/http"
	"text/template"
)

var tpls = map[string]*template.Template{
	"404":   newTemplate("views/layout.html", "views/404.html"),
	"index": newTemplate("views/layout.html", "views/index.html"),
}

func newTemplate(files ...string) *template.Template {
	return template.Must(template.New("*").ParseFiles(files...))
}

func executeTemplate(w http.ResponseWriter, name string, status int, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	return tpls[name].ExecuteTemplate(w, "layout", data)
}
