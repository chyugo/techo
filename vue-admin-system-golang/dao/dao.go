package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"vue-admin-system-golang/config"
)

var (
	SqlDB   *xorm.Engine
	RedisDB *redis.Client
)

func init() {
	//dbc := "root" + ":" + "123456" + "@tcp(" + "127.0.0.1" + ":" + "3306" + ")/" + "tetyou" + "?charset=utf8"
	dbc := config.MySqlCfg.Username + ":" + config.MySqlCfg.Password + "@tcp(" + config.MySqlCfg.Host + ":" + config.MySqlCfg.Port + ")/" + config.MySqlCfg.Database + "?charset=utf8"
	SqlDB, _ = xorm.NewEngine("mysql", dbc)
	fmt.Println("dbc", dbc)
	err := SqlDB.Ping()
	if err != nil {
		fmt.Println("err", err)
	}
	SqlDB.ShowSQL()

	fmt.Println("sqldb", SqlDB)
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetMaxOpenConns(100)

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisCfg.Addr,
		Password: config.RedisCfg.Password,
		DB:       config.RedisCfg.Db,
	})
	ctx := context.Background()
	cn := rdb.Conn(ctx)
	defer cn.Close()

	// 验证是否连接到redis服务端
	res, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("connect failed.err:", err)
		panic("redis err")
	}
	fmt.Printf("connect successful!\n Ping=>%v\n", res)
	RedisDB = rdb
}

//
//func GetRedis() *redis.Client {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "",
//		DB:       0,
//	})
//	ctx := context.Background()
//	cn := rdb.Conn(ctx)
//	defer cn.Close()
//
//	// 验证是否连接到redis服务端
//	res, err := rdb.Ping(ctx).Result()
//	if err != nil {
//		fmt.Println("connect failed.err:", err)
//		return nil
//	}
//	fmt.Printf("connect successful!\n Ping=>%v\n", res)
//	return rdb
//}
