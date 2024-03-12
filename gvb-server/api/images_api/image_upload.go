package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
)

// 有的图片上传成功有的失败 需要给用户返回响应
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// 上传图片 返回图片的url
func (ImagesApi) ImageUploadView(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", ctx)
		return
	}

	// 判断路径是否存在 不存在则创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []FileUploadResponse
	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)

		// 判断文件的大小 太大了就失败
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2f MB，设定大小为：%d MB", size, global.Config.Upload.Size),
			})
			continue
		}

		err := ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: true,
				Msg:       err.Error(),
			})
			continue
		}

		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}

	res.OkWithData(resList, ctx)

}
