package util

import (
	"path/filepath"
	"strings"
)

// 获取文件名，不含路径和后缀
func GetFileName(fp string) string {
	name := filepath.Base(fp)
	ext := filepath.Ext(fp)
	return strings.TrimSuffix(name, ext)
}
