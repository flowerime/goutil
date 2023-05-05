package ku

import (
	"encoding/binary"
)

// 字节（小端）转为整数
func BytesToInt(b []byte) int {
	var ret int
	for i, v := range b {
		ret |= int(v) << (i << 3)
	}
	return ret
}

type integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

// 转换为长度为 4 的字节切片
func To4Bytes[T integer](i T) []byte {
	u32 := uint32(i)
	ret := make([]byte, 4)
	binary.LittleEndian.PutUint32(ret, u32)
	return ret
}

// 转换为长度为 2 的字节切片
func To2Bytes[T integer](i T) []byte {
	u16 := uint16(i)
	ret := make([]byte, 2)
	binary.LittleEndian.PutUint16(ret, u16)
	return ret
}
