package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type HeartDao struct {
}

func (d HeartDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Heart, int, error, error) {
	var heartList []models.Heart
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&heartList)
	newHeartList := []models.Heart{}
	// 反序
	for i := range heartList {
		newHeartList = append(newHeartList, heartList[len(heartList)-1-i])
	}
	heart := new(models.Heart)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(heart)
	fmt.Println("total", total)
	return newHeartList, int(total), err1, err2
}

func (d HeartDao) List(user *models.User, heartPageNow int) ([]models.Heart, int, error, error) {
	var heartList []models.Heart
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (heartPageNow-1)*10).Find(&heartList)
	newHeartList := []models.Heart{}
	// 反序
	for i := range heartList {
		newHeartList = append(newHeartList, heartList[len(heartList)-1-i])
	}
	heart := new(models.Heart)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(heart)
	fmt.Println("total", total)
	return newHeartList, int(total), err1, err2
}

func (d HeartDao) Sum(user *models.User) ([]models.Heart, error) {
	var heartList []models.Heart
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&heartList)
	newHeartList := []models.Heart{}
	// 反序
	for i := range heartList {
		newHeartList = append(newHeartList, heartList[len(heartList)-1-i])
	}
	return newHeartList, err1
}

func (d HeartDao) SumAll(user *models.User) ([]models.Heart, error) {
	var heartList []models.Heart
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&heartList)
	return heartList, err1
}

func (d HeartDao) BaseRecord(user *models.User, timeStamp string) (*models.Heart, error) {
	heart := models.Heart{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&heart)
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
		effect, err := SqlDB.NewSession().Insert(&models.Heart{
			UserId:     user.Id,
			RecordDate: time,
		})
		logger.Logger.Info(fmt.Sprintf("Insert heart db data.Id:%d,RecordDate:%d", user.Id, time))
		fmt.Println("没有对应记录，执行插入", effect)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Heart
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d HeartDao) UpdateHeartById(heart models.Heart) (int64, error) {
	fmt.Printf("heartheart %#v", heart)
	i, err := SqlDB.NewSession().Where("id=?", heart.Id).Cols("normal", "abnormal").Update(&heart)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update heart db data.Id:%d,normal:%d,abnormal:%d", heart.Id, heart.Normal, heart.Abnormal))

	return i, nil
}

func (d HeartDao) Delete(user *models.User, heart models.Heart) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", heart.Id).And("user_id=?", user.Id).Delete(models.Heart{
		Id:     heart.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete heart db data.Id:%d,UserId:%d", heart.Id, user.Id))
	return i, err
}

func (d HeartDao) Find(user *models.User, heartPageNow int) ([]models.Heart, int, error, error) {
	var heartList []models.Heart
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (heartPageNow-1)*10).Find(&heartList)
	heart := new(models.Heart)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(heart)
	fmt.Println("total", total)
	return heartList, int(total), err1, err2
}
