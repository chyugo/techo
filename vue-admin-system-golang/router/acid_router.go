package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitAcidRouter(r *gin.RouterGroup) {
	acidApi := new(v1.AcidApi)
	group := r.Group("/acid")
	{
		group.Any("/list", acidApi.List)
		group.Any("/sum", acidApi.Sum)
		group.Any("/record", acidApi.Record)
		group.Any("/find", acidApi.Find)
		group.Any("/searchMonth", acidApi.SearchMonth)
		group.Any("/delete", acidApi.Delete)
		group.Any("/acidDate", acidApi.AcidDate)

	}
}
