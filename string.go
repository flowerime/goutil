package ku

import (
	"path/filepath"
	"runtime"
	"strings"
	uc "unicode"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

// 判断字符串中是否全是汉字
func IsHan(word string) bool {
	flag := true
	for _, r := range word {
		if !uc.Is(uc.Han, r) {
			flag = false
		}
	}
	return flag
}

// 去除字符串中的所有空白字符
func TrimSpace(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if !uc.IsSpace(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// []byte, encoding -> string
func Decode(b []byte, e string) (string, error) {
	enc := getEncoding(e)
	b, err := enc.NewDecoder().Bytes(b)
	return string(b), err
}

// string, encoding -> []byte
func Encode(s, e string) ([]byte, error) {
	enc := getEncoding(e)
	return enc.NewEncoder().Bytes([]byte(s))
}

func getEncoding(enc string) encoding.Encoding {
	var encoding encoding.Encoding
	switch enc {
	case "UTF-16LE":
		encoding = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "UTF-16BE":
		encoding = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	case "GBK":
		encoding = simplifiedchinese.GBK
	case "GB18030":
		encoding = simplifiedchinese.GB18030
	default:
		encoding = unicode.UTF8
	}
	// fmt.Println(enc, encoding)
	return encoding
}

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
