package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

// using asymmetric crypto/RSA keys
// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

func GetSignKey() []byte {
	var err error
	var signKey []byte
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pKeyPath := path.Join(dir, "../src/widgets-api/"+privKeyPath)

	signKey, err = ioutil.ReadFile(pKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return nil
	}
	return signKey
}

func GetVerifyKey() []byte {
	var err error
	var verifyKey []byte
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	publicKeyPath := path.Join(dir, "../src/widgets-api/"+pubKeyPath)

	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
	}
	return verifyKey
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}

	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
