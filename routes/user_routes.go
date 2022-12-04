package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang-gorm-postgresql/controllers"
	"github.com/onainadapdap1/golang-gorm-postgresql/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}

