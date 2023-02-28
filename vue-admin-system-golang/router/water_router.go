package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitWaterRouter(r *gin.RouterGroup) {
	waterApi := new(v1.WaterApi)
	group := r.Group("/water")
	{
		group.Any("/list", waterApi.List)
		group.Any("/sum", waterApi.Sum)
		group.Any("/record", waterApi.Record)
		group.Any("/find", waterApi.Find)
		group.Any("/searchMonth", waterApi.SearchMonth)
		group.Any("/delete", waterApi.Delete)
	}
}
