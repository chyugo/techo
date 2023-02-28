package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type WeightDao struct {
}

func (d WeightDao) List(user *models.User, weightPageNow int) ([]models.Weight, int, error, error) {
	var weightList []models.Weight
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (weightPageNow-1)*10).Find(&weightList)
	newWeightList := []models.Weight{}
	fmt.Println("weightList", weightList)
	// 反序
	for i := range weightList {
		newWeightList = append(newWeightList, weightList[len(weightList)-1-i])
	}
	weight := new(models.Weight)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(weight)
	fmt.Println("total", total)
	return newWeightList, int(total), err1, err2
}

func (d WeightDao) Sum(user *models.User) ([]models.Weight, error) {
	var weightList []models.Weight
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&weightList)
	newWeightList := []models.Weight{}
	// 反序
	for i := range weightList {
		newWeightList = append(newWeightList, weightList[len(weightList)-1-i])
	}

	return newWeightList, err
}

func (d WeightDao) SumAll(user *models.User) ([]models.Weight, error) {
	var weightList []models.Weight
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&weightList)
	return weightList, err
}

func (d WeightDao) BaseRecord(user *models.User, timeStamp string) (*models.Weight, error) {
	weight := models.Weight{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&weight)
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
		effect, err := SqlDB.NewSession().Insert(&models.Weight{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert WeightDao db RecordDate:%d,UserId:%d", time, user.Id))
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Weight
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d WeightDao) UpdateMorningById(weight models.Weight) (int64, error) {
	fmt.Printf("weightweightweight %#v", weight)
	i, err := SqlDB.NewSession().Where("id=?", weight.Id).Cols("morning_weight").Update(&weight)

	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update WeightDao db Id:%d,morning_weight:%d", weight.Id, weight.MorningWeight))

	return i, nil
}

func (d WeightDao) UpdateNoonById(weight models.Weight) (int64, error) {
	i, err := SqlDB.NewSession().Where("id=?", weight.Id).Cols("noon_weight").Update(&weight)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update WeightDao db Id:%d,noon_weight:%d", weight.Id, weight.NoonWeight))
	return i, nil
}

func (d WeightDao) UpdateNightById(weight models.Weight) (int64, error) {
	i, err := SqlDB.NewSession().Where("id=?", weight.Id).Cols("night_weight").Update(&weight)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update WeightDao db Id:%d,night_weight:%d", weight.Id, weight.NightWeight))

	return i, nil
}

func (d WeightDao) Delete(user *models.User, weight models.Weight) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", weight.Id).And("user_id=?", user.Id).Delete(models.Weight{
		Id:     weight.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete WeightDao db Id:%d,user_id:%d", weight.Id, user.Id))
	return i, err
}

func (d WeightDao) Find(user *models.User, weightPageNow int) ([]models.Weight, int, error, error) {
	var weightList []models.Weight
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (weightPageNow-1)*10).Find(&weightList)
	weight := new(models.Weight)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(weight)
	fmt.Println("total", total)
	return weightList, int(total), err1, err2
}

func (d WeightDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Weight, int, error, error) {
	var weightList []models.Weight
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&weightList)
	newWeightList := []models.Weight{}
	fmt.Println("weightList", weightList)
	// 反序
	for i := range weightList {
		newWeightList = append(newWeightList, weightList[len(weightList)-1-i])
	}
	weight := new(models.Weight)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(weight)
	fmt.Println("total", total)
	return newWeightList, int(total), err1, err2
}
