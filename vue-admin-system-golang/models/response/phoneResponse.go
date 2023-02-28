package resp

// mysql house 表
type PhoneResponse struct {
	RecordDate  string `xorm:"record_date"`
	PhoneMinute int    `xorm:"phone_minute"`
}
type PhoneTableResponse struct {
	Id          int    `json:"id" xorm:"id" `
	RecordDate  string `xorm:"record_date"`
	PhoneMinute string `xorm:"phone_minute"`
}
