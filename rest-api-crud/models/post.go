package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Body      string    `json:"body"`
    CreatedAt time.Time `json:"created"`
    UpdatedAt time.Time `json:"updated"`
}

var db *sql.DB

// SetDB sets the global DB connection for the package
func SetDB(database *sql.DB) {
    db = database
}

// CreatePost inserts a new post into the database
func CreatePost(post Post) (*Post, error) {
    result, err := db.Exec("INSERT INTO posts (title, body, created, updated) VALUES (?, ?, ?, ?)",
        post.Title, post.Body, post.CreatedAt, post.UpdatedAt)
    if err != nil {
        return nil, err
    }

    // Get the last inserted ID
    postId, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

    post.ID = int(postId)
    return &post, nil
}

// GetPost retrieves a single post by ID
func GetPost(id int) (*Post, error) {
    var post Post
    var createdAtRaw, updatedAtRaw []uint8 
    // QueryRow to retrieve the post with a specific ID
    err := db.QueryRow("SELECT id, title, body, created, updated FROM posts WHERE id = ?", id).Scan(
        &post.ID, &post.Title, &post.Body, &createdAtRaw, &updatedAtRaw)
        if(err != nil){
            log.Println("Error scanning post:", err)
            return nil, err
        }

        post.CreatedAt, err = parseTime(createdAtRaw)
        if err != nil {
            log.Println("Error parsing createdAt:", err)
            return nil, err
        }

        post.UpdatedAt, err = parseTime(updatedAtRaw)
        if err != nil {
            log.Println("Error parsing updatedAt:", err)
            return nil, err
        }
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle case where no rows were found
            log.Printf("No post found with ID %d", id)
            return &post, errors.New("post not found")
        }
        // Handle other errors
        log.Printf("Error retrieving post with ID %d: %v", id, err)
        return &post, err
    }

    return &post, nil
}

// GetPosts retrieves all posts from the database
func GetPosts() ([]Post, error) {
    rows, err := db.Query("SELECT id, title, body, created, updated FROM posts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        var createdAtRaw, updatedAtRaw []uint8 // For handling the []uint8 type from the DB
        
        err := rows.Scan(&post.ID, &post.Title, &post.Body, &createdAtRaw, &updatedAtRaw)
        if err != nil {
            return nil, err
        }
        
        // Parse createdAtRaw and updatedAtRaw to time.Time
        post.CreatedAt, err = parseTime(createdAtRaw)
        if err != nil {
            log.Println("Error parsing createdAt:", err)
            return nil, err
        }

        post.UpdatedAt, err = parseTime(updatedAtRaw)
        if err != nil {
            log.Println("Error parsing updatedAt:", err)
            return nil, err
        }

        posts = append(posts, post)
    }

    return posts, nil
}

// UpdatePost updates a post by ID
func UpdatePost(post Post) error {
    _, err := db.Exec("UPDATE posts SET title = ?, body = ?, updated = ? WHERE id = ?",
        post.Title, post.Body, post.UpdatedAt, post.ID)
    return err
}

// DeletePost deletes a post by ID
func DeletePost(id int) error {
    _, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
    return err
}
// Helper function to convert []uint8 to time.Time
func parseTime(rawTime []uint8) (time.Time, error) {
    timeStr := string(rawTime)
    // Use the appropriate layout based on your database format
    layout := "2006-01-02 15:04:05"  // MySQL DATETIME format
    return time.Parse(layout, timeStr)
}