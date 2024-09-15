package main

import (
	"log"
	"net/http"
	"blog-post/internal"
)

func main() {
	// database initialization
	db, err = internal.InitDB("./blog.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Application of migration
	err = internal.ApplyMigrations(db)
	if err != nil {
		log.Fatal("Error application migration: %v", err)
	}

	// Routers
	http.HandlerFunc("/posts", internal.HandlePosts) // Get all posts
	http.HandlerFunc("/post/", internalHandlePostDetails) // Get update delete by ID
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/assets"))))

	log.Println("Start server an http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

