package router

import (
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sofyan48/nemo/docs/swagger/docs"
	v1 "github.com/sofyan48/nemo/src/app/v1/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	serverHost := os.Getenv("SWAGGER_SERVER_ADDRESS")
	url := ginSwagger.URL(serverHost + "/swagger/doc.json") // The url pointing to API definition
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// End Declare Swagger

	version1 := v1.V1RouterLoaderHandler()
	version1.V1Routes(routers)
}
