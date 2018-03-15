package controller

import (
	"html/template"
	"net/http"
	"path"
	"widgets-api/app/util"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := path.Join(util.GetRootPath(), "../src/widgets-api/template/index.html")
	t, _ := template.ParseFiles(path)
	t.Execute(w, nil)
}

func Users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := path.Join(util.GetRootPath(), "../src/widgets-api/template/user.html")
	t, _ := template.ParseFiles(path)
	t.Execute(w, nil)
}

func Widgets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := path.Join(util.GetRootPath(), "../src/widgets-api/template/widget.html")
	t, _ := template.ParseFiles(path)
	t.Execute(w, nil)
}
