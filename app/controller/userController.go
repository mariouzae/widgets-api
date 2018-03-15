package controller

import (
	"encoding/json"
	"net/http"
	"widgets-api/app/auth"
	"widgets-api/app/model"
	"widgets-api/app/util"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		util.DisplayAppError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return
	}

	u, err := model.FindUserByName(user.Name)
	if err != nil {
		util.DisplayAppError(w, err, "Username or password is wrong", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		util.DisplayAppError(w, err, "Invalid login credentials", http.StatusUnauthorized)
		return
	}
	auth.LoginHandler(w, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := model.FindUsers()
	if err != nil {
		util.DisplayAppError(w, err, "error to retrieve users", http.StatusInternalServerError)
		return
	}
	util.JsonResponse(users, w)
}

func GetUserById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	user, err := model.FindUserById(id)
	if err != nil {
		util.DisplayAppError(w, err, "error to retrieve users", http.StatusInternalServerError)
		return
	}
	util.JsonResponse(user, w)
}
