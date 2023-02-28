package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"vue-admin-system-golang/models"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
)

type ShitApi struct {
	shitService service.ShitService
}

func (a ShitApi) SearchMonth(c *gin.Context) {
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
	shitList, count, err1, err2 := a.shitService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.ShitResponse{}
	for i := 0; i < len(shitList); i++ {
		responseList = append(responseList, resp.ShitResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(shitList[i].RecordDate, 10))[0:10],
			ShitTime:   shitList[i].ShitTime * 1000,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", responseList)
	resp.OK(c, gin.H{
		"shitList": responseList,
	})
}

func (a ShitApi) List(c *gin.Context) {
	shitPageNow, err := strconv.Atoi(c.PostForm("shitPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", shitPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	shitList, count, err1, err2 := a.shitService.List(user, shitPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	var responseList []resp.ShitResponse
	for i := 0; i < len(shitList); i++ {
		responseList = append(responseList, resp.ShitResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(shitList[i].RecordDate, 10))[0:10],
			ShitTime:   shitList[i].ShitTime * 1000,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", responseList)
	resp.OK(c, gin.H{
		"shitList":      responseList,
		"shitPageCount": count,
	})
}

func (a ShitApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	shitList, err := a.shitService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 处理数据
	var responseList []resp.ShitResponse
	for i := 0; i < len(shitList); i++ {
		responseList = append(responseList, resp.ShitResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(shitList[i].RecordDate, 10))[0:10],
			ShitTime:   shitList[i].ShitTime * 1000,
		})
	}
	resp.OK(c, gin.H{
		"shitList": responseList,
	})

}

func (a ShitApi) Record(c *gin.Context) {
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

	shitTime := c.PostForm("shitTime")
	hour := strings.Split(shitTime, ":")[0]
	minute := strings.Split(shitTime, ":")[1]
	h, _ := strconv.Atoi(hour)
	m, _ := strconv.Atoi(minute)
	shitTimeStamp := utils.StampConv(i, h, m)
	fmt.Println("timestamp", shitTimeStamp, user)

	// 判断有没有创建记录，没有就创建
	shit, err := a.shitService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("weight", shit)

	// 判断 修改
	fmt.Println("修改 shit记录", shit)
	shit.ShitTime = shitTimeStamp
	_, err = a.shitService.UpdateShitById(*shit)
	if err != nil {
		resp.Error(c, err)
	}
	resp.OK(c)

}

func (a ShitApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	shitId := c.PostForm("shitId")
	if shitId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	shitId_int, err := strconv.Atoi(shitId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.shitService.Delete(user, models.Shit{Id: shitId_int})
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

func (a ShitApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	shitPageNow, err := strconv.Atoi(c.PostForm("shitPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("shitPageNow", shitPageNow)
	// 查询数据
	shitList, count, err1, err2 := a.shitService.Find(user, shitPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	shitResponse := []resp.ShitTableResponse{}
	for i := 0; i < len(shitList); i++ {
		shitResponse = append(shitResponse, resp.ShitTableResponse{
			Id:         shitList[i].Id,
			RecordDate: utils.Stamp10(strconv.FormatInt(shitList[i].RecordDate, 10))[0:10],
			ShitTime:   utils.StamptoString(shitList[i].ShitTime),
		})
	}
	resp.OK(c, gin.H{
		"shitList":      shitResponse,
		"shitPageCount": count,
	})
}