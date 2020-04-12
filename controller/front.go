package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterController() *userController {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	return uc
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
