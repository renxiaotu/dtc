package tobytes

func BoolToBytes(bs ...bool) []byte {
	by := make([]byte, 0)
	n := uint8(0)
	place := uint8(0)
	for i := 0; i < len(bs); i++ {
		n = n << 1
		n += BoolToUint8(bs[i])
		place++
		if place == 8 {
			place = 0
			by = append(by, n)
			n = 0
		}
	}
	if place > 0 {
		by = append(by, n)
	}
	return by
}

func BoolToByte(b bool) byte {
	return BoolToUint8(b)
}

func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
