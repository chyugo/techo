package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitPhoneRouter(r *gin.RouterGroup) {
	phoneApi := new(v1.PhoneApi)
	group := r.Group("/phone")
	{
		group.Any("/list", phoneApi.List)
		group.Any("/sum", phoneApi.Sum)
		group.Any("/record", phoneApi.Record)
		group.Any("/find", phoneApi.Find)
		group.Any("/searchMonth", phoneApi.SearchMonth)

		group.Any("/delete", phoneApi.Delete)
	}
}
