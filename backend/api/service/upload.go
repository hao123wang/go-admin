package service

import (
	"fmt"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"go-admin-server/global"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type UploadService struct{}

func (s *UploadService) UploadFile(filename string) (string, string, error) {
	ext := path.Ext(filename)
	now := time.Now()

	// 新的文件名（根据时间戳命名）
	newFilename := strconv.Itoa(now.Nanosecond()) + ext

	// 存放文件的目录（根据不同的日期，创建不同目录）
	dirPath := filepath.Join("./uploads", now.Format("20060102"))
	if err := utils.CheckAndCreateDir(dirPath); err != nil {
		return "", "", response.ErrServerError
	}
	// 文件的完整存储路径（相对路径）
	fullPath := filepath.Join(dirPath, newFilename)
	// 文件的访问url
	url := fmt.Sprintf("%s:%d%s/%s/%s", "localhost", global.Config.Server.Port,
		"/uploads", now.Format("20060102"), newFilename)
	return fullPath, url, nil
}
