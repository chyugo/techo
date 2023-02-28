package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "vue-admin-system-golang/api/v1"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/utils"
	"vue-admin-system-golang/utils/limiter"
	"vue-admin-system-golang/utils/logger"
)

var limitChan = make(chan struct{}, 2000)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	router.Use(utils.Auth())
	router.Use(logger.LoggerToFile())
	router.Use(IPPass())
	router.Use(ServerPass())
	r := router.Group("/")
	{
		InitLoginRouter(r)
		InitPhoneRouter(r)
		InitHeartRouter(r)
		InitShitRouter(r)
		InitSleepRouter(r)
		InitWeightRouter(r)
		InitWakeRouter(r)
		//InitTestRouter(r)
		InitMusicRouter(r)
		InitWaterRouter(r)
		InitAcidRouter(r)
	}
	r.Any("doc/download", v1.OtherApi{}.DownloadDoc)
	return router
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func IPPass() gin.HandlerFunc {
	// 针对同一个Ip 每秒限制200次
	return func(c *gin.Context) {
		fmt.Println("c.ClientIP()", c.ClientIP())
		l := limiter.NewLimiter(200, 200, c.ClientIP())
		if l.Allow() {
			fmt.Println("pass")
		} else {
			fmt.Println("reject")
			resp.Error(c, -4, "访问频率过快")
			c.Abort()
			return
		}
	}

}

func ServerPass() gin.HandlerFunc {
	// 针对服务器并发 每秒限制2000次
	return func(c *gin.Context) {
		limitChan <- struct{}{}
		c.Next()
		<-limitChan
	}
}
