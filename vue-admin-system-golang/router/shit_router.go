package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitShitRouter(r *gin.RouterGroup) {
	shitApi := new(v1.ShitApi)
	group := r.Group("/shit")
	{
		group.Any("/list", shitApi.List)
		group.Any("/sum", shitApi.Sum)
		group.Any("/record", shitApi.Record)
		group.Any("/find", shitApi.Find)
		group.Any("/searchMonth", shitApi.SearchMonth)

		group.Any("/delete", shitApi.Delete)

		group.Any("/shitDate", shitApi.ShitDate)

	}
}
