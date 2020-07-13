package tobytes

import "errors"

func InterfaceToBates(v interface{}) ([]byte, error) {
	switch v.(type) {
	case bool:
		if v.(bool) {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case int:
		return IntToBytes(v.(int), ThisEndian), nil
	case int8:
		return []byte{uint8(v.(int8))}, nil
	case int16:
		return Int16ToBytes(v.(int16), ThisEndian), nil
	case int32:
		return Int32ToBytes(v.(int32), ThisEndian), nil
	case int64:
		return Int64ToBytes(v.(int64), ThisEndian), nil
	case uint:
		return UintToBytes(v.(uint), ThisEndian), nil
	case uint8:
		return []byte{v.(uint8)}, nil
	case uint16:
		return Uint16ToBytes(v.(uint16), ThisEndian), nil
	case uint32:
		return Uint32ToBytes(v.(uint32), ThisEndian), nil
	case uint64:
		return Uint64ToBytes(v.(uint64), ThisEndian), nil
	case float32:
		return Float32ToBytes(v.(float32), ThisEndian), nil
	case float64:
		return Float64ToBytes(v.(float64), ThisEndian), nil
	case string:
		return []byte(v.(string)), nil
	default:
		return nil, errors.New("type not supported")
	}
}
