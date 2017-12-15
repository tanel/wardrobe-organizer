package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter returns new router instance
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", GetSignup)
	router.POST("/signup", PostSignup)

	router.GET("/items/:id", GetItem)
	router.POST("/items/:id", PostItem)
	router.GET("/items", GetItems)
	router.GET("/new", GetItemsNew)
	router.POST("/new", PostItemsNew)

	router.GET("/logout", GetLogout)

	router.GET("/", GetIndex)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	return router
}
