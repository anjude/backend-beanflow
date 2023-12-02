package api_service

import (
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/domain/flow/flow_service"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/constant"
)

// 匿名定义，未实现接口可以报错
var _ IFlowController = &FlowController{}

type IFlowController interface {
	AddNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	GetUserNotes(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	GetNoteList(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	GetNoteDetail(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	DelNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	LikeNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)

	AddComment(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
	GetCommentList(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
}

type FlowController struct {
	flowService flow_service.IFlowService
}

func (u FlowController) AddNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.AddNoteReq)
	resp, bizErr := u.flowService.AddNote(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) GetUserNotes(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.GetUserNotesReq)
	if req.Limit == 0 {
		req.Limit = constant.DefaultLimit
	}
	resp, bizErr := u.flowService.GetUserNotes(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) GetNoteList(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.GetNoteListReq)
	if req.Limit == 0 {
		req.Limit = constant.DefaultLimit
	}
	resp, bizErr := u.flowService.GetNoteList(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) GetNoteDetail(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.GetNoteDetailReq)
	resp, bizErr := u.flowService.GetNoteDetail(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) DelNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.DelNoteReq)
	resp, bizErr := u.flowService.DelNote(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) LikeNote(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.LikeNoteReq)
	resp, bizErr := u.flowService.LikeNote(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) AddComment(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.AddCommentReq)
	resp, bizErr := u.flowService.AddComment(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}

func (u FlowController) GetCommentList(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.GetCommentListReq)
	if req.Limit == 0 {
		req.Limit = constant.DefaultLimit
	}
	resp, bizErr := u.flowService.GetCommentList(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}
