package entity

type Comment struct {
	ID         int64  `gorm:"column:id"`
	Openid     string `gorm:"column:openid"` //  小程序用户唯一标识符
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"`
	NoteId     int64  `gorm:"column:note_id"` //  笔记id
	Content    string `gorm:"column:content"` //  笔记内容
}

func (Comment) TableName() string {
	return "comment"
}
