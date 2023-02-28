package resp

type SleepResponse struct {
	RecordDate string `xorm:"record_date"`
	SleepTime  int64  `xorm:"sleep_time"`
	GetTime    int64  `xorm:"get_time"`
	Active     int    `xorm:"active"`
}

type SleepTableResponse struct {
	Id         int    `json:"id" xorm:"id" `
	RecordDate string `xorm:"record_date"`
	SleepTime  string `xorm:"sleep_time"`
	GetTime    string `xorm:"get_time"`
	Active     string `xorm:"active"`
}
