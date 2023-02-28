package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitWakeRouter(r *gin.RouterGroup) {
	wakeApi := new(v1.WakeApi)
	group := r.Group("/wake")
	{
		group.Any("/list", wakeApi.List)
		group.Any("/sum", wakeApi.Sum)
		group.Any("/searchMonth", wakeApi.SearchMonth)

	}
}
