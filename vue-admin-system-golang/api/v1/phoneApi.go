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

type PhoneApi struct {
	phoneService service.PhoneService
}

func (a PhoneApi) List(c *gin.Context) {
	phonePageNow, err := strconv.Atoi(c.PostForm("phonePageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", phonePageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	phoneList, count, err1, err2 := a.phoneService.List(user, phonePageNow)
	fmt.Printf("%#vphonelist", phoneList)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	var responseList []resp.PhoneResponse
	for i := 0; i < len(phoneList); i++ {
		responseList = append(responseList, resp.PhoneResponse{
			RecordDate:  utils.Stamp10(strconv.FormatInt(phoneList[i].RecordDate, 10))[0:10],
			PhoneMinute: phoneList[i].PhoneMinute,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", responseList)
	resp.OK(c, gin.H{
		"phoneList":      responseList,
		"phonePageCount": count,
	})

}

func (a PhoneApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	phoneList, err := a.phoneService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 处理数据
	var responseList []resp.PhoneResponse
	for i := 0; i < len(phoneList); i++ {
		responseList = append(responseList, resp.PhoneResponse{
			RecordDate:  utils.Stamp10(strconv.FormatInt(phoneList[i].RecordDate, 10))[0:10],
			PhoneMinute: phoneList[i].PhoneMinute,
		})
	}
	fmt.Println("adfasfsdfsdfadsfasdfsweightList", responseList)
	resp.OK(c, gin.H{
		"phoneList": responseList,
	})

}

func (a PhoneApi) Record(c *gin.Context) {
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
	timestamp := utils.NineStamp(i)

	phoneHour := c.PostForm("phoneHour")
	phoneMinute := c.PostForm("phoneMinute")
	phoneHour_int, _ := strconv.Atoi(phoneHour)
	phoneMinute_int, _ := strconv.Atoi(phoneMinute)

	phoneAllMinute := phoneHour_int*60 + phoneMinute_int
	// 判断有没有创建记录，没有就创建
	phone, err := a.phoneService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}

	// 判断 修改
	fmt.Println("修改 phone记录", phone)
	phone.PhoneMinute = phoneAllMinute
	_, err = a.phoneService.UpdatePhoneById(*phone)
	if err != nil {
		resp.Error(c, err)
		return

	}
	resp.OK(c)

}

func (a PhoneApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	phoneId := c.PostForm("phoneId")
	if phoneId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	phoneId_int, err := strconv.Atoi(phoneId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.phoneService.Delete(user, models.Phone{Id: phoneId_int})
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

func (a PhoneApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	phonePageNow, err := strconv.Atoi(c.PostForm("phonePageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("pagenow", phonePageNow)
	// 查询数据
	phoneList, count, err1, err2 := a.phoneService.Find(user, phonePageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	phoneResponse := []resp.PhoneTableResponse{}
	for i := 0; i < len(phoneList); i++ {
		phoneResponse = append(phoneResponse, resp.PhoneTableResponse{
			Id:          phoneList[i].Id,
			RecordDate:  utils.Stamp10(strconv.FormatInt(phoneList[i].RecordDate, 10))[0:10],
			PhoneMinute: utils.MinuteToHourMinute2(phoneList[i].PhoneMinute),
		})
	}
	resp.OK(c, gin.H{
		"phoneList":      phoneResponse,
		"phonePageCount": count,
	})
}

func (a PhoneApi) SearchMonth(c *gin.Context) {
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
	phoneList, count, err1, err2 := a.phoneService.SearchMonth(user, thisMonth, nextMonth)
	fmt.Printf("%#vphonelist", phoneList)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.PhoneResponse{}
	for i := 0; i < len(phoneList); i++ {
		responseList = append(responseList, resp.PhoneResponse{
			RecordDate:  utils.Stamp10(strconv.FormatInt(phoneList[i].RecordDate, 10))[0:10],
			PhoneMinute: phoneList[i].PhoneMinute,
		})
	}
	resp.OK(c, gin.H{
		"phoneList": responseList,
	})

}
