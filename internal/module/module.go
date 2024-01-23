package module

import (
	"gin-rush-template/internal/module/user"
	"github.com/gin-gonic/gin"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func init() {
	// Register your module here
	RegisterModule(&user.ModuleUser{})
}
