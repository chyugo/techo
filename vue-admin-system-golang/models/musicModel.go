package models

// mysql house 表
type Music struct {
	Id         int    `json:"id" xorm:"id" `
	OldName    string `xorm:"old_name"`
	Name       string `xorm:"name"`
	UserId     int    `xorm:"user_id"`
	Title      string `xorm:"title"`
	Artist     string `xorm:"artist"`
	PicName    string `xorm:"pic_name"`
	OldPicName string `xorm:"old_pic_name"`
	LrcName    string `xorm:"lrc_name"`
	OldLrcName string `xorm:"old_lrc_name"`
	DelFlag    int    `xorm:"del_flag"`
}

func (Music) TableName() string {
	return "musics"
}
