package frombytes

import (
	"encoding/binary"
	"math"
	"strconv"
	"unsafe"
)

const Version = "1.1.0"

type Endian struct {
	little bool
}

var LittleEndian = Endian{true}
var BigEndian = Endian{false}
var ThisEndian = Endian{IsLittleEndian()}

//有符号短整型
func BytesToInt16(b []byte, e Endian) int16 {
	return int16(BytesToUint16(b, e))
}

//无符号短整型
func BytesToUint16(b []byte, e Endian) uint16 {
	if e.little {
		return binary.LittleEndian.Uint16(b)
	}
	return binary.BigEndian.Uint16(b)
}

//有符号长整型
func BytesToInt32(b []byte, e Endian) int32 {
	return int32(BytesToUint32(b, e))
}

//无符号长整型
func BytesToUint32(b []byte, e Endian) uint32 {
	if e.little {
		return binary.LittleEndian.Uint32(b)
	}
	return binary.BigEndian.Uint32(b)
}

//有符号长长整型
func BytesToInt64(b []byte, e Endian) int64 {
	return int64(BytesToUint64(b, e))
}

//无符号长长整型
func BytesToUint64(b []byte, e Endian) uint64 {
	if e.little {
		return binary.LittleEndian.Uint64(b)
	} else {
		return binary.BigEndian.Uint64(b)
	}
}

//有符号整型
func BytesToInt(b []byte, e Endian) int {
	if strconv.IntSize == 32 {
		return int(BytesToInt32(b, e))
	}
	return int(BytesToInt64(b, e))
}

//无符号整型
func BytesToUint(b []byte, e Endian) uint {
	if strconv.IntSize == 32 {
		return uint(BytesToUint32(b, e))
	}
	return uint(BytesToUint64(b, e))
}

//32位浮点型
func BytesToFloat32(b []byte, e Endian) float32 {
	bits := uint32(0)
	if e.little {
		bits = binary.LittleEndian.Uint32(b)
	} else {
		bits = binary.BigEndian.Uint32(b)
	}
	return math.Float32frombits(bits)
}

//64位浮点型
func BytesToFloat64(b []byte, e Endian) float64 {
	bits := uint64(0)
	if e.little {
		bits = binary.LittleEndian.Uint64(b)
	} else {
		bits = binary.BigEndian.Uint64(b)
	}
	return math.Float64frombits(bits)
}

func IsLittleEndian() bool {
	n := 0x1234
	return *(*byte)(unsafe.Pointer(&n)) == 0x34
}

func LocalEndian() binary.ByteOrder {
	if IsLittleEndian() {
		return binary.LittleEndian
	} else {
		return binary.BigEndian
	}
}
