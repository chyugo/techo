package models

// mysql house 表
type User struct {
	Id       int    `json:"id" xorm:"id" `
	UserName string `xorm:"user_name"`
	PassWord string `xorm:"pass_word"`
	DelFlag  int    `xorm:"del_flag"`
}

func (User) TableName() string {
	return "user"
}
