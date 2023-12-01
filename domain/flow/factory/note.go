package factory

import (
	"encoding/json"

	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/domain/flow/entity"
)

func BuildNoteFromEntity(note *entity.Note) (*do.Note, error) {
	extra := do.Extra{}
	err := json.Unmarshal([]byte(note.Extra), &extra)
	if err != nil {
		return nil, err
	}
	return &do.Note{
		ID:         note.ID,
		Openid:     note.Openid,
		CreateTime: note.CreateTime,
		UpdateTime: note.UpdateTime,
		Content:    note.Content,
		IsPublic:   note.IsPublic,
		Extra:      extra,
	}, nil
}
