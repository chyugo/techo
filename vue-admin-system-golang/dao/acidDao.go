package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type AcidDao struct {
}

func (d AcidDao) BaseRecord(user *models.User, timeStamp string) (*models.Acid, error) {
	acid := models.Acid{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&acid)
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
		effect, err := SqlDB.NewSession().Insert(&models.Acid{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert AcidDao db RecordDate:%d,UserId:%d", time, user.Id))
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Acid
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d AcidDao) UpdateAcidById(acid models.Acid) (int64, error) {
	i, err := SqlDB.NewSession().Where("id=?", acid.Id).Cols("acid_value", "judge").Update(&acid)

	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update AcidDao db Id:%d,acid:%d,judge:%d", acid.Id, acid.AcidValue, acid.Judge))

	return i, nil
}

func (d AcidDao) List(user *models.User, acidPageNow int) ([]models.Acid, int, error, error) {
	var acidList []models.Acid
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (acidPageNow-1)*10).Find(&acidList)
	newAcidList := []models.Acid{}
	fmt.Println("acidList", acidList)
	// 反序
	for i := range acidList {
		newAcidList = append(newAcidList, acidList[len(acidList)-1-i])
	}
	acid := new(models.Acid)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(acid)
	fmt.Println("total", total)
	return newAcidList, int(total), err1, err2
}

func (d AcidDao) Sum(user *models.User) ([]models.Acid, error) {
	var acidList []models.Acid
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&acidList)
	newWeightList := []models.Acid{}
	// 反序
	for i := range acidList {
		newWeightList = append(newWeightList, acidList[len(acidList)-1-i])
	}

	return newWeightList, err
}

func (d AcidDao) Delete(user *models.User, acid models.Acid) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", acid.Id).And("user_id=?", user.Id).Delete(models.Acid{
		Id:     acid.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete AcidDao db Id:%d,user_id:%d", acid.Id, user.Id))
	return i, err
}

func (d AcidDao) Find(user *models.User, acidPageNow int) ([]models.Acid, int, error, error) {
	var acidList []models.Acid
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (acidPageNow-1)*10).Find(&acidList)
	acid := new(models.Acid)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(acid)
	fmt.Println("total", total)
	return acidList, int(total), err1, err2
}

func (d AcidDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Acid, int, error, error) {
	var acidList []models.Acid
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date >= ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&acidList)
	newAcidList := []models.Acid{}
	fmt.Println("acidList", acidList)
	// 反序
	for i := range acidList {
		newAcidList = append(newAcidList, acidList[len(acidList)-1-i])
	}
	acid := new(models.Acid)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(acid)
	fmt.Println("total", total)
	return newAcidList, int(total), err1, err2
}
func (d AcidDao) SumAll(user *models.User) ([]models.Acid, error) {
	var acidList []models.Acid
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&acidList)
	return acidList, err
}
