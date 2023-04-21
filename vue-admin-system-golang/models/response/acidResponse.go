package resp

type AcidResponse struct {
	RecordDate string `xorm:"record_date"`
	Acid       int    `xorm:"water"`
}
type AcidTableResponse struct {
	Id         int    `xorm:"id" json:"id" `
	RecordDate string `xorm:"record_date"`
	Acid       string `xorm:"water"`
	Judge      int
}
type AcidDateResponse struct {
	WorkingDay string        `json:"workingDay"`
	Content    []AcidContent `json:"content"`
}

type AcidContent struct {
	Notice string `json:"notice"`
	Type   string `json:"type"`
}
