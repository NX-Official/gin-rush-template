package user

import (
	"github.com/gin-gonic/gin"
)

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {

	r.POST("/login", Login)
	r.POST("/register", Create)

}
