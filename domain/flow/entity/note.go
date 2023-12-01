package entity

type Note struct {
	ID         int64  `gorm:"column:id"`
	Openid     string `gorm:"column:openid"` //  小程序用户唯一标识符
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"`
	Content    string `gorm:"column:content"`   //  笔记内容
	Extra      string `gorm:"column:extra"`     //  笔记额外数据放extra，存成json
	IsPublic   int8   `gorm:"column:is_public"` //  是否公开
}

func (Note) TableName() string {
	return "note"
}
