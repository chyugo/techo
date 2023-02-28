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

type SleepApi struct {
	sleepService service.SleepService
}

func (a SleepApi) List(c *gin.Context) {
	sleepPageNow, err := strconv.Atoi(c.PostForm("sleepPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", sleepPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	sleepList, count, err1, err2 := a.sleepService.List(user, sleepPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	sleepResponse := []resp.SleepResponse{}
	for i := 0; i < len(sleepList); i++ {
		sleepResponse = append(sleepResponse, resp.SleepResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(sleepList[i].RecordDate, 10))[0:10],
			Active:     sleepList[i].Active,
		})
	}
	resp.OK(c, gin.H{
		"sleepList":      sleepResponse,
		"sleepPageCount": count,
	})
}

func (a SleepApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	sleepList, err := a.sleepService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 处理数据
	sleepResponse := []resp.SleepResponse{}
	for i := 0; i < len(sleepList); i++ {
		sleepResponse = append(sleepResponse, resp.SleepResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(sleepList[i].RecordDate, 10))[0:10],
			Active:     sleepList[i].Active,
		})
	}
	resp.OK(c, gin.H{
		"sleepList": sleepResponse,
	})

}

func (a SleepApi) Record(c *gin.Context) {
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
	fmt.Println("时间戳", timestamp)
	// 判断有没有创建记录，没有就创建
	sleep, err := a.sleepService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("sleep", sleep)

	getTime := c.PostForm("getTime")
	sleepTime := c.PostForm("sleepTime")
	sleepHour := c.PostForm("sleepHour")
	sleepMinute := c.PostForm("sleepMinute")

	// 判断 修改
	if getTime != "" {
		hour := strings.Split(getTime, ":")[0]
		minute := strings.Split(getTime, ":")[1]
		h, _ := strconv.Atoi(hour)
		m, _ := strconv.Atoi(minute)
		getTimeStamp := utils.StampConv(i, h, m)
		fmt.Println("getTime", getTimeStamp, user)

		fmt.Println("修改 getTime记录", getTime)
		sleep.GetTime = getTimeStamp
		_, err = a.sleepService.UpdateGetTimeById(*sleep)
		if err != nil {
			resp.Error(c, err)
			return
		}
	}

	if sleepTime != "" {
		hour2 := strings.Split(sleepTime, ":")[0]
		minute2 := strings.Split(sleepTime, ":")[1]
		h2, _ := strconv.Atoi(hour2)
		m2, _ := strconv.Atoi(minute2)
		sleepTimeStamp := utils.StampConv(i, h2, m2)
		fmt.Println("sleepTimeStamp", sleepTimeStamp, user)

		fmt.Println("修改 sleepTime记录", sleepTime)
		sleep.SleepTime = sleepTimeStamp
		_, err = a.sleepService.UpdateSleepTimeById(*sleep)
		if err != nil {
			resp.Error(c, err)
			return

		}
	}

	if sleepHour != "" || sleepMinute != "" {
		sleepHour_int, _ := strconv.Atoi(sleepHour)
		sleepMinute_int, _ := strconv.Atoi(sleepMinute)
		sleepAllMinute := sleepHour_int*60 + sleepMinute_int
		sleep.Active = sleepAllMinute
		_, err = a.sleepService.UpdateSleepActiveById(*sleep)
		if err != nil {
			resp.Error(c, err)
			return

		}
	}
	resp.OK(c)
}

func (a SleepApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	sleepId := c.PostForm("sleepId")
	if sleepId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	sleepId_int, err := strconv.Atoi(sleepId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.sleepService.Delete(user, models.Sleep{Id: sleepId_int})
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

func (a SleepApi) Find(c *gin.Context) {

	sleepPageNow, err := strconv.Atoi(c.PostForm("sleepPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", sleepPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	sleepList, count, err1, err2 := a.sleepService.Find(user, sleepPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	sleepResponse := []resp.SleepTableResponse{}
	for i := 0; i < len(sleepList); i++ {
		sleepResponse = append(sleepResponse, resp.SleepTableResponse{
			Id:         sleepList[i].Id,
			RecordDate: utils.Stamp10(strconv.FormatInt(sleepList[i].RecordDate, 10))[0:10],
			SleepTime:  utils.StamptoString(sleepList[i].SleepTime),
			GetTime:    utils.StamptoString(sleepList[i].GetTime),
			Active:     utils.MinuteToHourMinute2(sleepList[i].Active),
		})
	}
	resp.OK(c, gin.H{
		"sleepList":      sleepResponse,
		"sleepPageCount": count,
	})
}

func (a SleepApi) SearchMonth(c *gin.Context) {
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
	sleepList, count, err1, err2 := a.sleepService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	sleepResponse := []resp.SleepResponse{}
	for i := 0; i < len(sleepList); i++ {
		sleepResponse = append(sleepResponse, resp.SleepResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(sleepList[i].RecordDate, 10))[0:10],
			Active:     sleepList[i].Active,
		})
	}
	fmt.Println("sleepList", sleepList)

	resp.OK(c, gin.H{
		"sleepList": sleepResponse,
		"thisMonth": utils.Stamp10_2(fmt.Sprintf("%d", thisMonth)),
		"nextMonth": utils.Stamp10_2(fmt.Sprintf("%d", nextMonth)),
	})

}
