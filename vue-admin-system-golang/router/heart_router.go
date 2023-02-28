package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitHeartRouter(r *gin.RouterGroup) {
	heartApi := new(v1.HeartApi)
	group := r.Group("/heart")
	{
		group.Any("/list", heartApi.List)
		group.Any("/sum", heartApi.Sum)
		group.Any("/record", heartApi.Record)
		group.Any("/find", heartApi.Find)
		group.Any("/delete", heartApi.Delete)
		group.Any("/searchMonth", heartApi.SearchMonth)

	}
}
