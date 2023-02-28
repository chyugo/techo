package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
	"vue-admin-system-golang/dao"
	"vue-admin-system-golang/models"
	resp "vue-admin-system-golang/models/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if doSqure(c) {
			fmt.Println("在放行请求内")
			return
		}
		cookie := c.GetHeader("Authorization")
		_, err := dao.RedisDB.Get(context.Background(), cookie).Result()
		fmt.Println(c.Request.RequestURI, "权限验证中间件", cookie)
		if err != nil {
			fmt.Println("err", err)
			fmt.Println("无权")
			resp.Error(c, -2, "无权")
			c.Abort()
			return
		}
	}
}
func doSqure(c *gin.Context) bool {
	var request []string
	request = append(request, "/login")
	request = append(request, "/home")
	request = append(request, "/test")

	request = append(request, "/more_static")

	for i := 0; i < len(request); i++ {
		replace := strings.Contains(c.Request.RequestURI, request[i])
		if replace {
			return true
		}
	}
	return false
}

func GetUserByRedis(c *gin.Context) (*models.User, error) {
	session, err := dao.RedisDB.Get(context.Background(), c.Request.Header.Get("Authorization")).Result()
	fmt.Println("c.Request.Header.Get(\"Authorization\")", c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	var user models.User
	err = json.Unmarshal([]byte(session), &user)
	fmt.Println("当前用户是", user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 上传文件文件名进行base64
func CreateFileSecret() string {
	rand.Seed(time.Now().Unix())
	fileName := fmt.Sprintf("%d", rand.Intn(10000000))
	fmt.Println("文件名", fileName)
	return fmt.Sprintf("%s&%d", fileName, time.Now().Unix())
}
