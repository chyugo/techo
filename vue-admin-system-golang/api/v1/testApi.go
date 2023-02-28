package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/service"
)

type TestApi struct {
	heartService  service.HeartService
	phoneService  service.PhoneService
	shitService   service.ShitService
	sleepService  service.SleepService
	userService   service.UserService
	wakeService   service.WakeService
	weightService service.WeightService
}

func (a TestApi) Test1(c *gin.Context) {
	//心脏数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Heart{}.TableName())
	heartList := []models.Heart{}
	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		abnormal := rand.Intn(50)
		normal := 1000 - abnormal
		recordDate := today + int64(86400*i)
		fmt.Println("recorddate", strconv.FormatInt(recordDate, 10))
		heartList = append(heartList, models.Heart{
			UserId:     3,
			RecordDate: recordDate,
			Normal:     normal,
			Abnormal:   abnormal,
		})
	}
	s, err := session.Insert(heartList)
	if err != nil {
		fmt.Println("err", err)
	}
	resp.OK(c, s)
}

func (a TestApi) Test2(c *gin.Context) {
	//体重数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Weight{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		//morning, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*5), 64)
		//noon, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*5), 64)
		//night, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*5), 64)

		morning := rand.Intn(50)
		noon := rand.Intn(50)
		night := rand.Intn(50)
		recordDate := today + int64(86400*i)
		fmt.Println("recorddate", strconv.FormatInt(recordDate, 10))

		_, err := session.Insert(models.Weight{
			UserId:        3,
			RecordDate:    recordDate,
			MorningWeight: 1250 + morning,
			NoonWeight:    1250 + noon,
			NightWeight:   1250 + night,
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)
}

func (a TestApi) Test3(c *gin.Context) {
	//拉屎数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Shit{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		recordDate := today + int64(86400*i)
		shitTime := recordDate + int64(rand.Intn(21600))
		fmt.Println("recorddate", strconv.FormatInt(recordDate, 10))

		_, err := session.Insert(models.Shit{
			UserId:     3,
			ShitTime:   shitTime,
			RecordDate: recordDate,
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)
}

func (a TestApi) Test5(c *gin.Context) {
	//sleep数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Sleep{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		recordDate := today + int64(86400*i)
		get := recordDate + int64(rand.Intn(21600))
		sleep := recordDate + int64(rand.Intn(51600))
		fmt.Println("recorddate", strconv.FormatInt(recordDate, 10))

		_, err := session.Insert(models.Sleep{
			UserId:     3,
			GetTime:    get,
			SleepTime:  sleep,
			RecordDate: recordDate,
			Active:     360 + rand.Intn(135),
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)
}

func (a TestApi) Test6(c *gin.Context) {
	//phone数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Phone{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		recordDate := today + int64(86400*i)
		_, err := session.Insert(models.Phone{
			UserId:      3,
			RecordDate:  recordDate,
			PhoneMinute: rand.Intn(360),
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)

}

func (a TestApi) Test7(c *gin.Context) {
	//喝水数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Phone{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		recordDate := today + int64(86400*i)
		_, err := session.Insert(models.Water{
			UserId:     3,
			RecordDate: recordDate,
			WaterValue: rand.Intn(2500),
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)

}
func (a TestApi) Test8(c *gin.Context) {
	//反酸数据填充
	rand.Seed(time.Now().UnixNano())
	session := dao.SqlDB.NewSession().Table(models.Phone{}.TableName())

	var today int64
	today = 1672790400
	for i := 0; i < 700; i++ {
		recordDate := today + int64(86400*i)
		_, err := session.Insert(models.Acid{
			UserId:     3,
			RecordDate: recordDate,
			AcidValue:  rand.Intn(60),
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
	resp.OK(c)

}
