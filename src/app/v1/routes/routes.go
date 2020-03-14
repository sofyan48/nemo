package routes

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/nemo/src/app/v1/api/health/controller"
	user "github.com/sofyan48/nemo/src/app/v1/api/user/controller"
	"github.com/sofyan48/nemo/src/middleware"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	User       user.UserControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health: health.HealthControllerHandler(),
		User:   user.UserControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initUser(router)
}
