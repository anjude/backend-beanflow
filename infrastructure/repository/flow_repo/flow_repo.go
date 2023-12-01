package flow_repo

import (
	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/domain/flow/entity"
	"github.com/anjude/backend-beanflow/domain/flow/factory"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
)

type IFlowRepo interface {
	CreateNote(ctx *beanctx.BizContext, note *do.Note) error
	GetUserNotes(ctx *beanctx.BizContext, openid string, offset, limit int64) ([]*do.Note, error)
	GetNoteList(ctx *beanctx.BizContext, openid string, isPublic int8, offset, limit int64) ([]*do.Note, error)
	GetNoteById(ctx *beanctx.BizContext, id int64) (*do.Note, error)
	DelNoteById(ctx *beanctx.BizContext, id int64) error
}

type FlowRepo struct {
}

func (u *FlowRepo) CreateNote(ctx *beanctx.BizContext, note *do.Note) error {
	noteEntity, err := note.ToEntity()
	if err != nil {
		return err
	}
	err = ctx.GetDb().Create(noteEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *FlowRepo) GetUserNotes(ctx *beanctx.BizContext, openid string, offset, limit int64) ([]*do.Note, error) {
	var notes []*entity.Note
	err := ctx.GetDb().Where("openid = ?", openid).Offset(int(offset)).Limit(int(limit)).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	var dos []*do.Note
	for _, note := range notes {
		noteFromEntity, err := factory.BuildNoteFromEntity(note)
		if err != nil {
			return nil, err
		}
		dos = append(dos, noteFromEntity)
	}
	return dos, nil
}

func (u *FlowRepo) GetNoteList(ctx *beanctx.BizContext, openid string, isPublic int8, offset, limit int64) ([]*do.Note, error) {
	var notes []*entity.Note
	err := ctx.GetDb().Where("openid = ? and is_public = ?", openid, isPublic).Offset(int(offset)).Limit(int(limit)).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	var dos []*do.Note
	for _, note := range notes {
		noteFromEntity, err := factory.BuildNoteFromEntity(note)
		if err != nil {
			return nil, err
		}
		dos = append(dos, noteFromEntity)
	}
	return dos, nil
}

func (u *FlowRepo) GetNoteById(ctx *beanctx.BizContext, id int64) (*do.Note, error) {
	var note *entity.Note
	err := ctx.GetDb().Where("id = ?", id).First(&note).Error
	if err != nil {
		return nil, err
	}
	noteDO, err := factory.BuildNoteFromEntity(note)
	if err != nil {
		return nil, err
	}
	return noteDO, nil
}

func (u *FlowRepo) DelNoteById(ctx *beanctx.BizContext, id int64) error {
	err := ctx.GetDb().Where("id = ?", id).Delete(&entity.Note{}).Error
	if err != nil {
		return err
	}
	return nil
}
