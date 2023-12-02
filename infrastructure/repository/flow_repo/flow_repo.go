package flow_repo

import (
	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/domain/flow/entity"
	"github.com/anjude/backend-beanflow/domain/flow/factory"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IFlowRepo interface {
	CreateNote(ctx *beanctx.BizContext, note *do.Note) error
	GetUserNotes(ctx *beanctx.BizContext, openid string, offset, limit int64) ([]*do.Note, error)
	GetNoteList(ctx *beanctx.BizContext, openid string, isPublic int8, offset, limit int64) ([]*do.Note, error)
	GetNoteById(ctx *beanctx.BizContext, id int64) (*do.Note, error)
	DelNoteById(ctx *beanctx.BizContext, id int64) error
	LikeNoteExist(ctx *beanctx.BizContext, openid string, id int64) (bool, error)
	UpdateNote(ctx *beanctx.BizContext, note *do.Note) error

	AddLikeNoteRecord(ctx *beanctx.BizContext, openid string, noteId int64) error
	DelLikeNoteRecord(ctx *beanctx.BizContext, openid string, noteId int64) error

	CreateComment(ctx *beanctx.BizContext, comment *do.Comment) error
	GetCommentList(ctx *beanctx.BizContext, noteId, offset, limit int64) ([]*do.Comment, error)
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
	err := ctx.GetDb().Where("openid = ?", openid).Offset(int(offset)).Limit(int(limit)).Order("update_time desc").Find(&notes).Error
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
	err := ctx.GetDb().Where("openid = ? and is_public = ?", openid, isPublic).Offset(int(offset)).Limit(int(limit)).Order("update_time desc").Find(&notes).Error
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

func (u *FlowRepo) LikeNoteExist(ctx *beanctx.BizContext, openid string, id int64) (bool, error) {
	err := ctx.GetDb().Where("openid = ? and note_id = ?", openid, id).Take(&entity.NoteLikeRecord{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		ctx.Log().Errorf("get note like record error: %v", err)
		return false, err
	}
	return true, nil
}

func (u *FlowRepo) UpdateNote(ctx *beanctx.BizContext, note *do.Note) error {
	noteEntity, err := note.ToEntity()
	if err != nil {
		return err
	}
	err = ctx.GetDb().Model(&entity.Note{}).Where("id = ?", note.ID).Updates(noteEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *FlowRepo) AddLikeNoteRecord(ctx *beanctx.BizContext, openid string, noteId int64) error {
	err := ctx.GetDb().Create(&entity.NoteLikeRecord{
		Openid: openid,
		NoteId: noteId,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *FlowRepo) DelLikeNoteRecord(ctx *beanctx.BizContext, openid string, noteId int64) error {
	err := ctx.GetDb().Where("openid = ? and note_id = ?", openid, noteId).Delete(&entity.NoteLikeRecord{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *FlowRepo) CreateComment(ctx *beanctx.BizContext, comment *do.Comment) error {
	commentEntity, err := comment.ToEntity()
	if err != nil {
		return err
	}
	err = ctx.GetDb().Create(commentEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *FlowRepo) GetCommentList(ctx *beanctx.BizContext, noteId, offset, limit int64) ([]*do.Comment, error) {
	var comments []*entity.Comment
	err := ctx.GetDb().Where("note_id = ?", noteId).Offset(int(offset)).Limit(int(limit)).Order("update_time desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	var dos []*do.Comment
	for _, comment := range comments {
		commentDO, err := factory.BuildCommentFromEntity(comment)
		if err != nil {
			return nil, err
		}
		dos = append(dos, commentDO)
	}
	return dos, nil
}
