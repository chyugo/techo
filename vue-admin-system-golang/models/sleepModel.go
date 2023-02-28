package models

// mysql house 表
type Sleep struct {
	Id         int   `json:"id" xorm:"id" `
	UserId     int   `xorm:"user_id"`
	RecordDate int64 `xorm:"record_date"`
	SleepTime  int64 `xorm:"sleep_time"`
	GetTime    int64 `xorm:"get_time"`
	Active     int   `xorm:"active"`
}

func (Sleep) TableName() string {
	return "sleep"
}
