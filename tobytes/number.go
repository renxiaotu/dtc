package tobytes

import (
	"encoding/binary"
	"math"
	"strconv"
	"unsafe"
)

type Endian bool

var LittleEndian Endian = true
var BigEndian Endian = false
var ThisEndian = thisEndian()

//有符号短整型
func Int16ToBytes(n int16, e Endian) []byte {
	return Uint16ToBytes(uint16(n), e)
}

//无符号短整型
func Uint16ToBytes(n uint16, e Endian) []byte {
	if e {
		return []byte{
			byte(n),
			byte(n >> 8),
		}
	}
	return []byte{
		byte(n >> 8),
		byte(n),
	}
}

//有符号长整型
func Int32ToBytes(n int32, e Endian) []byte {
	return Uint32ToBytes(uint32(n), e)
}

//无符号长整型
func Uint32ToBytes(n uint32, e Endian) []byte {
	if e {
		return []byte{byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24)}
	}
	return []byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
}

//有符号长长整型
func Int64ToBytes(n int64, e Endian) []byte {
	return Uint64ToBytes(uint64(n), e)
}

//无符号长长整型
func Uint64ToBytes(n uint64, e Endian) []byte {
	if e {
		return []byte{byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24), byte(n >> 32), byte(n >> 40), byte(n >> 48), byte(n >> 56)}
	} else {
		return []byte{byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32), byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
	}
}

//有符号整型
func IntToBytes(n int, e Endian) []byte {
	if strconv.IntSize == 32 {
		return Int32ToBytes(int32(n), e)
	}
	return Int64ToBytes(int64(n), e)
}

//无符号整型
func UintToBytes(n uint, e Endian) []byte {
	if strconv.IntSize == 32 {
		return Uint32ToBytes(uint32(n), e)
	}
	return Uint64ToBytes(uint64(n), e)
}

//32位浮点型
func Float32ToBytes(f float32, e Endian) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	if e {
		binary.LittleEndian.PutUint32(bytes, bits)
	} else {
		binary.BigEndian.PutUint32(bytes, bits)
	}
	return bytes
}

//64位浮点型
func Float64ToBytes(f float64, e Endian) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	if e {
		binary.LittleEndian.PutUint64(bytes, bits)
	} else {
		binary.BigEndian.PutUint64(bytes, bits)
	}
	return bytes
}

func thisEndian() Endian {
	if IsLittleEndian() {
		return true
	}
	return false
}

func IsLittleEndian() bool {
	n := 0x1234
	return *(*byte)(unsafe.Pointer(&n)) == 0x34
}
