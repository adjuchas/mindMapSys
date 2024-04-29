package tools

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func DraftToDot(sourcePath string) (string, error) {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("open source err")
	}
	defer sourceFile.Close()
	fileName := filepath.Base(sourcePath)
	tar := filepath.Join("uploads/dot", fileName)
	destFile, _ := os.Create(tar)
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return "", fmt.Errorf("拷贝文件失败: %s", err)
	}
	return tar, nil
}
