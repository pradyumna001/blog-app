package handlers

import (
	"net/http"

	controllers "github.com/pradyumna001/blog-app/Controllers"
)

//Handlers Routes For Posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	var pc *controllers.PostsController
	switch r.Method {
	case http.MethodGet:
		// Get All Posts.
		pc.ReadRecords(w, r)
	case http.MethodPost:
		// Create a new Post.
		pc.CreateRecord(w, r)
	case http.MethodPut:
		// Update an existing Post.
		pc.UpdateRecord(w, r)
	case http.MethodDelete:
		// Remove the Post.
		pc.DeleteRecord(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
