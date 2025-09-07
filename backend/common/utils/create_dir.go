package utils

import "os"

func CheckAndCreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，则创建
		err := os.MkdirAll(dirPath, 0755)
		return err
	}
	return nil
}
