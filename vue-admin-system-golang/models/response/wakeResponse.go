package resp

type WakeResponse struct {
	RecordDate string `xorm:"record_date"`
	GetTime    int64  `xorm:"get_time"`
	SleepTime  int64  `xorm:"sleep_time"`
}
