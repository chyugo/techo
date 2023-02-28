package router

import (
	"github.com/gin-gonic/gin"
	v1 "vue-admin-system-golang/api/v1"
)

func InitLoginRouter(r *gin.RouterGroup) {
	userApi := new(v1.UserApi)
	group := r.Group("/login")
	{
		group.Any("/dologin", userApi.DoLogin)
		group.Any("/doRegist", userApi.DoRegist)

		group.Any("/notOutTime", userApi.NotOutTime)

	}
	r.Any("/export", userApi.Export)
	userGroup := r.Group("/user")
	{
		userGroup.Any("/resetPassword", userApi.ResetPassword)
		userGroup.Any("/sentSuggest", userApi.SentSuggest)
	}

}
