package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"vue-admin-system-golang/models"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
)

type HeartApi struct {
	heartService service.HeartService
}

func (a HeartApi) List(c *gin.Context) {
	heartPageNow, err := strconv.Atoi(c.PostForm("heartPageNow"))
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("pagenow", heartPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	heartList, count, err1, err2 := a.heartService.List(user, heartPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	var responseList []resp.HeartResponse
	for i := 0; i < len(heartList); i++ {
		responseList = append(responseList, resp.HeartResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(heartList[i].RecordDate, 10))[0:10],
			Normal:     float64(heartList[i].Normal) / 10,
			Abnormal:   float64(heartList[i].Abnormal) / 10,
		})
	}
	resp.OK(c, gin.H{
		"heartList":      responseList,
		"heartPageCount": count,
	})
}

func (a HeartApi) SearchMonth(c *gin.Context) {
	month := c.PostForm("month")
	thisMonth, nextMonth, err := utils.SearchMonthConv(month)
	if err != nil {
		resp.Error(c)
		return
	}
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	heartList, count, err1, err2 := a.heartService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.HeartResponse{}
	for i := 0; i < len(heartList); i++ {
		responseList = append(responseList, resp.HeartResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(heartList[i].RecordDate, 10))[0:10],
			Normal:     float64(heartList[i].Normal) / 10,
			Abnormal:   float64(heartList[i].Abnormal) / 10,
		})
	}
	resp.OK(c, gin.H{
		"heartList": responseList,
	})
}

func (a HeartApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	heartList, err := a.heartService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}

	// 处理数据
	var responseList []resp.HeartResponse
	for i := 0; i < len(heartList); i++ {
		responseList = append(responseList, resp.HeartResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(heartList[i].RecordDate, 10))[0:10],
			Normal:     float64(heartList[i].Normal) / 10,
			Abnormal:   float64(heartList[i].Abnormal) / 10,
		})
	}
	resp.OK(c, gin.H{
		"heartList": responseList,
	})
}

func (a HeartApi) Record(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	// 获取参数 发送13位时间戳 取前十位
	//获取记录时间
	t := c.PostForm("recordDate")[0:10]
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	//fmt.Println("iii", i)
	timestamp := utils.NineStamp(i)

	heartall := c.PostForm("heartall")

	heartAbnomalAll := c.PostForm("heartAbnomalAll")

	// 判断有没有创建记录，没有就创建
	heart, err := a.heartService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("heart", heart)

	// 判断 修改
	fmt.Println("修改 heart记录", heart)
	//heart.Abnormal =
	hAll, _ := strconv.ParseFloat(heartall, 10)
	hAAll, _ := strconv.ParseFloat(heartAbnomalAll, 10)
	abnormal := (hAAll / hAll) * 1000
	normal := 1000 - abnormal
	//abnormal_2f := fmt.Sprintf("%.2f", abnormal)
	//normal_2f := fmt.Sprintf("%.2f", normal)
	//fmt.Println("异常率", abnormal_2f, "正常率", normal_2f)
	heart.Normal = int(normal)
	heart.Abnormal = int(abnormal)
	_, err = a.heartService.UpdateHeartById(*heart)
	if err != nil {
		resp.Error(c, err)
	}
	resp.OK(c)

}

func (a HeartApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	heartId := c.PostForm("heartId")
	if heartId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	heartId_int, err := strconv.Atoi(heartId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.heartService.Delete(user, models.Heart{Id: heartId_int})
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

func (a HeartApi) Find(c *gin.Context) {

	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	heartPageNow, err := strconv.Atoi(c.PostForm("heartPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", heartPageNow)
	// 查询数据
	heartList, count, err1, err2 := a.heartService.Find(user, heartPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	heartResponse := []resp.HeartTableResponse{}
	for i := 0; i < len(heartList); i++ {
		heartResponse = append(heartResponse, resp.HeartTableResponse{
			Id:         heartList[i].Id,
			RecordDate: utils.Stamp10(strconv.FormatInt(heartList[i].RecordDate, 10))[0:10],
			Normal:     fmt.Sprintf("%.1f%%", float64(heartList[i].Normal)/10),
			Abnormal:   fmt.Sprintf("%.1f%%", float64(heartList[i].Abnormal)/10),
		})
	}
	resp.OK(c, gin.H{
		"heartList":      heartResponse,
		"heartPageCount": count,
	})
}
