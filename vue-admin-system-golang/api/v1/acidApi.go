package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/models/request"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
)

type AcidApi struct {
	acidService service.AcidService
}

func (a AcidApi) SearchMonth(c *gin.Context) {
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
	acidList, count, err1, err2 := a.acidService.SearchMonth(user, thisMonth, nextMonth)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	//fmt.Println("查询数据", acidList)
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.AcidResponse{}
	for i := 0; i < len(acidList); i++ {
		responseList = append(responseList, resp.AcidResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10],
			Acid:       acidList[i].AcidValue,
		})
	}
	//fmt.Println("adfasfsdfsdfadsfasdfsweightList", acidList)
	resp.OK(c, gin.H{
		"acidList": responseList,
	})
}

func (a AcidApi) List(c *gin.Context) {
	acidPageNow, err := strconv.Atoi(c.PostForm("acidPageNow"))
	if err != nil {
		fmt.Println("err", err)
		resp.Error(c, err)
		return
	}
	fmt.Println("acidPageNow", acidPageNow)
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	acidList, count, err1, err2 := a.acidService.List(user, acidPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	//fmt.Println("查询数据", acidList)
	fmt.Println("count", count)
	// 处理数据
	responseList := []resp.AcidResponse{}
	for i := 0; i < len(acidList); i++ {
		responseList = append(responseList, resp.AcidResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10],
			Acid:       acidList[i].AcidValue,
		})
	}
	//fmt.Println("adfasfsdfsdfadsfasdfsweightList", acidList)
	resp.OK(c, gin.H{
		"acidList":      responseList,
		"acidPageCount": count,
	})
}

func (a AcidApi) Sum(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	// 查询数据
	acidList, err := a.acidService.Sum(user)
	if err != nil {
		resp.Error(c, err)
		return
	}
	//fmt.Println("查询数据", acidList)
	// 处理数据
	var responseList []resp.AcidResponse
	for i := 0; i < len(acidList); i++ {
		responseList = append(responseList, resp.AcidResponse{
			RecordDate: utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10],
			Acid:       acidList[i].AcidValue,
		})
	}
	//fmt.Println("adfasfsdfsdfadsfasdfsacidList", acidList)
	resp.OK(c, gin.H{
		"acidList": responseList,
	})
}

func (a AcidApi) Record(c *gin.Context) {
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
	judge, err := strconv.Atoi(c.PostForm("judge"))
	if err != nil {
		resp.ParamError(c)
		return
	}
	//fmt.Println("iii", i)
	timestamp := utils.NineStamp(i)
	//fmt.Println("timestamp", timestamp)
	//fmt.Println("strconv.FormatInt(timestamp, 64)", strconv.FormatInt(timestamp, 10))
	in_acid := c.PostForm("acid")

	// 判断有没有创建记录，没有就创建
	getAcid, err := a.acidService.BaseRecord(user, strconv.FormatInt(timestamp, 10))
	if err != nil {
		resp.Error(c, err)
		return
	}
	fmt.Println("getAcid", getAcid)

	// 判断 修改
	if in_acid != "" {
		fmt.Println("修改acid记录", in_acid)
		in_acid_int, err := strconv.Atoi(in_acid)
		getAcid.AcidValue = in_acid_int
		getAcid.Judge = judge
		_, err = a.acidService.UpdateAcidById(*getAcid)
		if err != nil {
			resp.Error(c, err)
			return
		}
		fmt.Println("i", i)
		resp.OK(c)
	}
}

func (a AcidApi) Delete(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	acidId := c.PostForm("acidId")
	if acidId == "" {
		resp.Error(c, "缺少参数")
		return
	}

	acidId_int, err := strconv.Atoi(acidId)
	if err != nil {
		resp.Error(c, "参数转换错误")
		return
	}

	i, err := a.acidService.Delete(user, models.Acid{Id: acidId_int})
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

func (a AcidApi) Find(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	acidPageNow, err := strconv.Atoi(c.PostForm("acidPageNow"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("acidPageNow", acidPageNow)
	// 查询数据
	acidList, count, err1, err2 := a.acidService.Find(user, acidPageNow)
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	acidResponse := []resp.AcidTableResponse{}
	for i := 0; i < len(acidList); i++ {
		acidResponse = append(acidResponse, resp.AcidTableResponse{
			Id:         acidList[i].Id,
			RecordDate: utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10],
			Acid:       fmt.Sprintf("%d", acidList[i].AcidValue),
			Judge:      acidList[i].Judge,
		})
	}
	resp.OK(c, gin.H{
		"acidList":      acidResponse,
		"acidPageCount": count,
	})
}

func (a AcidApi) AcidDate(c *gin.Context) {
	getModel := request.AcidDateModel{}
	if err := c.Bind(&getModel); err != nil {
		resp.ParamError(c)
		return
	}
	fmt.Printf("%#v", getModel)

	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}

	thisMonth := time.Unix(utils.StampConv2(time.UnixMilli(getModel.Month).Unix(), 1, 8, 0), 0)
	nextMonth := thisMonth.AddDate(0, 1, 0)
	fmt.Println(thisMonth.Unix(), nextMonth.Unix())

	// 查询数据
	acidList, count, err1, err2 := a.acidService.SearchMonth(user, thisMonth.Unix(), nextMonth.Unix())
	if err1 != nil || err2 != nil {
		resp.Error(c, err1, err2)
		return
	}
	fmt.Println("count", count)
	// 处理数据
	var responseList []resp.AcidDateResponse
	for i := 0; i < len(acidList); i++ {
		var notice string
		var type_s string

		if acidList[i].Judge == 1 {
			notice = "不涨"
			type_s = "safe"
		}
		if acidList[i].Judge == 2 {
			notice = "小胀"
			type_s = "important"
		}
		if acidList[i].Judge == 3 {
			notice = "中胀"
			type_s = "important"
		}
		if acidList[i].Judge == 4 {
			notice = "大胀"
			type_s = "important"
		}
		responseList = append(responseList, resp.AcidDateResponse{
			WorkingDay: utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10],
			Content:    []resp.AcidContent{resp.AcidContent{Notice: notice, Type: type_s}},
		})
	}
	resp.OK(c, responseList)
}
