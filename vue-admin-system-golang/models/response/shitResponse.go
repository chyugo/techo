package resp

type ShitResponse struct {
	RecordDate string `xorm:"record_date"`
	ShitTime   int64  `xorm:"shit_time"`
}
type ShitTableResponse struct {
	Id         int    `json:"id" xorm:"id" `
	RecordDate string `xorm:"record_date"`
	ShitTime   string `xorm:"shit_time"`
}
