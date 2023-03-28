package routers

import (
	"server/admin/schemas/req"
	"server/core"
	"server/core/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

var CommonGroup = core.Group("/common")

func init() {
	group := CommonGroup
	group.AddPOST("/upload/image", uploadImage)
}

// uploadImage 上传图片
func uploadImage(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	util.VerifyUtil.VerifyBody(c, &uReq)
	response.OkWithData(c, "")
}
