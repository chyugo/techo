package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type UserService struct {
	userDao dao.UserDao
}

func (d UserService) Login(user *models.User) (*models.User, error) {
	return d.userDao.Login(user)
}

func (d UserService) Regist(user *models.User) (int64, error) {
	return d.userDao.Regist(user)
}

func (d UserService) UpdatePassword(user *models.User) (int64, error) {
	return d.userDao.UpdatePassword(user)
}
func (d UserService) SaveSuggest(suggest models.Suggest) (int64, error) {
	return d.userDao.SaveSuggest(suggest)
}
