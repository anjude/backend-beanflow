package do

import "github.com/anjude/backend-beanflow/domain/flow/entity"

type Comment struct {
	ID         int64
	Openid     string
	CreateTime int64
	UpdateTime int64
	NoteId     int64
	Content    string
}

func (c Comment) ToEntity() (*entity.Comment, error) {
	return &entity.Comment{
		ID:      c.ID,
		Openid:  c.Openid,
		NoteId:  c.NoteId,
		Content: c.Content,
	}, nil
}
