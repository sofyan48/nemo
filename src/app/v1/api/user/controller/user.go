package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/nemo/src/app/utility/rest"
)

// UserController ...
type UserController struct{}

// UserControllerHandler ...
func UserControllerHandler() *UserController {
	return &UserController{}
}

// UserControllerInterface ...
type UserControllerInterface interface{}

// UserCreate ...
func (ctrl *UserController) UserCreate(contet *gin.Context) {
	rest.ResponseMessages(context, http.StatusOK, "OK")
}
