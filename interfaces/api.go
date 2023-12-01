package interfaces

import (
	"github.com/anjude/backend-beanflow/application/api_service"
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
	user api_service.IUserController
}

func (a ApiService) RegisterRouter(engine *gin.Engine) {
	userGroup := engine.Group("api/user")
	userGroup.GET("/login", middleware.NoJWTAuth(), middleware.HandleRequest(a.user.Login, dto.LoginReq{}))
}
