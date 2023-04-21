package resp

type ShitResponse struct {
	RecordDate string `xorm:"record_date"`
	ShitTime   int64  `xorm:"shit_time"`
}
type ShitTableResponse struct {
	Id         int    `json:"id" xorm:"id" `
	RecordDate string `xorm:"record_date"`
	ShitTime   string `xorm:"shit_time"`
	Judge      int    `xorm:"judge"`
}
type ShitDateResponse struct {
	WorkingDay string    `json:"workingDay"`
	Content    []Content `json:"content"`
}

type Content struct {
	Notice string `json:"notice"`
	Type   string `json:"type"`
}
