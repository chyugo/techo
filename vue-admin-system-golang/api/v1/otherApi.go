package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	resp "vue-admin-system-golang/models/response"
	"vue-admin-system-golang/utils"
)

type OtherApi struct {
}

func (a OtherApi) DownloadDoc(c *gin.Context) {
	month := c.PostForm("title")
	filePath, err := utils.DocGenerate(month)
	if err != nil {
		resp.Error(c)
		return
	}
	fmt.Println("file", filePath)
	c.File(filePath)
}
