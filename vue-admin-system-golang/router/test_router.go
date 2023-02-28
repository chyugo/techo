package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitTestRouter(r *gin.RouterGroup) {
	testApi := new(v1.TestApi)
	group := r.Group("/test")
	{
		group.Any("/test1", testApi.Test1)
		group.Any("/test2", testApi.Test2)
		group.Any("/test3", testApi.Test3)
		group.Any("/test5", testApi.Test5)
		group.Any("/test6", testApi.Test6)
		group.Any("/test7", testApi.Test7)
		group.Any("/test8", testApi.Test8)

	}
}
