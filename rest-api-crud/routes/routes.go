package routes

import (
	"github.com/gorilla/mux"
	"yehtun.com/rest-api-crud/controllers"
)

// SetupRoutes defines all routes for the API
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // CRUD routes for posts
    router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
    router.HandleFunc("/posts/{id:[0-9]+}", controllers.GetPost).Methods("GET")
    router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
    router.HandleFunc("/posts/{id:[0-9]+}", controllers.UpdatePost).Methods("PUT")
    router.HandleFunc("/posts/{id:[0-9]+}", controllers.DeletePost).Methods("DELETE")

    return router
}