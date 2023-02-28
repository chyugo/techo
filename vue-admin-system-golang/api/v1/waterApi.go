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

type WaterApi struct {
	waterService service.WaterService
}

func (a WaterApi) SearchMonth(c *gin.Context) {
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
	waterList, count, err1, err2 := a.waterService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("查询数据", waterList)
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.WaterResponse{}
	for i := 0; i < len(waterList); i++ {
		responseList = append(responseList, resp.WaterResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(waterList[i].RecordDate, 10))[0:10],
			Water:      waterList[i].WaterValue,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", waterList)
	resp.OK(c, gin.H{
		"waterList": responseList,
	})
}

func (a WaterApi) List(c *gin.Context) {
	waterPageNow, err := strconv.Atoi(c.PostForm("waterPageNow"))
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("waterPageNow", waterPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	waterList, count, err1, err2 := a.waterService.List(user, waterPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("查询数据", waterList)
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.WaterResponse{}
	for i := 0; i < len(waterList); i++ {
		responseList = append(responseList, resp.WaterResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(waterList[i].RecordDate, 10))[0:10],
			Water:      waterList[i].WaterValue,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", waterList)
	resp.OK(c, gin.H{
		"waterList":      responseList,
		"waterPageCount": count,
	})
}

func (a WaterApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	waterList, err := a.waterService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("查询数据", waterList)
	// 处理数据
	var responseList []resp.WaterResponse
	for i := 0; i < len(waterList); i++ {
		responseList = append(responseList, resp.WaterResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(waterList[i].RecordDate, 10))[0:10],
			Water:      waterList[i].WaterValue,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfswaterList", waterList)
	resp.OK(c, gin.H{
		"waterList": responseList,
	})
}

func (a WaterApi) Record(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	// 获取参数 发送13位时间戳 取前十位
	t := c.PostForm("recordDate")[0:10]
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	//fmt.Println("iii", i)
	timestamp := utils.NineStamp(i)
	//fmt.Println("timestamp", timestamp)
	//fmt.Println("strconv.FormatInt(timestamp, 64)", strconv.FormatInt(timestamp, 10))
	in_water := c.PostForm("water")

	// 判断有没有创建记录，没有就创建
	getWater, err := a.waterService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("getWater", getWater)

	// 判断 修改
	if in_water != "" {
		fmt.Println("修改water记录", in_water)
		in_water_int, err := strconv.Atoi(in_water)
		getWater.WaterValue = getWater.WaterValue + in_water_int
		_, err = a.waterService.UpdateWaterById(*getWater)
		if err != nil {
			resp.Error(c, err)
			return
		}
		fmt.Println("i", i)
		resp.OK(c)
	}
}

func (a WaterApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	waterId := c.PostForm("waterId")
	if waterId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	waterId_int, err := strconv.Atoi(waterId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.waterService.Delete(user, models.Water{Id: waterId_int})
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

func (a WaterApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	waterPageNow, err := strconv.Atoi(c.PostForm("waterPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("waterPageNow", waterPageNow)
	// 查询数据
	waterList, count, err1, err2 := a.waterService.Find(user, waterPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	waterResponse := []resp.WaterTableResponse{}
	for i := 0; i < len(waterList); i++ {
		waterResponse = append(waterResponse, resp.WaterTableResponse{
			Id:         waterList[i].Id,
			RecordDate: utils.Stamp10(strconv.FormatInt(waterList[i].RecordDate, 10))[0:10],
			Water:      fmt.Sprintf("%d", waterList[i].WaterValue),
		})
	}
	resp.OK(c, gin.H{
		"waterList":      waterResponse,
		"waterPageCount": count,
	})
}
