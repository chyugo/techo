package resp

type WeightResponse struct {
	RecordDate    string  `xorm:"record_date"`
	MorningWeight float32 `xorm:"morning_weight"`
	NoonWeight    float32 `xorm:"noon_weight"`
	NightWeight   float32 `xorm:"night_weight"`
}
type WeightTableResponse struct {
	Id            int    `xorm:"id" json:"id" `
	RecordDate    string `xorm:"record_date"`
	MorningWeight string `xorm:"morning_weight"`
	NoonWeight    string `xorm:"noon_weight"`
	NightWeight   string `xorm:"night_weight"`
}
