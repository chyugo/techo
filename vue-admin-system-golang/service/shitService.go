package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type ShitService struct {
	shitDao dao.ShitDao
}

func (s ShitService) List(user *models.User, phonePageNow int) ([]models.Shit, int, error, error) {
	return s.shitDao.List(user, phonePageNow)
}
func (s ShitService) Sum(user *models.User) ([]models.Shit, error) {
	return s.shitDao.Sum(user)
}

func (s ShitService) SumAll(user *models.User) ([]models.Shit, error) {
	return s.shitDao.SumAll(user)
}

func (s ShitService) BaseRecord(user *models.User, timeStamp string) (*models.Shit, error) {
	return s.shitDao.BaseRecord(user, timeStamp)
}

func (s ShitService) UpdateShitById(shit models.Shit) (int64, error) {
	return s.shitDao.UpdateShitById(shit)
}
func (s ShitService) Delete(user *models.User, shit models.Shit) (int64, error) {
	return s.shitDao.Delete(user, shit)
}

func (s ShitService) Find(user *models.User, shitPageNow int) ([]models.Shit, int, error, error) {
	return s.shitDao.Find(user, shitPageNow)
}
func (s ShitService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Shit, int, error, error) {
	return s.shitDao.SearchMonth(user, thisMonth, nextMonth)
}
