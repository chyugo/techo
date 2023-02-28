package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitMusicRouter(r *gin.RouterGroup) {
	musicApi := new(v1.MusicApi)
	group := r.Group("/music")
	{
		//获取歌曲详情
		group.Any("/getMusic", musicApi.GetMusic)
		//上传音乐
		group.Any("/uploadMusic", musicApi.UploadMusic)
		//	确认上传音乐
		group.Any("/submitMusic", musicApi.SubmitMusic)
		//上传图片
		group.Any("/uploadPic", musicApi.UploadPic)
		// 上传歌词
		group.Any("/uploadLrc", musicApi.UploadLrc)

		// 歌曲表查询
		group.Any("/find", musicApi.Find)
		// 歌曲表删除
		group.Any("/delete", musicApi.Delete)

	}
}
