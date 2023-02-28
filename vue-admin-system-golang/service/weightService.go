package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type WeightService struct {
	weightDao dao.WeightDao
}

func (s WeightService) List(user *models.User, phonePageNow int) ([]models.Weight, int, error, error) {
	return s.weightDao.List(user, phonePageNow)
}

func (s WeightService) Sum(user *models.User) ([]models.Weight, error) {
	return s.weightDao.Sum(user)
}

func (s WeightService) SumAll(user *models.User) ([]models.Weight, error) {
	return s.weightDao.SumAll(user)
}

func (s WeightService) BaseRecord(user *models.User, timeStamp string) (*models.Weight, error) {
	return s.weightDao.BaseRecord(user, timeStamp)
}

func (s WeightService) UpdateMorningById(weight models.Weight) (int64, error) {
	return s.weightDao.UpdateMorningById(weight)
}

func (s WeightService) UpdateNoonById(weight models.Weight) (int64, error) {
	return s.weightDao.UpdateNoonById(weight)
}
func (s WeightService) UpdateNightById(weight models.Weight) (int64, error) {
	return s.weightDao.UpdateNightById(weight)
}
func (s WeightService) Delete(user *models.User, weight models.Weight) (int64, error) {
	return s.weightDao.Delete(user, weight)
}

func (s WeightService) Find(user *models.User, weightPageNow int) ([]models.Weight, int, error, error) {
	return s.weightDao.Find(user, weightPageNow)
}

func (s WeightService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Weight, int, error, error) {
	return s.weightDao.SearchMonth(user, thisMonth, nextMonth)
}
