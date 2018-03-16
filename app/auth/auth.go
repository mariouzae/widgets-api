package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"widgets-api/app/model"
	"widgets-api/app/util"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Text string `json:"text"`
}

type Token struct {
	Token string `json:"token"`
}

func AuthHandler(handleFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// validate the token
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return util.GetSignKey(), nil
				})
				if token.Valid {
					fmt.Println("Valid token: " + bearerToken[1])
					handleFunc(w, r, params)
				} else {
					util.DisplayAppError(w, err, "Authentication failed", http.StatusUnauthorized)
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Printf("%s", err)
		//return err
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"ext":  time.Now().Add(time.Minute * 20160).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(util.GetSignKey())
	if err != nil {
		//return err
	}
	response := Token{tokenString}
	util.JsonResponse(response, w)
}
