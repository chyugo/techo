package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type WaterDao struct {
}

func (d WaterDao) BaseRecord(user *models.User, timeStamp string) (*models.Water, error) {
	water := models.Water{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&water)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("count，", count)
	time, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		fmt.Println("err", err)
	}
	if count == 0 {
		effect, err := SqlDB.NewSession().Insert(&models.Water{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert WaterDao db RecordDate:%d,UserId:%d", time, user.Id))
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Water
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d WaterDao) UpdateWaterById(water models.Water) (int64, error) {
	i, err := SqlDB.NewSession().Where("id=?", water.Id).Cols("water_value").Update(&water)

	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update WaterDao db Id:%d,water:%d", water.Id, water.WaterValue))

	return i, nil
}

func (d WaterDao) List(user *models.User, waterPageNow int) ([]models.Water, int, error, error) {
	var waterList []models.Water
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (waterPageNow-1)*10).Find(&waterList)
	newWaterList := []models.Water{}
	fmt.Println("waterList", waterList)
	// 反序
	for i := range waterList {
		newWaterList = append(newWaterList, waterList[len(waterList)-1-i])
	}
	water := new(models.Water)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(water)
	fmt.Println("total", total)
	return newWaterList, int(total), err1, err2
}

func (d WaterDao) Sum(user *models.User) ([]models.Water, error) {
	var waterList []models.Water
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&waterList)
	newWeightList := []models.Water{}
	// 反序
	for i := range waterList {
		newWeightList = append(newWeightList, waterList[len(waterList)-1-i])
	}

	return newWeightList, err
}

func (d WaterDao) Delete(user *models.User, water models.Water) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", water.Id).And("user_id=?", user.Id).Delete(models.Water{
		Id:     water.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete WaterDao db Id:%d,user_id:%d", water.Id, user.Id))
	return i, err
}

func (d WaterDao) Find(user *models.User, waterPageNow int) ([]models.Water, int, error, error) {
	var waterList []models.Water
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (waterPageNow-1)*10).Find(&waterList)
	water := new(models.Water)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(water)
	fmt.Println("total", total)
	return waterList, int(total), err1, err2
}

func (d WaterDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Water, int, error, error) {
	var waterList []models.Water
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&waterList)
	newWaterList := []models.Water{}
	fmt.Println("waterList", waterList)
	// 反序
	for i := range waterList {
		newWaterList = append(newWaterList, waterList[len(waterList)-1-i])
	}
	water := new(models.Water)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(water)
	fmt.Println("total", total)
	return newWaterList, int(total), err1, err2
}

func (d WaterDao) SumAll(user *models.User) ([]models.Water, error) {
	var waterList []models.Water
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&waterList)
	return waterList, err
}
