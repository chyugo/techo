package models

// mysql house 表
type Phone struct {
	Id          int   `json:"id" xorm:"id" `
	UserId      int   `xorm:"user_id"`
	RecordDate  int64 `xorm:"record_date"`
	PhoneMinute int   `xorm:"phone_minute"`
}

func (Phone) TableName() string {
	return "phone"
}
