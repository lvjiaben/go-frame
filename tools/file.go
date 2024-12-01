package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func getRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "configs")); err == nil {
			return dir, nil
		}

		// 如果已经到达文件系统的根目录，则返回错误
		if dir == filepath.Dir(dir) {
			return "", fmt.Errorf("go.mod not found")
		}

		dir = filepath.Dir(dir)
	}
}
