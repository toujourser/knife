package conv

import "strconv"

// ParseStringSliceToUint64 将字符串切片转换为 uint64 切片
func ParseStringSliceToUint64(s []string) []uint64 {
	iv := make([]uint64, len(s))
	for i, v := range s {
		iv[i], _ = strconv.ParseUint(v, 10, 64)
	}
	return iv
}

// StringInt 将字符串转换为整数
func StringInt(data string) int {
	stringInt, _ := strconv.Atoi(data)
	return stringInt
}

// IntString 将整数转换为字符串
func IntString(data int) string {
	return strconv.Itoa(data)
}

// Int64Int 将 int64 转换为 int
func Int64Int(data int64) int {
	int64Int, _ := strconv.Atoi(strconv.FormatInt(data, 10))
	return int64Int
}

// IntInt64 将 int 转换为 int64
func IntInt64(data int) int64 {
	intInt64, _ := strconv.ParseInt(strconv.Itoa(data), 10, 64)
	return intInt64
}

// Int64String 将 int64 转换为字符串
func Int64String(data int64) string {
	return strconv.FormatInt(data, 10)
}

// StringInt64 将字符串转换为 int64
func StringInt64(data string) int64 {
	stringInt64, _ := strconv.ParseInt(data, 10, 64)
	return stringInt64
}

// Float64String 将 float64 转换为字符串
func Float64String(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 64)
}

// StringFloat64 将字符串转换为 float64
func StringFloat64(data string) float64 {
	stringFloat64, _ := strconv.ParseFloat(data, 64)
	return stringFloat64
}
