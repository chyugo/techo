package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var MySqlCfg MySqlConfig
var AppCfg AppConfig
var RedisCfg RedisConfig

type MySqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type AppConfig struct {
	Server string
}
type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

func LoadMySqlCfg(cfg *ini.File) {
	MySqlCfg.Username = cfg.Section("mysql").Key("username").String()
	MySqlCfg.Password = cfg.Section("mysql").Key("password").String()
	MySqlCfg.Host = cfg.Section("mysql").Key("host").String()
	MySqlCfg.Port = cfg.Section("mysql").Key("port").String()
	MySqlCfg.Database = cfg.Section("mysql").Key("database").String()
}

func LoadAppCfg(cfg *ini.File) {
	AppCfg.Server = cfg.Section("app").Key("server").String()
}

func LoadRedisCfg(cfg *ini.File) {
	RedisCfg.Addr = cfg.Section("redis").Key("addr").String()
	RedisCfg.Password = cfg.Section("redis").Key("password").String()
	RedisCfg.Db, _ = cfg.Section("redis").Key("db").Int()
}

func init() {
	cfg, err := ini.Load("./config/development.ini")
	//cfg, err := ini.Load("./config/product.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
		panic("无法获取配置文件")
	}
	LoadMySqlCfg(cfg)
	LoadAppCfg(cfg)
	LoadRedisCfg(cfg)
}
