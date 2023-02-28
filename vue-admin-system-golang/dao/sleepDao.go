package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type SleepDao struct {
}

func (d SleepDao) List(user *models.User, sleepPageNow int) ([]models.Sleep, int, error, error) {
	var sleepList []models.Sleep
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("active!=0").Desc("record_date").Limit(10, (sleepPageNow-1)*10).Find(&sleepList)
	newSleepList := []models.Sleep{}
	// 反序
	for i := range sleepList {
		newSleepList = append(newSleepList, sleepList[len(sleepList)-1-i])
	}
	sleep := new(models.Sleep)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(sleep)
	fmt.Println("total", total)
	return newSleepList, int(total), err1, err2
}
func (d SleepDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Sleep, int, error, error) {
	var sleepList []models.Sleep
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("active!=0").And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&sleepList)
	newSleepList := []models.Sleep{}
	// 反序
	for i := range sleepList {
		newSleepList = append(newSleepList, sleepList[len(sleepList)-1-i])
	}
	sleep := new(models.Sleep)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(sleep)
	fmt.Println("total", total)
	return newSleepList, int(total), err1, err2
}

func (d SleepDao) Delete(user *models.User, sleep models.Sleep) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", sleep.Id).And("user_id=?", user.Id).Delete(models.Sleep{
		Id:     sleep.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete sleep db id:%d,UserId:%d", sleep.Id, user.Id))

	return i, err
}

func (d SleepDao) Find(user *models.User, sleepPageNow int) ([]models.Sleep, int, error, error) {
	var sleepList []models.Sleep
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (sleepPageNow-1)*10).Find(&sleepList)
	sleep := new(models.Sleep)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(sleep)
	fmt.Println("total", total)
	return sleepList, int(total), err1, err2
}

func (d SleepDao) Sum(user *models.User) ([]models.Sleep, error) {
	var sleepList []models.Sleep
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&sleepList)
	newSleepList := []models.Sleep{}
	// 反序
	for i := range sleepList {
		newSleepList = append(newSleepList, sleepList[len(sleepList)-1-i])
	}
	return newSleepList, err
}

func (d SleepDao) SumAll(user *models.User) ([]models.Sleep, error) {
	var sleepList []models.Sleep
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&sleepList)
	return sleepList, err
}

func (d SleepDao) BaseRecord(user *models.User, timeStamp string) (*models.Sleep, error) {
	sleep := models.Sleep{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&sleep)
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
		effect, err := SqlDB.NewSession().Insert(&models.Sleep{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert sleep db UserId:%d,RecordDate:%d", user.Id, time))

		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Sleep
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d SleepDao) UpdateSleepTimeById(sleep models.Sleep) (int64, error) {
	fmt.Printf("SleepDaoSleepDaoSleepDao %#v", sleep)
	i, err := SqlDB.NewSession().Where("id=?", sleep.Id).Cols("sleep_time").Update(&sleep)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update sleep db id:%d,sleep_time:%d", sleep.Id, sleep.SleepTime))

	return i, nil
}
func (d SleepDao) UpdateGetTimeById(sleep models.Sleep) (int64, error) {
	fmt.Printf("SleepDaoSleepDaoSleepDao %#v", sleep)
	i, err := SqlDB.NewSession().Where("id=?", sleep.Id).Cols("get_time").Update(&sleep)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update sleep db id:%d,get_time:%d", sleep.Id, sleep.GetTime))

	return i, nil
}

func (d SleepDao) UpdateSleepActiveById(sleep models.Sleep) (int64, error) {
	fmt.Printf("SleepDaoSleepDaoSleepDao %#v", sleep)
	i, err := SqlDB.NewSession().Where("id=?", sleep.Id).Cols("active").Update(&sleep)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update sleep db id:%d,active:%d", sleep.Id, sleep.Active))
	return i, nil
}
