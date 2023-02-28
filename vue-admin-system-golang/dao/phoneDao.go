package dao

import (
	"fmt"
	"strconv"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type PhoneDao struct {
}

func (d PhoneDao) List(user *models.User, phonePageNow int) ([]models.Phone, int, error, error) {
	var phoneList []models.Phone
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (phonePageNow-1)*10).Find(&phoneList)
	newPhoneList := []models.Phone{}
	// 反序
	for i := range phoneList {
		newPhoneList = append(newPhoneList, phoneList[len(phoneList)-1-i])
	}
	phone := new(models.Phone)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(phone)
	fmt.Println("total", total)
	return newPhoneList, int(total), err1, err2
}

func (d PhoneDao) Sum(user *models.User) ([]models.Phone, error) {
	var phoneList []models.Phone
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(730, 0).Find(&phoneList)
	newPhoneList := []models.Phone{}
	// 反序
	for i := range phoneList {
		newPhoneList = append(newPhoneList, phoneList[len(phoneList)-1-i])
	}
	return newPhoneList, err
}

func (d PhoneDao) SumAll(user *models.User) ([]models.Phone, error) {
	var phoneList []models.Phone
	// desc 降序
	err := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Find(&phoneList)
	return phoneList, err
}

func (d PhoneDao) BaseRecord(user *models.User, timeStamp string) (*models.Phone, error) {
	phone := models.Phone{}
	count, err := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Count(&phone)
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
		effect, err := SqlDB.NewSession().Insert(&models.Phone{
			UserId:     user.Id,
			RecordDate: time,
		})
		fmt.Println("没有对应记录，执行插入", effect)
		logger.Logger.Info(fmt.Sprintf("insert phone db UserId:%d,RecordDate:%d", user.Id, time))
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}
	var response models.Phone
	SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date=?", timeStamp).Get(&response)
	return &response, nil
}

func (d PhoneDao) UpdatePhoneById(phone models.Phone) (int64, error) {
	fmt.Printf("phonephonephone %#v", phone)
	i, err := SqlDB.NewSession().Where("id=?", phone.Id).Cols("phone_minute").Update(&phone)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update phone db Id:%d,phone_minute:%d", phone.Id, phone.PhoneMinute))

	return i, nil
}

func (d PhoneDao) Delete(user *models.User, phone models.Phone) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", phone.Id).And("user_id=?", user.Id).Delete(models.Phone{
		Id:     phone.Id,
		UserId: user.Id,
	})
	logger.Logger.Info(fmt.Sprintf("delete phone db Id:%d,UserId:%d", phone.Id, user.Id))

	return i, err
}

func (d PhoneDao) Find(user *models.User, phonePageNow int) ([]models.Phone, int, error, error) {
	var phoneList []models.Phone
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).Desc("record_date").Limit(10, (phonePageNow-1)*10).Find(&phoneList)
	phone := new(models.Phone)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(phone)
	fmt.Println("total", total)
	return phoneList, int(total), err1, err2
}

func (d PhoneDao) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Phone, int, error, error) {
	var phoneList []models.Phone
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("record_date > ?", thisMonth).And("record_date< ?", nextMonth).Desc("record_date").Find(&phoneList)
	newPhoneList := []models.Phone{}
	// 反序
	for i := range phoneList {
		newPhoneList = append(newPhoneList, phoneList[len(phoneList)-1-i])
	}
	phone := new(models.Phone)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(phone)
	fmt.Println("total", total)
	return newPhoneList, int(total), err1, err2
}
