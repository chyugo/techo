package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitWeightRouter(r *gin.RouterGroup) {
	weightApi := new(v1.WeightApi)
	group := r.Group("/weight")
	{
		group.Any("/list", weightApi.List)
		group.Any("/sum", weightApi.Sum)
		group.Any("/record", weightApi.Record)
		group.Any("/find", weightApi.Find)
		group.Any("/searchMonth", weightApi.SearchMonth)

		group.Any("/delete", weightApi.Delete)

	}
}
