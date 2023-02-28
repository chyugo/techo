package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type ShitDao struct {
}

func (d ShitDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Shit, int, error, error) {
	var shitList []models.Shit
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&shitList)
	newShitList := []models.Shit{}
	// 反序
	for i := range shitList {
		newShitList = append(newShitList, shitList[len(shitList)-1-i])
	}
	shit := new(models.Shit)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(shit)
	fmt.Println("total", total)
	return newShitList, int(total), err1, err2
}

func (d ShitDao) List(user *models.User, shitPageNow int) ([]models.Shit, int, error, error) {
	var shitList []models.Shit
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (shitPageNow-1)*10).Find(&shitList)
	newShitList := []models.Shit{}
	// 反序
	for i := range shitList {
		newShitList = append(newShitList, shitList[len(shitList)-1-i])
	}
	shit := new(models.Shit)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(shit)
	fmt.Println("total", total)
	return newShitList, int(total), err1, err2
}
func (d ShitDao) Sum(user *models.User) ([]models.Shit, error) {
	var shitList []models.Shit
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&shitList)
	newShitList := []models.Shit{}
	// 反序
	for i := range shitList {
		newShitList = append(newShitList, shitList[len(shitList)-1-i])
	}
	return newShitList, err
}

func (d ShitDao) SumAll(user *models.User) ([]models.Shit, error) {
	var shitList []models.Shit
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&shitList)
	return shitList, err
}

func (d ShitDao) BaseRecord(user *models.User, timeStamp string) (*models.Shit, error) {
	shit := models.Shit{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&shit)
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
		effect, err := SqlDB.NewSession().Insert(&models.Shit{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert shit db UserId:%d,RecordDate:%d", user.Id, time))

		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Shit
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d ShitDao) UpdateShitById(shit models.Shit) (int64, error) {
	fmt.Printf("shitshitshit %#v", shit)
	i, err := SqlDB.NewSession().Where("id=?", shit.Id).Cols("shit_time").Update(&shit)

	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update shit db id:%d,shit_time:%d", shit.Id, shit.ShitTime))

	return i, nil
}

func (d ShitDao) Delete(user *models.User, shit models.Shit) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", shit.Id).And("user_id=?", user.Id).Delete(models.Shit{
		Id:     shit.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete shit db id:%d,UserId:%d", shit.Id, user.Id))

	return i, err
}

func (d ShitDao) Find(user *models.User, shitPageNow int) ([]models.Shit, int, error, error) {
	var shitList []models.Shit
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (shitPageNow-1)*10).Find(&shitList)
	shit := new(models.Shit)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(shit)
	fmt.Println("total", total)
	return shitList, int(total), err1, err2
}
