package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// 用户密码进行md5+salt加密
func CreatePassWordSecret(s string) string {
	fmt.Println("传入的password", s)
	s = fmt.Sprintf("%s&&&secret=%s", s, "fmoalmfklsd")
	code := md5.Sum([]byte(s))
	c := fmt.Sprintf("%x", code)
	fmt.Println("c", c)
	return string(c)
}

// 用户密码进行md5+salt加密
func CreateCookie(s string) string {
	fmt.Println("传入", s)
	s = fmt.Sprintf("%s&&&secret=%s&&&&%s", s, "fsagdasf", time.Now().Unix())
	code := md5.Sum([]byte(s))
	c := fmt.Sprintf("%x", code)
	return string(c)
}

// excel下载链接进行md5+salt加密
func CreateFilename(s string) string {
	fmt.Println("传入", s)
	rand.Seed(time.Now().Unix())
	s = fmt.Sprintf("%s&&&secret=%s&&&&%s%d", s, "ffasbjksasf", time.Now().Unix(), rand.Intn(100000000))
	code := md5.Sum([]byte(s))
	c := fmt.Sprintf("%x%d", code, time.Now().Unix())
	return string(c)
}
