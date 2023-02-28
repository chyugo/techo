package resp

type GetMusicResponse struct {
	Title  string `json:"title" `
	Artist string `json:"artist"`
	Src    string `json:"src"`
	Pic    string `json:"pic"`
	Lrc    string `json:"lrc"`
}

type MusicTableResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title" `
	Artist     string `json:"artist"`
	OldPicName string `json:"pic"`
	OldLrcName string `json:"lrc"`
}
