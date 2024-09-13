package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"yehtun.com/rest-api-crud/models"
)

// CreatePost handles POST request to create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    json.NewDecoder(r.Body).Decode(&post)
    post.CreatedAt = time.Now()
    post.UpdatedAt = time.Now()

    createdPost, err := models.CreatePost(post); //id
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdPost)
}

// GetPost handles GET request to retrieve a post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["id"])
    fmt.Println("get single post id", postID)
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }


    post, err := models.GetPost(postID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(post)
}

// GetPosts handles GET request to retrieve all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
    posts, err := models.GetPosts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(posts)
}

// UpdatePost handles PUT request to update a post by ID
func UpdatePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["id"])
    fmt.Println("updaing post id", postID)
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    var post models.Post
    err = json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Set the post ID from the route, not from the request body
    post.ID = postID
    post.UpdatedAt = time.Now()

    err = models.UpdatePost(post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
  
    json.NewEncoder(w).Encode(map[string]string{"message": "Post updated successfully"})
}

// DeletePost handles DELETE request to delete a post by ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["id"])
    fmt.Println("updaing post id", postID)
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    err = models.DeletePost(postID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // w.WriteHeader(http.StatusFailedDependency)
    json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted"})
}