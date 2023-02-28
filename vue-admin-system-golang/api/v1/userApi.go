package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
	"vue-admin-system-golang/utils"
	"vue-admin-system-golang/utils/logger"
)

type UserApi struct {
	userService   service.UserService
	heartService  service.HeartService
	phoneService  service.PhoneService
	shitService   service.ShitService
	sleepService  service.SleepService
	weightService service.WeightService
	acidService   service.AcidService
	waterService  service.WaterService
}

func (a UserApi) DoLogin(c *gin.Context) {
	passWord := c.PostForm("passWord")
	userName := c.PostForm("userName")
	secreat := utils.CreatePassWordSecret(passWord)

	if userName == "" || passWord == "" {
		resp.Error(c, "缺少参数")
		return
	}
	user := &models.User{
		UserName: userName,
		PassWord: secreat,
	}
	finduser, err := a.userService.Login(user)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "找不到用户" {
			resp.Error(c, -3, "找不到用户")
			return
		}
		resp.Error(c, err.Error())
		return
	}
	if user.PassWord != finduser.PassWord {
		resp.Error(c, "密码错误")
		return
	}
	//验证通过 生成cookie和session
	cookie := utils.CreateCookie(string(finduser.Id))
	session, err := json.Marshal(models.User{
		Id:       finduser.Id,
		UserName: finduser.UserName,
	})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	_, err = dao.RedisDB.Set(context.Background(), cookie, session, time.Hour*24*30).Result()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	resp.OK(c, gin.H{
		"cookie":  cookie,
		"outTime": 60 * 60 * 24 * 30,
	})

}

func (a UserApi) DoRegist(c *gin.Context) {
	registName := c.PostForm("registName")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")
	if password1 != password2 {
		resp.Error(c, "两次输入的密码不一致")
		return
	}
	if password1 == "" || registName == "" {
		resp.Error(c, "缺少参数")
		return
	}

	secreat := utils.CreatePassWordSecret(password1)
	user := &models.User{
		UserName: registName,
		PassWord: secreat,
	}

	count, err := a.userService.Regist(user)
	if int(count) != 0 {
		resp.Error(c, "有同名的用户")
		return
	}
	if err != nil {
		resp.Error(c, err)
		return
	}
	resp.OK(c)

}

func (a UserApi) NotOutTime(c *gin.Context) {
	cookie := c.PostForm("cookie")
	fmt.Println("cookie", cookie)
	_, err := dao.RedisDB.Get(context.Background(), cookie).Result()
	if err != nil {
		resp.OK(c, gin.H{
			"notOutTime": false,
		})
		return
	}
	resp.OK(c, gin.H{
		"notOutTime": true,
	})

}

