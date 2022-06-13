package file

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type File struct {
	Size int64  `json:"size"`
	Path string `json:"path"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func Upload(c *gin.Context, path string) (*File, error) {
	files, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	// 上传文件至指定目录
	guid := uuid.New().String()
	fileName := guid + GetExt(files.Filename)
	err = IsNotExistMkDir(path)
	if err != nil {
		return nil, err
	}
	singleFile := path + fileName
	err = c.SaveUploadedFile(files, singleFile)
	if err != nil {
		return nil, err
	}
	fileType, _ := GetType(singleFile)
	return &File{
		Size: GetFileSize(singleFile),
		Path: singleFile,
		Name: files.Filename,
		Type: fileType,
	}, nil
}
