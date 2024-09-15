package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Getting all posts and creating a new one
func HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		posts, err := getAllPosts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Display via template
		renderTemplate(w, "index", posts)
	
	case "POST":
		var post Post
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = createPost(post.Title, post.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Post created")
	}
}

// Get update and delete post by id
func HandlePostDetails(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/post/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		post, err := getPostByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(post)
		
	case "PUT":
		var post Post
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = updatePost(id, post.Title, post.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Post update")

	case "DELETE":
		err = deletePost(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Post deleted")
	}
}

// Database functions
func getAllPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := eows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return post, nil
}

func createPost(title, Content string) error {
	_, err := db.Exec("INSERT INTO posts (title, content) VALUS (?, ?)", title, content)
	return err
}

func updatePost(id int, title, content string) error {
	_, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	return err
}

func daletePost(id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

// Display html template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := fmt.Sprintf("./static/tmpl/%s.html", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
