package interfaces

import (
	"github.com/anjude/backend-beanflow/application/api_service"
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
	user api_service.IUserController
	flow api_service.IFlowController
}

func (a ApiService) RegisterRouter(engine *gin.Engine) {
	userGroup := engine.Group("api/user")
	userGroup.GET("/login", middleware.NoJWTAuth(), middleware.HandleRequest(a.user.Login, dto.LoginReq{}))

	flowGroup := engine.Group("api/flow")
	flowGroup.POST("/note/add", middleware.HandleRequest(a.flow.AddNote, dto.AddNoteReq{}))
	// 获取用户笔记列表
	flowGroup.GET("/note/user_notes", middleware.HandleRequest(a.flow.GetUserNotes, dto.GetUserNotesReq{}))
	flowGroup.GET("/note/list", middleware.HandleRequest(a.flow.GetNoteList, dto.GetNoteListReq{}))
	flowGroup.GET("/note/detail", middleware.HandleRequest(a.flow.GetNoteDetail, dto.GetNoteDetailReq{}))
}
