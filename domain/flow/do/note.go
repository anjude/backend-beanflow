package do

import (
	"encoding/json"

	"github.com/anjude/backend-beanflow/domain/flow/entity"
)

type Note struct {
	ID         int64
	Openid     string
	CreateTime int64
	UpdateTime int64
	Content    string
	IsPublic   int8
	Extra      Extra
}

type Extra struct {
	LikeNum int64 `json:"like_num"`
}

func (n Note) ToEntity() (*entity.Note, error) {
	extraStr, err := json.Marshal(n.Extra)
	if err != nil {
		return nil, err
	}
	return &entity.Note{
		Openid:   n.Openid,
		Content:  n.Content,
		IsPublic: n.IsPublic,
		Extra:    string(extraStr),
	}, nil
}
