package handlers

import (
	"net/http"

	controllers "github.com/pradyumna001/blog-app/Controllers"
)

//Handlers Routes For User
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var uc *controllers.UsersController
	switch r.Method {
	case http.MethodGet:
		uc.ReadRecords(w, r)
	case http.MethodPost:
		// Create a new record.
		uc.CreateRecord(w, r)
	case http.MethodPut:
		// Update an existing record.
		uc.UpdateRecord(w, r)
	case http.MethodDelete:
		// Remove the record.
		uc.DeleteRecord(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//Handlers Routes For User
func UserSearchHandler(w http.ResponseWriter, r *http.Request) {
	var uc *controllers.UsersController

	uc.SearchRecords(w, r)

}
