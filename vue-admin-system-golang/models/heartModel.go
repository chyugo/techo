package models

// mysql house 表
type Heart struct {
	Id         int   `json:"id" xorm:"id" `
	UserId     int   `xorm:"user_id"`
	RecordDate int64 `xorm:"record_date" json:"date"`
	Normal     int   `xorm:"normal" json:"normal"`
	Abnormal   int   `xorm:"abnormal" json:"abnormal"`
}

func (Heart) TableName() string {
	return "heart"
}
