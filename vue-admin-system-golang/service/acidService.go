package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type AcidService struct {
	acidDao dao.AcidDao
}

func (s AcidService) BaseRecord(user *models.User, timeStamp string) (*models.Acid, error) {
	return s.acidDao.BaseRecord(user, timeStamp)
}

func (s AcidService) UpdateAcidById(acid models.Acid) (int64, error) {
	return s.acidDao.UpdateAcidById(acid)
}
func (s AcidService) List(user *models.User, acidPageNow int) ([]models.Acid, int, error, error) {
	return s.acidDao.List(user, acidPageNow)
}
func (s AcidService) Sum(user *models.User) ([]models.Acid, error) {
	return s.acidDao.Sum(user)
}
func (s AcidService) Delete(user *models.User, acid models.Acid) (int64, error) {
	return s.acidDao.Delete(user, acid)
}

func (s AcidService) Find(user *models.User, acidPageNow int) ([]models.Acid, int, error, error) {
	return s.acidDao.Find(user, acidPageNow)
}
func (s AcidService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Acid, int, error, error) {
	return s.acidDao.SearchMonth(user, thisMonth, nextMonth)
}
func (s AcidService) SumAll(user *models.User) ([]models.Acid, error) {
	return s.acidDao.SumAll(user)
}
