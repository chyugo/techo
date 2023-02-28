package dao

import (
	"fmt"
	"vue-admin-system-golang/models"
)

type WakeDao struct {
}

func (d WakeDao) List(user *models.User, wakePageNow int) ([]models.Wake, int, error, error) {
	var wakeList []models.Wake
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (wakePageNow-1)*10).Find(&wakeList)
	newWakeList := []models.Wake{}
	// 反序
	for i := range wakeList {
		newWakeList = append(newWakeList, wakeList[len(wakeList)-1-i])
	}
	wake := new(models.Wake)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(wake)
	fmt.Println("total", total)
	return newWakeList, int(total), err1, err2
}

func (d WakeDao) Sum(user *models.User) ([]models.Wake, error) {
	var wakeList []models.Wake
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&wakeList)
	newWakeList := []models.Wake{}
	// 反序
	for i := range wakeList {
		newWakeList = append(newWakeList, wakeList[len(wakeList)-1-i])
	}
	return newWakeList, err
}

func (d WakeDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Wake, int, error, error) {
	var wakeList []models.Wake
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&wakeList)
	newWakeList := []models.Wake{}
	// 反序
	for i := range wakeList {
		newWakeList = append(newWakeList, wakeList[len(wakeList)-1-i])
	}
	wake := new(models.Wake)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(wake)
	fmt.Println("total", total)
	return newWakeList, int(total), err1, err2
}
