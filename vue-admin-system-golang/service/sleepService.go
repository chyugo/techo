package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type SleepService struct {
	sleepDao dao.SleepDao
}

func (s SleepService) List(user *models.User, phonePageNow int) ([]models.Sleep, int, error, error) {
	return s.sleepDao.List(user, phonePageNow)
}
func (s SleepService) SearchMonth(user *models.User, thisMonth int64, nextMonth int64) ([]models.Sleep, int, error, error) {
	return s.sleepDao.SearchMonth(user, thisMonth, nextMonth)
}
func (s SleepService) Delete(user *models.User, sleep models.Sleep) (int64, error) {
	return s.sleepDao.Delete(user, sleep)
}

func (s SleepService) Find(user *models.User, phonePageNow int) ([]models.Sleep, int, error, error) {
	return s.sleepDao.Find(user, phonePageNow)
}

func (s SleepService) Sum(user *models.User) ([]models.Sleep, error) {
	return s.sleepDao.Sum(user)
}

func (s SleepService) SumAll(user *models.User) ([]models.Sleep, error) {
	return s.sleepDao.SumAll(user)
}
func (s SleepService) BaseRecord(user *models.User, timeStamp string) (*models.Sleep, error) {
	return s.sleepDao.BaseRecord(user, timeStamp)
}
func (s SleepService) UpdateGetTimeById(sleep models.Sleep) (int64, error) {
	return s.sleepDao.UpdateGetTimeById(sleep)
}
func (s SleepService) UpdateSleepTimeById(sleep models.Sleep) (int64, error) {
	return s.sleepDao.UpdateSleepTimeById(sleep)
}

func (s SleepService) UpdateSleepActiveById(sleep models.Sleep) (int64, error) {
	return s.sleepDao.UpdateSleepActiveById(sleep)
}
