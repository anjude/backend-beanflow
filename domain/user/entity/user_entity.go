package entity

type User struct {
	ID         int64  `gorm:"column:id"`
	Openid     string `gorm:"column:openid"` //  小程序用户唯一标识符
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}
