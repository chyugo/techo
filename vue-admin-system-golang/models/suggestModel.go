package models

type Suggest struct {
	Id      int    `json:"id" xorm:"id" `
	Content string `xorm:"content"`
}

func (Suggest) TableName() string {
	return "suggest"
}
