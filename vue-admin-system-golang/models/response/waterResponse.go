package resp

type WaterResponse struct {
	RecordDate string `xorm:"record_date"`
	Water      int    `xorm:"water"`
}
type WaterTableResponse struct {
	Id         int    `xorm:"id" json:"id" `
	RecordDate string `xorm:"record_date"`
	Water      string `xorm:"water"`
}
