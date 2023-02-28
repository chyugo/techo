package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"strings"
	"vue-admin-system-golang/config"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/models/request"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
)

type MusicApi struct {
	musicService service.MusicService
}

func (a MusicApi) UploadMusic(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println("file", file)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	// 获取文件后缀
	temp := strings.Split(file.Filename, ".")
	format := temp[len(temp)-1]
	filename := utils.CreateFileSecret()
	filepath := fmt.Sprintf("%s.%s", filename, format)
	dst := path.Join("./static/music", filepath)
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		resp.Error(c, saveErr.Error())
		return
	}
	resp.OK(c, gin.H{
		"oldFileName": file.Filename,
		"baseName":    temp[0],
		"filePath":    filepath,
	})
}

func (a MusicApi) UploadPic(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println("file", file)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	// 获取文件后缀
	temp := strings.Split(file.Filename, ".")
	format := temp[len(temp)-1]
	filename := utils.CreateFileSecret()
	filepath := fmt.Sprintf("%s.%s", filename, format)
	dst := path.Join("./static/music", filepath)
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		resp.Error(c, saveErr.Error())
		return
	}
	resp.OK(c, gin.H{
		"oldPicName": file.Filename,
		"picPath":    filepath,
	})
}

func (a MusicApi) UploadLrc(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println("file", file)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	// 获取文件后缀
	temp := strings.Split(file.Filename, ".")
	format := temp[len(temp)-1]
	filename := utils.CreateFileSecret()
	filepath := fmt.Sprintf("%s.%s", filename, format)
	dst := path.Join("./static/music", filepath)
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		resp.Error(c, saveErr.Error())
		return
	}
	resp.OK(c, gin.H{
		"oldLrcName": file.Filename,
		"lrcPath":    filepath,
	})
}

func (a MusicApi) SubmitMusic(c *gin.Context) {
	query := request.SubmitMusicRequest{}
	if err := c.Bind(&query); err != nil {
		resp.Error(c)
		fmt.Println("err", err)
		return
	}
	fmt.Println("query", query)
	if query.FilePath == "" || query.OldFileName == "" || query.Title == "" || query.Artist == "" {
		resp.Error(c, "缺少参数")
		return
	}
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("user", user)
	err = a.musicService.CreateMusic(models.Music{
		UserId:     user.Id,
		OldName:    query.OldFileName,
		Name:       query.FilePath,
		Title:      query.Title,
		Artist:     query.Artist,
		PicName:    query.PicPath,
		OldPicName: query.OldPicName,
		LrcName:    query.LrcPath,
		OldLrcName: query.OldLrcPath,
	})
	if err != nil {
		fmt.Println(err)
		resp.Error(c)
		return
	}
	resp.OK(c)

}

func (a MusicApi) GetMusic(c *gin.Context) {
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("user", user)
	musicList, err := a.musicService.GetMusic(models.Music{
		UserId: user.Id,
	})
	if err != nil {
		fmt.Println(err)
		resp.Error(c)
		return
	}
	var response []resp.GetMusicResponse
	for i := 0; i < len(musicList); i++ {
		response = append(response, resp.GetMusicResponse{
			Title:  musicList[i].Title,
			Artist: musicList[i].Artist,
			Src:    fmt.Sprintf("http://%s%s/more_static/%s", config.MySqlCfg.Host, config.AppCfg.Server, musicList[i].Name),
			Pic:    fmt.Sprintf("http://%s%s/more_static/%s", config.MySqlCfg.Host, config.AppCfg.Server, musicList[i].PicName),
			Lrc:    fmt.Sprintf("http://%s%s/more_static/%s", config.MySqlCfg.Host, config.AppCfg.Server, musicList[i].LrcName),
		})
	}
	resp.OK(c, response)

}
func (a MusicApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	musicPageNow, err := strconv.Atoi(c.PostForm("musicPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("musicPageNow", musicPageNow)
	// 查询数据
	musicList, count, err1, err2 := a.musicService.Find(user, musicPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	musicResponse := []resp.MusicTableResponse{}
	for i := 0; i < len(musicList); i++ {
		musicResponse = append(musicResponse, resp.MusicTableResponse{
			Id:         musicList[i].Id,
			Title:      musicList[i].Title,
			Artist:     musicList[i].Artist,
			OldPicName: musicList[i].OldPicName,
			OldLrcName: musicList[i].OldLrcName,
		})
	}
	resp.OK(c, gin.H{
		"musicList":      musicResponse,
		"musicPageCount": count,
	})
}

func (a MusicApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	musicId := c.PostForm("musicId")
	if musicId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	musicId_int, err := strconv.Atoi(musicId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.musicService.Delete(user, models.Music{Id: musicId_int})
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("i", i)
	if i != 0 {
		resp.OK(c)
		return
	}
	resp.Error(c)
}
