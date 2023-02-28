package models

// mysql house 表
type Water struct {
	Id         int   `xorm:"id" json:"id" `
	UserId     int   `xorm:"user_id"`
	RecordDate int64 `xorm:"record_date"`
	WaterValue int   `xorm:"water_value"`
}

func (Water) TableName() string {
	return "water"
}
