package handlers

import "net/http"

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// parsing data
	//validation data
	// create model
	err := service.CreatePost(post)
	if err != nil {
		// handlers err
	}

	// ...
}