package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/controller"
)

// New returns new router instance
func New() *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", controller.GetSignup)
	router.POST("/signup", controller.PostSignup)

	router.GET("/items/:id", controller.GetItem)
	router.POST("/items/:id", controller.PostItem)
	router.GET("/items", controller.GetItems)
	router.GET("/new", controller.GetItemsNew)
	router.POST("/new", controller.PostItemsNew)

	router.GET("/logout", controller.GetLogout)

	router.GET("/", controller.GetIndex)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	return router
}