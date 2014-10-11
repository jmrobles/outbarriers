package outbarriers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Context struct {
	Handler *gin.Engine
	REST    *gin.RouterGroup
	DB      gorm.DB
	/*
		DB gorm.DB
		REST *gin.RouterGroup
	*/
}

func NewContext() *Context {
	return &Context{}
}
