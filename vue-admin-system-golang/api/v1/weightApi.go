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

type WeightApi struct {
	weightService service.WeightService
}

func (a WeightApi) SearchMonth(c *gin.Context) {
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
	weightList, count, err1, err2 := a.weightService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("查询数据", weightList)
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.WeightResponse{}
	for i := 0; i < len(weightList); i++ {
		responseList = append(responseList, resp.WeightResponse{
			RecordDate:    utils.Stamp10(strconv.FormatInt(weightList[i].RecordDate, 10))[0:10],
			MorningWeight: float32(weightList[i].MorningWeight) / 10,
			NoonWeight:    float32(weightList[i].NoonWeight) / 10,
			NightWeight:   float32(weightList[i].NightWeight) / 10,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", weightList)
	resp.OK(c, gin.H{
		"weightList": responseList,
	})
}

func (a WeightApi) List(c *gin.Context) {
	weightPageNow, err := strconv.Atoi(c.PostForm("weightPageNow"))
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("pagenowaaaa", weightPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	weightList, count, err1, err2 := a.weightService.List(user, weightPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("查询数据", weightList)
	fmt.Println("count", count)
	// 处理数据
	var responseList []resp.WeightResponse
	for i := 0; i < len(weightList); i++ {
		responseList = append(responseList, resp.WeightResponse{
			RecordDate:    utils.Stamp10(strconv.FormatInt(weightList[i].RecordDate, 10))[0:10],
			MorningWeight: float32(weightList[i].MorningWeight) / 10,
			NoonWeight:    float32(weightList[i].NoonWeight) / 10,
			NightWeight:   float32(weightList[i].NightWeight) / 10,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", weightList)
	resp.OK(c, gin.H{
		"weightList":      responseList,
		"weightPageCount": count,
	})
}

func (a WeightApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	weightList, err := a.weightService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("查询数据", weightList)
	// 处理数据
	var responseList []resp.WeightResponse
	for i := 0; i < len(weightList); i++ {
		responseList = append(responseList, resp.WeightResponse{
			RecordDate:    utils.Stamp10(strconv.FormatInt(weightList[i].RecordDate, 10))[0:10],
			MorningWeight: float32(weightList[i].MorningWeight) / 10,
			NoonWeight:    float32(weightList[i].NoonWeight) / 10,
			NightWeight:   float32(weightList[i].NightWeight) / 10,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", weightList)
	resp.OK(c, gin.H{
		"weightList": responseList,
	})
}

func (a WeightApi) Record(c *gin.Context) {
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
	morning := c.PostForm("morningWeight")
	noon := c.PostForm("noonWeight")
	night := c.PostForm("nightWeight")

	// 判断有没有创建记录，没有就创建
	weight, err := a.weightService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("monringh", morning, "morning", noon, "night", night)
	fmt.Println("weight", weight)

	// 判断 修改
	if morning != "" {
		fmt.Println("修改 morning记录", morning)
		w, err := strconv.Atoi(morning)
		weight.MorningWeight = w
		_, err = a.weightService.UpdateMorningById(*weight)
		if err != nil {
			resp.Error(c, err)
			return
		}
		resp.OK(c)
	}

	// 判断 修改
	if noon != "" {
		fmt.Println("修改 noon记录", noon)
		w, err := strconv.Atoi(noon)
		weight.NoonWeight = w
		_, err = a.weightService.UpdateNoonById(*weight)
		if err != nil {
			resp.Error(c, err)
			return
		}
		fmt.Println("i", i)
		resp.OK(c)

	}

	// 判断 修改
	if night != "" {
		fmt.Println("修改night记录", night)
		w, err := strconv.Atoi(night)
		weight.NightWeight = w
		_, err = a.weightService.UpdateNightById(*weight)
		if err != nil {
			resp.Error(c, err)
			return
		}
		fmt.Println("i", i)
		resp.OK(c)
	}
}

func (a WeightApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	weightId := c.PostForm("weightId")
	if weightId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	weightId_int, err := strconv.Atoi(weightId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.weightService.Delete(user, models.Weight{Id: weightId_int})
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

func (a WeightApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	weightPageNow, err := strconv.Atoi(c.PostForm("weightPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("weightPageNow", weightPageNow)
	// 查询数据
	weightList, count, err1, err2 := a.weightService.Find(user, weightPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	weightResponse := []resp.WeightTableResponse{}
	for i := 0; i < len(weightList); i++ {
		weightResponse = append(weightResponse, resp.WeightTableResponse{
			Id:            weightList[i].Id,
			RecordDate:    utils.Stamp10(strconv.FormatInt(weightList[i].RecordDate, 10))[0:10],
			MorningWeight: fmt.Sprintf("%.1f斤", float64(weightList[i].MorningWeight)/10),
			NoonWeight:    fmt.Sprintf("%.1f斤", float64(weightList[i].NoonWeight)/10),
			NightWeight:   fmt.Sprintf("%.1f斤", float64(weightList[i].NightWeight)/10),
		})
	}
	resp.OK(c, gin.H{
		"weightList":      weightResponse,
		"weightPageCount": count,
	})
}
