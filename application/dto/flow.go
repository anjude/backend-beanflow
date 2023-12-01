package dto

import "github.com/anjude/backend-beanflow/domain/flow/do"

// AddNoteReq 用指针可以判断是否传入，go里int8默认值为0，不使用指针无法判断是否传入
type AddNoteReq struct {
	Content  string `json:"content" binding:"required"`
	IsPublic *int8  `json:"is_public" binding:"required"`
}

// GetUserNotesReq 用form可以获取url中的参数，用json可以获取body中的参数
type GetUserNotesReq struct {
	Offset int64 `form:"offset" binding:"-"`
	Limit  int64 `form:"limit" binding:"-"`
}

type GetUserNotesResp struct {
	List   []*NoteView `json:"list"`
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
}

type NoteView struct {
	ID         int64  `json:"id"`
	Openid     string `json:"openid"` //  小程序用户唯一标识符
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Content    string `json:"content"` //  笔记内容
	LikeNum    int64  `json:"like_num"`
	IsPublic   int8   `json:"is_public"` //  是否公开
}

func BuildNoteView(note *do.Note) *NoteView {
	return &NoteView{
		ID:         note.ID,
		Openid:     note.Openid,
		CreateTime: note.CreateTime,
		UpdateTime: note.UpdateTime,
		Content:    note.Content,
		LikeNum:    note.Extra.LikeNum,
		IsPublic:   note.IsPublic,
	}
}

type GetNoteListReq struct {
	Openid string `form:"openid" binding:"required"`
	Offset int64  `form:"offset" binding:"-"`
	Limit  int64  `form:"limit" binding:"-"`
}

type GetNoteListResp struct {
	List   []*NoteView `json:"list"`
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
}

type GetNoteDetailReq struct {
	NoteId int64 `form:"note_id" binding:"required"`
}

type GetNoteDetailResp NoteView

type DelNoteReq struct {
	NoteId int64 `json:"note_id" binding:"required"`
}
