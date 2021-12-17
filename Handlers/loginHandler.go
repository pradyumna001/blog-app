package handlers

import (
	"net/http"

	controllers "github.com/pradyumna001/blog-app/Controllers"
)

//Handle Login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var uc *controllers.UsersController
		auth := uc.AuthenticateUser(w, r)
		if auth {
			uc.ReadRecord(w, r)
		}

	}
}
