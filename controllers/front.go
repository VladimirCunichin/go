package controllers

import (
	"net/http"
	"encoding/json"
	"io"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
