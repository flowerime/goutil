package ku

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/gogs/chardet"
	"golang.org/x/net/html/charset"
)

// 将字节流转换为 utf-8
func NewReader(input io.Reader) io.Reader {
	brd := bufio.NewReader(input)
	buf, _ := brd.Peek(1024)
	detector := chardet.NewTextDetector()
	cs, err := detector.DetectBest(buf) // 检测编码格式
	if err != nil {
		return brd
	}
	if cs.Confidence != 100 && cs.Charset != "UTF-8" {
		cs.Charset = "GB18030"
	}
	// 删除 BOM 文件头
	boms := make(map[string][]byte)
	boms["UTF-16BE"] = []byte{0xfe, 0xff}
	boms["UTF-16LE"] = []byte{0xff, 0xfe}
	boms["UTF-8"] = []byte{0xef, 0xbb, 0xbf}
	if b, ok := boms[cs.Charset]; ok {
		if bytes.HasPrefix(buf, b) {
			brd.Read(b)
		}
	}
	rd, _ := charset.NewReaderLabel(cs.Charset, brd) // 转换字节流
	return rd
}

// 读取文件
func Read(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return f, err
	}
	return NewReader(f), nil
}

// 读 2 个字节，按小端序转为 uint16
func ReadUint16(r io.Reader) uint16 {
	tmp := make([]byte, 2)
	r.Read(tmp)
	return binary.LittleEndian.Uint16(tmp)
}

// 读 4 个字节，按小端序转为 uint32
func ReadUint32(r io.Reader) uint32 {
	tmp := make([]byte, 4)
	r.Read(tmp)
	return binary.LittleEndian.Uint32(tmp)
}
