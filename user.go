package outbarriers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type LoginJSON struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func doLogin(c *gin.Context) {

	//log.Printf("\tLOGIN")
	var json LoginJSON
	var user *User
	ctx := c.MustGet("Context").(*Context)
	c.Bind(&json)
	user = ctx.AuthUser(json.Email, json.Password)
	if user.Id > 0 {
		log.Printf("User logged")
		token := ctx.LoginUser(user)
		c.JSON(200, gin.H{"status": true, "auth-token": token})
	} else {
		c.JSON(401, gin.H{"status": false})
	}
}
func doAuth(c *gin.Context) {

	//c.JSON(200, gin.H{"status": true})
	method := c.Request.Header.Get("X-Original-Request-Method")
	//log.Printf("Auth method: %s", method)
	if method == "GET" {
		c.Writer.WriteHeader(200)
		return
	}
	token := c.Request.Header.Get("X-Auth-Token")
	if token == "" {
		c.Writer.WriteHeader(403)
		return
	}
	ctx := c.MustGet("Context").(*Context)
	// Get session
	session := ctx.GetSessionByToken(token)
	if session == nil {
		c.Writer.WriteHeader(403)
		return
	}
	// TODO: check for expired time
	c.Writer.WriteHeader(200)
}

func (c *Context) SetupUserEP() {

	// Login
	c.REST.POST("/login", doLogin)
	// Auth
	c.REST.GET("/auth", doAuth)
	c.UserREST = c.REST.Group("user")

}
