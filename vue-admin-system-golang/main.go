package main

import (
	"fmt"
	"net/http"
	"vue-admin-system-golang/config"
	_ "vue-admin-system-golang/config"
	_ "vue-admin-system-golang/dao"
	"vue-admin-system-golang/router"
	_ "vue-admin-system-golang/utils"
	"vue-admin-system-golang/utils/fileFunc"
)

func main() {
	router := router.Init()
	router.StaticFS("/more_static", http.Dir("./static/music/"))
	fmt.Printf("RedisCfg%#v", config.RedisCfg)
	fmt.Printf("AppCfg%#v", config.AppCfg)
	fmt.Printf("MySqlCfg%#v", config.MySqlCfg)
	err := fileFunc.CreateDir("./Excels")
	if err != nil {
		panic("创建导出excel文件夹失败")
	}
	err = fileFunc.CreateDir("./monthReport")
	if err != nil {
		panic("创建导出月报文件夹失败")
	}
	err = router.Run(":8092")
	if err != nil {
		fmt.Println("err", err)
	}

}
