package service

import (
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
)

type MusicService struct {
	musicDao dao.MusicDao
}

func (s MusicService) CreateMusic(music models.Music) error {
	return s.musicDao.CreateMusic(music)
}

func (s MusicService) GetMusic(music models.Music) ([]models.Music, error) {
	return s.musicDao.GetMusic(music)

}

func (s MusicService) Find(user *models.User, musicPageNow int) ([]models.Music, int, error, error) {
	return s.musicDao.Find(user, musicPageNow)
}

func (s MusicService) Delete(user *models.User, music models.Music) (int64, error) {
	return s.musicDao.Delete(user, music)
}
