package controller

import (
	"net/http"
	"regexp"

	"github.com/leon0707/go_webservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
	w.Write([]byte("hello"))
}

func (uc *userController) getAll(w http.ResponseWriter) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	user, error := models.GetUserByID(id)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(user, w)
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
