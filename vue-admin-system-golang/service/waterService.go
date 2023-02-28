package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type WaterService struct {
	waterDao dao.WaterDao
}

func (s WaterService) BaseRecord(user *models.User, timeStamp string) (*models.Water, error) {
	return s.waterDao.BaseRecord(user, timeStamp)
}

func (s WaterService) UpdateWaterById(water models.Water) (int64, error) {
	return s.waterDao.UpdateWaterById(water)
}
func (s WaterService) List(user *models.User, waterPageNow int) ([]models.Water, int, error, error) {
	return s.waterDao.List(user, waterPageNow)
}
func (s WaterService) Sum(user *models.User) ([]models.Water, error) {
	return s.waterDao.Sum(user)
}
func (s WaterService) Delete(user *models.User, water models.Water) (int64, error) {
	return s.waterDao.Delete(user, water)
}

func (s WaterService) Find(user *models.User, waterPageNow int) ([]models.Water, int, error, error) {
	return s.waterDao.Find(user, waterPageNow)
}
func (s WaterService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Water, int, error, error) {
	return s.waterDao.SearchMonth(user, thisMonth, nextMonth)
}
func (s WaterService) SumAll(user *models.User) ([]models.Water, error) {
	return s.waterDao.SumAll(user)
}
