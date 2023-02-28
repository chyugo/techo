package models

// mysql house 表
type Shit struct {
	Id         int   `json:"id" xorm:"id" `
	UserId     int   `xorm:"user_id"`
	RecordDate int64 `xorm:"record_date"`
	ShitTime   int64 `xorm:"shit_time"`
}

func (Shit) TableName() string {
	return "shit"
}
