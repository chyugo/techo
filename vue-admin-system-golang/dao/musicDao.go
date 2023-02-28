package dao

import (
	"fmt"
	"vue-admin-system-golang/models"
	"vue-admin-system-golang/utils/logger"
)

type MusicDao struct {
}

func (d MusicDao) CreateMusic(music models.Music) error {
	fmt.Println("music", music)
	_, err := SqlDB.NewSession().Insert(music)
	logger.Logger.Info(fmt.Sprintf("insert music db UserId:%d,oldname:%s,name:%s", music.UserId, music.OldName, music.Name))
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (d MusicDao) GetMusic(music models.Music) ([]models.Music, error) {
	var musicList []models.Music
	err := SqlDB.NewSession().Where("user_id=?", music.UserId).And("del_flag=0").Find(&musicList)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return musicList, nil
}

func (d MusicDao) Find(user *models.User, musicPageNow int) ([]models.Music, int, error, error) {
	var musicList []models.Music
	// desc 降序
	err1 := SqlDB.NewSession().Where("user_id=?", user.Id).And("del_flag=0").Desc("id").Limit(10, (musicPageNow-1)*10).Find(&musicList)
	music := new(models.Music)
	total, err2 := SqlDB.NewSession().Where("user_id=?", user.Id).Count(music)
	fmt.Println("total", total)
	return musicList, int(total), err1, err2
}

func (d MusicDao) Delete(user *models.User, music models.Music) (int64, error) {
	// desc 降序
	i, err := SqlDB.NewSession().Where("id=?", music.Id).And("user_id=?", user.Id).Update(models.Music{
		DelFlag: 1,
	})
	logger.Logger.Info(fmt.Sprintf("delete shit db id:%d,UserId:%d", music.Id, user.Id))
	return i, err
}
