package models

// mysql house 表
type Wake struct {
	Id         int    `json:"id" xorm:"id" `
	UserId     int    `xorm:"user_id"`
	RecordDate int64  `xorm:"record_date"`
	GetTime    string `xorm:"get_time"`
	SleepTime  string `xorm:"sleep_time"`
}

func (Wake) TableName() string {
	return "sleep"
}
