package flow_service

import (
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/enum"
	"github.com/anjude/backend-beanflow/infrastructure/repository/flow_repo"
)

type IFlowService interface {
	AddNote(ctx *beanctx.BizContext, req dto.AddNoteReq) (interface{}, *beanerr.BizError)
	GetUserNotes(ctx *beanctx.BizContext, req dto.GetUserNotesReq) (interface{}, *beanerr.BizError)
	GetNoteList(ctx *beanctx.BizContext, req dto.GetNoteListReq) (interface{}, *beanerr.BizError)
	GetNoteDetail(ctx *beanctx.BizContext, req dto.GetNoteDetailReq) (interface{}, *beanerr.BizError)
	DelNote(ctx *beanctx.BizContext, req dto.DelNoteReq) (interface{}, *beanerr.BizError)
}

type FlowService struct {
	flowRepo flow_repo.IFlowRepo
}

func (u FlowService) AddNote(ctx *beanctx.BizContext, req dto.AddNoteReq) (interface{}, *beanerr.BizError) {
	err := u.flowRepo.CreateNote(ctx, &do.Note{
		Openid:   ctx.GetOpenid(),
		Content:  req.Content,
		IsPublic: *req.IsPublic,
	})
	if err != nil {
		ctx.Log().Errorf("create note error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	return nil, nil
}

func (u FlowService) GetUserNotes(ctx *beanctx.BizContext, req dto.GetUserNotesReq) (interface{}, *beanerr.BizError) {
	notes, err := u.flowRepo.GetUserNotes(ctx, ctx.GetOpenid(), req.Offset, req.Limit)
	if err != nil {
		ctx.Log().Errorf("get user notes error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	var noteList []*dto.NoteView
	for _, note := range notes {
		noteList = append(noteList, dto.BuildNoteView(note))
	}
	return &dto.GetUserNotesResp{
		Offset: req.Offset,
		Limit:  req.Limit,
		List:   noteList,
	}, nil
}

func (u FlowService) GetNoteList(ctx *beanctx.BizContext, req dto.GetNoteListReq) (interface{}, *beanerr.BizError) {
	// 只获取公开的笔记
	notes, err := u.flowRepo.GetNoteList(ctx, req.Openid, int8(enum.PublicNote), req.Offset, req.Limit)
	if err != nil {
		ctx.Log().Errorf("get note list error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	var noteList []*dto.NoteView
	for _, note := range notes {
		noteList = append(noteList, dto.BuildNoteView(note))
	}
	return &dto.GetUserNotesResp{
		Offset: req.Offset,
		Limit:  req.Limit,
		List:   noteList,
	}, nil
}

func (u FlowService) GetNoteDetail(ctx *beanctx.BizContext, req dto.GetNoteDetailReq) (interface{}, *beanerr.BizError) {
	note, err := u.flowRepo.GetNoteById(ctx, req.NoteId)
	if err != nil {
		ctx.Log().Errorf("get note detail error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	return dto.BuildNoteView(note), nil
}

func (u FlowService) DelNote(ctx *beanctx.BizContext, req dto.DelNoteReq) (interface{}, *beanerr.BizError) {
	note, err := u.flowRepo.GetNoteById(ctx, req.NoteId)
	if err != nil {
		ctx.Log().Errorf("get note detail error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	if note.Openid != ctx.GetOpenid() {
		return nil, beanerr.NoPermission.SetDetail("can only delete your own note")
	}
	err = u.flowRepo.DelNoteById(ctx, req.NoteId)
	if err != nil {
		ctx.Log().Errorf("del note error: %v", err)
		return nil, beanerr.DBError.SetDetail(err.Error())
	}
	return nil, nil
}
