package controller

import (
	"encoding/json"
	"net/http"
	"widgets-api/app/model"
	"widgets-api/app/util"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func GetWidgets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	widgets, err := model.FindWidgets()
	if err != nil {
		util.DisplayAppError(w, err, "error to retrieve widgets", http.StatusInternalServerError)
		return
	}
	util.JsonResponse(widgets, w)
}

func GetWidgetById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	widget, err := model.FindWidgetById(id)
	if err != nil {
		util.DisplayAppError(w, err, "error to get widget", http.StatusInternalServerError)
		return
	}
	util.JsonResponse(widget, w)
}

func CreateWidget(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var wi model.Widget

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&wi)
	if err != nil {
		panic(err)
	}

	widget, err := model.CreateWidget(wi)
	if err != nil {
		util.DisplayAppError(w, err, "error to create new widget", http.StatusInternalServerError)
		return
	}
	util.JsonResponse(widget, w)
}

func UpdateWidget(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	objID := param.ByName("id")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var wi model.Widget
	wi.ID = bson.ObjectIdHex(objID)
	err := decoder.Decode(&wi)
	if err != nil {
		panic(err)
	}
	err = model.UpdateWidget(wi)
	if err != nil {
		util.DisplayAppError(w, err, "error to update a widget", http.StatusInternalServerError)
		return
	}
}
