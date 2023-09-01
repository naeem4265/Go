package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/api-server/handlers"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Post("/signin", handlers.SignIn)
	router.Get("/signout", handlers.SignOut)

	router.Get("/albums", handlers.GetAlbums)
	router.Get("/albums/{id}", handlers.GetAlbumById)
	router.Put("/albums/{id}", handlers.PutAlbum)
	router.Post("/albums", handlers.PostAlbum)
	router.Delete("/albums/{id}", handlers.DeleteAlbum)

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// move the album related funcs
// return http status code
// create a middleware for auth
// implement signout
//
