package outbarriers

import (
	"github.com/gin-gonic/gin"
	"log"
)

/* Middleware */
func SetContext(ctx *Context) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("Context", ctx)
	}
}
func (c *Context) Init() {
	log.Printf("\tServer init")
	c.Handler = gin.Default()
	c.Handler.Use(SetContext(c))
	c.REST = c.Handler.Group("/v1")
}

func (c *Context) Start() {
	log.Printf("\tServer running...")
	c.Handler.Run(LISTENADDR)
}
