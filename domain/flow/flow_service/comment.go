package flow_service

import (
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
)

func (u FlowService) AddComment(ctx *beanctx.BizContext, req dto.AddCommentReq) (interface{}, *beanerr.BizError) {
	err := u.flowRepo.CreateComment(ctx, &do.Comment{
		Openid:  ctx.GetOpenid(),
		Content: req.Content,
		NoteId:  req.NoteId,
	})
	if err != nil {
		ctx.Log().Errorf("create note error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	return nil, nil
}

func (u FlowService) GetCommentList(ctx *beanctx.BizContext, req dto.GetCommentListReq) (interface{}, *beanerr.BizError) {
	comments, err := u.flowRepo.GetCommentList(ctx, req.NoteId, req.Offset, req.Limit)
	if err != nil {
		ctx.Log().Errorf("get comment list error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	commentList := make([]*dto.CommentView, len(comments))
	for i, comment := range comments {
		commentList[i] = dto.BuildCommentView(comment)
	}
	return &dto.GetCommentListResp{
		Offset: req.Offset,
		Limit:  req.Limit,
		List:   commentList,
	}, nil
}
