package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type PhoneService struct {
	phoneDao dao.PhoneDao
}

func (s PhoneService) List(user *models.User, phonePageNow int) ([]models.Phone, int, error, error) {
	return s.phoneDao.List(user, phonePageNow)
}

func (s PhoneService) Sum(user *models.User) ([]models.Phone, error) {
	return s.phoneDao.Sum(user)
}

func (s PhoneService) SumAll(user *models.User) ([]models.Phone, error) {
	return s.phoneDao.SumAll(user)
}

func (s PhoneService) BaseRecord(user *models.User, timeStamp string) (*models.Phone, error) {
	return s.phoneDao.BaseRecord(user, timeStamp)
}

func (s PhoneService) UpdatePhoneById(phone models.Phone) (int64, error) {
	return s.phoneDao.UpdatePhoneById(phone)
}
func (s PhoneService) Delete(user *models.User, phone models.Phone) (int64, error) {
	return s.phoneDao.Delete(user, phone)
}

func (s PhoneService) Find(user *models.User, phonePageNow int) ([]models.Phone, int, error, error) {
	return s.phoneDao.Find(user, phonePageNow)
}
func (s PhoneService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Phone, int, error, error) {
	return s.phoneDao.SearchMonth(user, thisMonth, nextMonth)
}
