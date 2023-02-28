package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type WakeService struct {
	wakeDao dao.WakeDao
}

func (s WakeService) List(user *models.User, wakePageNow int) ([]models.Wake, int, error, error) {
	return s.wakeDao.List(user, wakePageNow)
}

func (s WakeService) Sum(user *models.User) ([]models.Wake, error) {
	return s.wakeDao.Sum(user)
}

func (s WakeService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Wake, int, error, error) {
	return s.wakeDao.SearchMonth(user, thisMonth, nextMonth)
}
