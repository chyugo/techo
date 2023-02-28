package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
)

type WakeApi struct {
	wakeService service.WakeService
}

func (a WakeApi) SearchMonth(c *gin.Context) {
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
	wakeList, count, err1, err2 := a.wakeService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	wakeResponse := []resp.WakeResponse{}
	for i := 0; i < len(wakeList); i++ {
		sleeptime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].SleepTime, "000"), 10, 64)
		gettime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].GetTime, "000"), 10, 64)
		wakeResponse = append(wakeResponse, resp.WakeResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(wakeList[i].RecordDate, 10))[0:10],
			SleepTime:  sleeptime,
			GetTime:    gettime,
		})
	}
	resp.OK(c, gin.H{
		"wakeList": wakeResponse,
	})

}

func (a WakeApi) List(c *gin.Context) {
	wakePageNow, err := strconv.Atoi(c.PostForm("wakePageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", wakePageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	wakeList, count, err1, err2 := a.wakeService.List(user, wakePageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	wakeResponse := []resp.WakeResponse{}
	for i := 0; i < len(wakeList); i++ {
		sleeptime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].SleepTime, "000"), 10, 64)
		gettime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].GetTime, "000"), 10, 64)
		wakeResponse = append(wakeResponse, resp.WakeResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(wakeList[i].RecordDate, 10))[0:10],
			SleepTime:  sleeptime,
			GetTime:    gettime,
		})
	}
	resp.OK(c, gin.H{
		"wakeList":      wakeResponse,
		"wakePageCount": count,
	})

}

func (a WakeApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	wakeList, err := a.wakeService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 处理数据
	wakeResponse := []resp.WakeResponse{}
	for i := 0; i < len(wakeList); i++ {
		sleeptime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].SleepTime, "000"), 10, 64)
		gettime, _ := strconv.ParseInt(fmt.Sprintf("%s%s", wakeList[i].GetTime, "000"), 10, 64)
		wakeResponse = append(wakeResponse, resp.WakeResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(wakeList[i].RecordDate, 10))[0:10],
			SleepTime:  sleeptime,
			GetTime:    gettime,
		})
	}

	resp.OK(c, gin.H{
		"wakeList": wakeResponse,
	})

}
