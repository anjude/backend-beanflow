package dto

import "github.com/anjude/backend-beanflow/domain/flow/do"

type AddCommentReq struct {
	Content string `json:"content" binding:"required"`
	NoteId  int64  `json:"note_id" binding:"required"`
}

type GetCommentListReq struct {
	NoteId int64 `form:"note_id" binding:"required"`
	Offset int64 `form:"offset" binding:"-"`
	Limit  int64 `form:"limit" binding:"-"`
}

type GetCommentListResp struct {
	List   []*CommentView `json:"list"`
	Offset int64          `json:"offset"`
	Limit  int64          `json:"limit"`
}

type CommentView struct {
	ID         int64  `json:"id"`
	Openid     string `json:"openid"` //  小程序用户唯一标识符
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	NoteId     int64  `json:"note_id"` //  笔记id
	Content    string `json:"content"` //  留言内容
}

func BuildCommentView(note *do.Comment) *CommentView {
	return &CommentView{
		ID:         note.ID,
		Openid:     note.Openid,
		CreateTime: note.CreateTime,
		UpdateTime: note.UpdateTime,
		Content:    note.Content,
		NoteId:     note.NoteId,
	}
}