func (a UserApi) Export(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	heartList, err := a.heartService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export heartList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	phoneList, err := a.phoneService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export phoneList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	shitList, err := a.shitService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export shitList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	sleepList, err := a.sleepService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export sleepList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	weightList, err := a.weightService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export weightList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	acidList, err := a.acidService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export acidList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}
	waterList, err := a.waterService.SumAll(user)
	if err != nil {
		logger.Logger.Error("export waterList error.", err)
		fmt.Println("err", err)
		resp.Error(c)
		return
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	//heart表
	start := 3
	f.SetCellValue("Sheet1", "A1", "heart")
	f.SetCellValue("Sheet1", "A2", "record_date")
	f.SetCellValue("Sheet1", "B2", "normal")
	f.SetCellValue("Sheet1", "C2", "abnormal")
	for i := 0; i < len(heartList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(heartList[i].RecordDate, 10))[0:10]
		Normal := fmt.Sprintf("%.1f%%", float64(heartList[i].Normal)/10)
		Abnormal := fmt.Sprintf("%.1f%%", float64(heartList[i].Abnormal)/10)
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", start), Normal)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", start), Abnormal)
		start++
	}

	// phone表
	start = 3
	f.SetCellValue("Sheet1", "E1", "phone")
	f.SetCellValue("Sheet1", "E2", "record_date")
	f.SetCellValue("Sheet1", "F2", "PhoneMinute")
	for i := 0; i < len(phoneList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(phoneList[i].RecordDate, 10))[0:10]
		PhoneMinute := utils.MinuteToHourMinute2(phoneList[i].PhoneMinute)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", start), PhoneMinute)
		start++
	}

	// shit
	start = 3
	f.SetCellValue("Sheet1", "H1", "shit")
	f.SetCellValue("Sheet1", "H2", "record_date")
	f.SetCellValue("Sheet1", "I2", "ShitTime")

	for i := 0; i < len(shitList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(shitList[i].RecordDate, 10))[0:10]
		ShitTime := utils.StamptoString(shitList[i].ShitTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", start), ShitTime)
		start++
	}

	// sleep
	start = 3
	f.SetCellValue("Sheet1", "K1", "sleep")
	f.SetCellValue("Sheet1", "K2", "record_date")
	f.SetCellValue("Sheet1", "L2", "SleepTime")
	f.SetCellValue("Sheet1", "M2", "GetTime")
	f.SetCellValue("Sheet1", "N2", "Active")

	for i := 0; i < len(sleepList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(sleepList[i].RecordDate, 10))[0:10]
		SleepTime := utils.StamptoString(sleepList[i].SleepTime)
		GetTime := utils.StamptoString(sleepList[i].GetTime)
		Active := utils.MinuteToHourMinute2(sleepList[i].Active)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", start), SleepTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", start), GetTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", start), Active)
		start++
	}

	// weight
	start = 3
	f.SetCellValue("Sheet1", "P1", "weight")
	f.SetCellValue("Sheet1", "P2", "record_date")
	f.SetCellValue("Sheet1", "Q2", "MorningWeight")
	f.SetCellValue("Sheet1", "R2", "NoonWeight")
	f.SetCellValue("Sheet1", "S2", "NightWeight")

	for i := 0; i < len(weightList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(weightList[i].RecordDate, 10))[0:10]
		MorningWeight := fmt.Sprintf("%.1f斤", float64(weightList[i].MorningWeight)/10)
		NoonWeight := fmt.Sprintf("%.1f斤", float64(weightList[i].NoonWeight)/10)
		NightWeight := fmt.Sprintf("%.1f斤", float64(weightList[i].NightWeight)/10)
		f.SetCellValue("Sheet1", fmt.Sprintf("P%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%d", start), MorningWeight)
		f.SetCellValue("Sheet1", fmt.Sprintf("R%d", start), NoonWeight)
		f.SetCellValue("Sheet1", fmt.Sprintf("S%d", start), NightWeight)
		start++
	}

	// water
	start = 3
	f.SetCellValue("Sheet1", "U1", "water")
	f.SetCellValue("Sheet1", "U2", "record_date")
	f.SetCellValue("Sheet1", "V2", "WaterValue")

	for i := 0; i < len(waterList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(waterList[i].RecordDate, 10))[0:10]
		WaterValue := fmt.Sprintf("%.d", waterList[i].WaterValue)
		f.SetCellValue("Sheet1", fmt.Sprintf("U%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("V%d", start), WaterValue)
		start++
	}

	// acid
	start = 3
	f.SetCellValue("Sheet1", "X1", "acid")
	f.SetCellValue("Sheet1", "X2", "record_date")
	f.SetCellValue("Sheet1", "Y2", "AcidValue")

	for i := 0; i < len(acidList); i++ {
		RecordDate := utils.Stamp10(strconv.FormatInt(acidList[i].RecordDate, 10))[0:10]
		WaterValue := fmt.Sprintf("%.d", acidList[i].AcidValue)
		f.SetCellValue("Sheet1", fmt.Sprintf("X%d", start), RecordDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%d", start), WaterValue)
		start++
	}

	// Set active sheet of the workbook.
	//设置列宽
	f.SetColWidth("Sheet1", "A", "A", 11)
	f.SetColWidth("Sheet1", "E", "E", 11)
	f.SetColWidth("Sheet1", "H", "H", 11)
	f.SetColWidth("Sheet1", "K", "K", 11)
	f.SetColWidth("Sheet1", "F", "F", 11)
	f.SetColWidth("Sheet1", "N", "N", 11)
	f.SetColWidth("Sheet1", "P", "P", 11)
	f.SetColWidth("Sheet1", "U", "U", 11)
	f.SetColWidth("Sheet1", "X", "X", 11)

	f.SetColWidth("Sheet1", "Q", "Q", 13)
	f.SetColWidth("Sheet1", "R", "R", 13)
	f.SetColWidth("Sheet1", "S", "S", 13)
	f.SetColWidth("Sheet1", "V", "V", 13)
	f.SetColWidth("Sheet1", "Y", "Y", 13)

	f.SetColWidth("Sheet1", "I", "I", 18)
	f.SetColWidth("Sheet1", "L", "L", 18)
	f.SetColWidth("Sheet1", "M", "M", 18)
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	path := "./Excels/"
	url := fmt.Sprintf("%s%s.xlsx", path, utils.CreateFilename(user.UserName))
	if err := f.SaveAs(url); err != nil {
		fmt.Println(err)
		logger.Logger.Error("create file err.", err)
		resp.Error(c)
		return
	}
	fmt.Println(url)
	c.File(url)
}

func (a UserApi) ResetPassword(c *gin.Context) {
	// 鉴权
	user, err := utils.GetUserByRedis(c)
	if err != nil {
		resp.Error(c, err)
		return
	}
	if user.Id == 3 {
		resp.Error(c, "体验账户不可修改密码")
		return
	}
	passWord1 := c.PostForm("passWord1")
	passWord2 := c.PostForm("passWord2")
	if passWord1 != passWord2 {
		resp.Error(c, "两次输入的密码不一致")
		return
	}
	if passWord1 == "" {
		resp.Error(c, "缺少参数")
		return
	}

	secreat := utils.CreatePassWordSecret(passWord1)
	updateUser := &models.User{
		Id:       user.Id,
		PassWord: secreat,
	}
	_, err = a.userService.UpdatePassword(updateUser)
	if err != nil {
		fmt.Println("err")
		logger.Logger.Error("update password err.", err)
		resp.Error(c, err)
		return
	}
	resp.OK(c)

}

func (a UserApi) SentSuggest(c *gin.Context) {
	suggestContent := c.PostForm("suggestContent")
	if suggestContent == "" {
		resp.Error(c, "获取参数错误")
		return
	}

	_, err := a.userService.SaveSuggest(models.Suggest{Content: suggestContent})
	if err != nil {
		fmt.Println("err")
		logger.Logger.Error("update password err.", err)
		resp.Error(c, err)
		return
	}
	resp.OK(c)

}
