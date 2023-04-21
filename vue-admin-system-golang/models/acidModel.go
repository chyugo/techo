package models

// mysql house 表
type Acid struct {
	Id         int   `xorm:"id" json:"id" `
	UserId     int   `xorm:"user_id"`
	RecordDate int64 `xorm:"record_date"`
	AcidValue  int   `xorm:"acid_value"`
	Judge      int   `xorm:"judge"`
}

func (Acid) TableName() string {
	return "acid"
}
