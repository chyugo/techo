package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitSleepRouter(r *gin.RouterGroup) {
	sleepApi := new(v1.SleepApi)
	group := r.Group("/sleep")
	{
		group.Any("/list", sleepApi.List)
		group.Any("/sum", sleepApi.Sum)
		group.Any("/record", sleepApi.Record)
		group.Any("/find", sleepApi.Find)
		group.Any("/delete", sleepApi.Delete)

		group.Any("/searchMonth", sleepApi.SearchMonth)

	}
}
