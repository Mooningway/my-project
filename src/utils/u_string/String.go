package u_string

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	Empty string = ``
)

func String(value interface{}) string {
	if value == nil {
		return ``
	}

	switch result := value.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	case int:
		return strconv.Itoa(result)
	case int8:
		return strconv.Itoa(int(result))
	case int16:
		return strconv.Itoa(int(result))
	case int32:
		return strconv.Itoa(int(result))
	case int64:
		return strconv.FormatInt(result, 10)
	case uint:
		return strconv.Itoa(int(result))
	case uint8:
		return strconv.Itoa(int(result))
	case uint16:
		return strconv.Itoa(int(result))
	case uint32:
		return strconv.Itoa(int(result))
	case uint64:
		return strconv.FormatUint(result, 10)
	case float32:
		return strconv.FormatFloat(float64(result), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(result, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(result)
	default:
		jsonBytes, _ := json.Marshal(value)
		return string(jsonBytes)
	}
}

func Int(val string) (int, error) {
	return strconv.Atoi(val)
}

func Int64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

func Float64(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

func IntToInt64(val int) int64 {
	return int64(val)
}

func IntToString(val int) string {
	return strconv.Itoa(val)
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func IsInt(value interface{}) bool {
	if value == nil {
		return false
	}
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return true
	default:
		return false
	}
}

func Length(value interface{}) int {
	value0 := String(value)
	return len(value0)
}

func Asterisk(val string) string {
	valLen := len(val)
	switch valLen {
	case 0:
		return ""
	case 1:
		return val
	case 2:
		return val[:1] + "*"
	case 3:
		return val[:1] + "**"
	case 4:
		return val[:1] + "**" + val[valLen-1:valLen]
	case 5:
		return val[:1] + "***" + val[valLen-1:valLen]
	case 6:
		return val[:1] + "****" + val[valLen-1:valLen]
	case 7:
		return val[:1] + "*****" + val[valLen-1:valLen]
	case 8:
		return val[:2] + "*****" + val[valLen-1:valLen]
	case 9:
		return val[:2] + "*****" + val[valLen-2:valLen]
	case 10:
		return val[:3] + "*****" + val[valLen-2:valLen]
	case 11:
		return val[:3] + "*****" + val[valLen-3:valLen]
	default:
		var result strings.Builder
		result.WriteString(val[:3])
		for i := 0; i < valLen-6; i++ {
			result.WriteString("*")
		}
		result.WriteString(val[valLen-3 : valLen])
		return result.String()
	}
}

func AsteriskMail(val string) string {
	if len(val) == 0 {
		return ""
	}
	array := strings.Split(val, "@")
	arrayLen := len(array)
	if arrayLen == 0 || arrayLen > 2 {
		return ""
	}
	array1 := strings.Split(array[1], ".")
	return Asterisk(array[0]) + "@" + Asterisk(array1[0]) + "." + array1[1]
}
