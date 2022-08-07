package router

import (
	"auth0/noahjin/platform/authenticator"
	"auth0/noahjin/web/app/callback"
	"auth0/noahjin/web/app/login"
	"auth0/noahjin/web/app/logout"
	"auth0/noahjin/web/app/user"
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Register the routes and return router
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom type in cookies, we need to register
	gob.Register(map[string]interface{}{})

	// Init cookie for sessions
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
