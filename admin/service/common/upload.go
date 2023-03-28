package common

import (
	"mime/multipart"
)

var UploadService = uploadService{}

// uploadService 上传服务实现类
type uploadService struct{}

// UploadImage 上传图片
func (upSrv uploadService) UploadImage(file *multipart.FileHeader, cid uint, aid uint) bool {
	return true
}
