package request

type SubmitMusicRequest struct {
	OldFileName string `form:"oldFileName"`
	FilePath    string `form:"filePath"`
	Title       string `form:"title"`
	Artist      string `form:"artist"`
	PicPath     string `form:"picPath"`
	OldPicName  string `form:"oldPicName"`
	LrcPath     string `form:"lrcPath"`
	OldLrcPath  string `form:"oldLrcName"`
}
