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
func (c *Context) SetupUserEP() {

	// Login
	c.REST.POST("/login", doLogin)
	c.UserREST = c.REST.Group("user")

}
