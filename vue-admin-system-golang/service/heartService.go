package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type HeartService struct {
	heartDao dao.HeartDao
}

func (s HeartService) List(user *models.User, phonePageNow int) ([]models.Heart, int, error, error) {
	return s.heartDao.List(user, phonePageNow)
}
func (s HeartService) Sum(user *models.User) ([]models.Heart, error) {
	return s.heartDao.Sum(user)
}

func (s HeartService) SumAll(user *models.User) ([]models.Heart, error) {
	return s.heartDao.SumAll(user)
}

func (s HeartService) BaseRecord(user *models.User, timeStamp string) (*models.Heart, error) {
	return s.heartDao.BaseRecord(user, timeStamp)
}

func (s HeartService) UpdateHeartById(heart models.Heart) (int64, error) {
	return s.heartDao.UpdateHeartById(heart)
}

func (s HeartService) Delete(user *models.User, heart models.Heart) (int64, error) {
	return s.heartDao.Delete(user, heart)
}

func (s HeartService) Find(user *models.User, heartPageNow int) ([]models.Heart, int, error, error) {
	return s.heartDao.Find(user, heartPageNow)
}
func (s HeartService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Heart, int, error, error) {
	return s.heartDao.SearchMonth(user, thisMonth, nextMonth)
}
