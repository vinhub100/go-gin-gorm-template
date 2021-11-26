package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home ...
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"message": "Hello Pong"})
}

// Hello ...
func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "greet.html", gin.H{"msg": "How Are You"})
}
