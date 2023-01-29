package enc

import (
	"encoding/binary"
	"io"
)

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

// 字节（小端）转为整数
func BytesToInt(b []byte) int {
	var ret int
	for i, v := range b {
		ret |= int(v) << (i << 3)
	}
	return ret
}

// 将 uint32 转换为长度为 4 的字节切片
func To4Bytes(i uint32) []byte {
	ret := make([]byte, 4)
	binary.LittleEndian.PutUint32(ret, i)
	return ret
}

// 将 uint16 转换为长度为 2 的字节切片
func To2Bytes(i uint16) []byte {
	ret := make([]byte, 2)
	binary.LittleEndian.PutUint16(ret, i)
	return ret
}
