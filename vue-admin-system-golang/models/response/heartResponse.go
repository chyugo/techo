package resp

type HeartResponse struct {
	RecordDate string  `xorm:"record_date" json:"date"`
	Normal     float64 `xorm:"normal" json:"normal"`
	Abnormal   float64 `xorm:"abnormal" json:"abnormal"`
}
type HeartTableResponse struct {
	Id         int    `json:"id" xorm:"id" `
	UserId     int    `xorm:"user_id"`
	RecordDate string `xorm:"record_date" json:"date"`
	Normal     string `xorm:"normal" json:"normal"`
	Abnormal   string `xorm:"abnormal" json:"abnormal"`
}
