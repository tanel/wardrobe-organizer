package router

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/controller"
	"github.com/tanel/webapp/middleware"
	"github.com/tanel/webapp/session"
)

// New returns new router instance
func New(databaseConnection *sql.DB, sessionStore *session.Store) *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", middleware.HandlePublic(databaseConnection, sessionStore, controller.GetSignup))
	router.POST("/signup", middleware.HandlePublic(databaseConnection, sessionStore, controller.PostSignup))
	router.GET("/logout", middleware.HandlePublic(databaseConnection, sessionStore, controller.GetLogout))
	router.GET("/", middleware.HandlePublic(databaseConnection, sessionStore, controller.GetIndex))

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	return router
}
