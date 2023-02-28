package models

// mysql house 表
type Weight struct {
	Id            int   `xorm:"id" json:"id" `
	UserId        int   `xorm:"user_id"`
	RecordDate    int64 `xorm:"record_date"`
	MorningWeight int   `xorm:"morning_weight"`
	NoonWeight    int   `xorm:"noon_weight"`
	NightWeight   int   `xorm:"night_weight"`
}

func (Weight) TableName() string {
	return "weight"
}
