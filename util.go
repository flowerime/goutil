package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

// 获取文件名，不含路径和后缀
func GetFileName(fp string) string {
	name := filepath.Base(fp)
	ext := filepath.Ext(fp)
	return strings.TrimSuffix(name, ext)
}

// 不同操作系统的换行符
var LineBreak string

func init() {
	switch runtime.GOOS {
	case "windows":
		LineBreak = "\r\n"
	case "darwin":
		LineBreak = "\r"
	default:
		LineBreak = "\n"
	}
}
