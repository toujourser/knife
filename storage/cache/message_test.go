package cache

import (
	"github.com/robinjoseph08/redisqueue/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageMethods(t *testing.T) {
	// 创建一个测试用的 Message 实例
	message := &Message{
		Message: redisqueue.Message{
			ID:     "testID",
			Stream: "testStream",
			Values: map[string]interface{}{"key": "value"},
		},
	}

	// 测试 GetID 方法
	assert.Equal(t, "testID", message.GetID())

	// 测试 GetStream 方法
	assert.Equal(t, "testStream", message.GetStream())

	// 测试 GetValues 方法
	assert.Equal(t, map[string]interface{}{"key": "value"}, message.GetValues())

	// 测试 SetID 方法
	message.SetID("newID")
	assert.Equal(t, "newID", message.ID)

	// 测试 SetStream 方法
	message.SetStream("newStream")
	assert.Equal(t, "newStream", message.Stream)

	// 测试 SetValues 方法
	message.SetValues(map[string]interface{}{"newKey": "newValue"})
	assert.Equal(t, map[string]interface{}{"newKey": "newValue"}, message.Values)

	// 测试 GetPrefix 方法
	assert.Equal(t, "", message.GetPrefix()) // 因为 Values 中没有 PrefixKey

	// 测试 SetPrefix 方法
	message.SetPrefix("newPrefix")
	assert.Equal(t, "newPrefix", message.GetPrefix())
}
