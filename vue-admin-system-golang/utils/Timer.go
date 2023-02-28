package utils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/utils/fileFunc"
)

func init() {
	err := fileFunc.CreateDir("./dbBackup")
	if err != nil {
		panic("创建日志文件夹失败")
	}
	cron := cron.New()
	_, err = cron.AddFunc("@every 8h", BackUpDb)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("timer.")
	go cron.Start()
	defer cron.Stop()

}

func BackUpDb() {
	filename := fmt.Sprintf("./dbBackup/%sBackup.sql", time.Now().Format("2006-01-02 150405"))
	dao.SqlDB.DumpAllToFile(filename)
	fmt.Printf("\n定时备份数据库")
}
