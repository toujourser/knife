package uuid

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
	mu            sync.Mutex
)

func init() {
	machineID = 101 << 12
	lastTimeStamp = time.Now().UnixNano() / 1000
}

// 生成雪花ID
func GetSnowflakeIdProcess() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000
	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 4095 {
			time.Sleep(time.Microsecond)
			curTimeStamp = time.Now().UnixNano() / 1000
			sn = 0
		}
	} else {
		sn = 0
	}
	sn++

	lastTimeStamp = curTimeStamp
	// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
	// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
	rightBinValue := curTimeStamp & 0x1FFFFFF
	// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
	rightBinValue <<= 15
	id := rightBinValue | machineID | sn
	return id
}

func GetSnowflakeId() int64 {
	mu.Lock()
	defer mu.Unlock()
	return GetSnowflakeIdProcess()
}

// 生成随机字符串(带有位数)
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机字符串(特殊字符)
func GetRandomSpecialString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.,*@#&+~!"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机数生成,带有位数限制
func CreateValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var data strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&data, "%d", numeric[rand.Intn(r)])
	}
	return data.String()
}

// 生成一个在区间范围的随机数
func GenerateRangeNum(min, max int) int {
	// rand.Seed(time.Now().Unix())//时间戳秒，跟新的随机数相同，暂时不用
	rand.Seed(time.Now().UnixNano()) //获取纳秒
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}
