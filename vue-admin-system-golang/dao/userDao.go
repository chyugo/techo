package dao

import (
	"errors"
	"fmt"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type UserDao struct {
}

func (d UserDao) Login(user *models.User) (*models.User, error) {
	findUser := models.User{}
	session := SqlDB.NewSession().Table(models.User{}.TableName())
	session.Where("user_name=?", user.UserName).And("del_flag=?", 0)
	get_flag, err := session.Get(&findUser)
	if !get_flag {
		return nil, errors.New("找不到用户")
	}
	return &findUser, err
}

func (d UserDao) Regist(user *models.User) (int64, error) {
	findUser := models.User{
		UserName: user.UserName,
	}
	session := SqlDB.NewSession().Table(models.User{}.TableName())
	// 判断用户名是否重复
	count, err := session.Where("user_name=?", user.UserName).Count(findUser)
	if err != nil {
		fmt.Println("err")
		return -1, err
	}
	if count != 0 {
		return count, err
	}
	_, err = SqlDB.NewSession().Table(models.User{}.TableName()).Insert(user)
	logger.Logger.Info(fmt.Sprintf("insert user .user_name:%s ,pass_word:%s", user.UserName, user.PassWord))
	return 0, err
}

func (d UserDao) UpdatePassword(user *models.User) (int64, error) {
	session := SqlDB.NewSession().Table(models.User{}.TableName())
	i, err := session.Where("id=?", user.Id).Cols("pass_word").Update(user)
	if err != nil {
		logger.Logger.Info(fmt.Sprintf("update password err. Id:%d,pass_word:%s", user.Id, user.PassWord))
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("update password success. Id:%d,pass_word:%s", user.Id, user.PassWord))
	return i, nil
}

func (d UserDao) SaveSuggest(suggest models.Suggest) (int64, error) {
	session := SqlDB.NewSession().Table(models.Suggest{}.TableName())
	i, err := session.Insert(suggest)
	if err != nil {
		logger.Logger.Info(fmt.Sprintf("insert suggest err. content:%s", suggest.Content))
		fmt.Println("err", err)
		return -1, err
	}
	logger.Logger.Info(fmt.Sprintf("insert suggest success. content:%s", suggest.Content))
	return i, nil
}
