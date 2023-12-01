package factory

import (
	"github.com/anjude/backend-beanflow/domain/flow/do"
	"github.com/anjude/backend-beanflow/domain/flow/entity"
)

func BuildCommentFromEntity(comment *entity.Comment) (*do.Comment, error) {
	return &do.Comment{
		ID:         comment.ID,
		Openid:     comment.Openid,
		CreateTime: comment.CreateTime,
		UpdateTime: comment.UpdateTime,
		Content:    comment.Content,
	}, nil
}
