package conv

import "strconv"

func ParseStringSliceToUint64(s []string) []uint64 {
	iv := make([]uint64, len(s))
	for i, v := range s {
		iv[i], _ = strconv.ParseUint(v, 10, 64)
	}
	return iv
}

//string转换int
func StringInt(data string) int {
	var stringInt int
	if data != "" {
		stringInt, _ = strconv.Atoi(data)
	} else {
		stringInt = 0
	}
	return stringInt
}

//int转换string
func IntString(data int) string {
	var intString string
	intString = strconv.Itoa(data)
	return intString
}

//int64转换int
func Int64Int(data int64) int {
	var int64Int int
	int64Int, err := strconv.Atoi(strconv.FormatInt(data, 10))
	if err != nil {
		int64Int = 0
	}
	return int64Int
}

//int转换int64
func IntInt64(data int) int64 {
	var intInt64 int64
	intInt64, err := strconv.ParseInt(strconv.Itoa(data), 10, 64)
	if err != nil {
		intInt64 = 0
	}
	return intInt64
}

//int64转换string
func Int64String(data int64) string {
	var intString string
	intString = strconv.FormatInt(data, 10)
	return intString
}

//string转换int64
func StringInt64(data string) int64 {
	var stringInt64 int64
	if data != "" {
		stringInt64, _ = strconv.ParseInt(data, 10, 64)
	} else {
		stringInt64 = 0
	}
	return stringInt64
}

//float64转换string
func Float64String(data float64) string {
	var float64Strig string
	float64Strig = strconv.FormatFloat(data, 'f', -1, 64)
	return float64Strig
}

//string转换float64
func StringFloat64(data string) float64 {
	var stringFloat64 float64
	if data != "" {
		stringFloat64, _ = strconv.ParseFloat(data, 64)
	} else {
		stringFloat64 = 0
	}
	return stringFloat64
}
