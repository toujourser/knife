# 项目总览（中文）

本项目是一个基于 Go 1.21 的模块化工具库（module: github.com/toujourser/utils）。我们以“按领域划分包（package-per-domain）”为核心设计思想，对常用能力进行轻量封装，帮助你在业务开发中快速落地稳定的基础功能。

- 语言与版本：Go 1.21
- 许可协议：GPL，详见根目录 LICENSE
- 安装：
  - go get -v -u github.com/toujourser/utils@latest

## 设计理念与架构

- 领域解耦：顶层目录即领域包（algorithm、concurrency、conv、slice/stream、storage、logger、mail、request、validation、uuid、ws 等），各自独立、低耦合。
- 轻封装：对社区通用库进行直观、惯用的二次封装（如 go-redis、NSQ、zap/logrus、go-playground/validator、gin、Google generative-ai 等），隐藏繁琐配置。
- 组合优先：跨包复用通过简洁接口组合实现（如 Feishu Webhook 组合了签名算法与 HTTP 请求）。
- 泛型友好：大量集合操作（slice、stream）使用 Go 泛型，简化业务编码。
- 易于迁移：模块间边界清晰，选择性引入，不会把无关依赖“拖进来”。

目录结构示例（节选）：
- algorithm：签名、LRU、排序、查找
- concurrency：通道（channel）与并发模式
- conv：常见类型转换
- datetime / timeparse：日期时间格式化与解析
- encrypt：AES/DES/RSA/HMAC/Hash/Base64 等
- file：文件读写、上传、下载、监控、gzip
- gemini：Google Generative AI 示例
- im_hook：飞书自定义机器人 WebHook
- logger / logrus / beego-log：多日志实现的便捷封装
- mail：邮件发送与模板
- random / mathutil：随机与数学工具
- request / retry：HTTP 请求与重试
- slice / stream：集合操作与流式处理
- storage（cache/queue）：内存/Redis 缓存，Redis/NSQ 队列
- structure：结构体复制与映射
- uuid：UUID/Snowflake/随机串
- validation：基于 go-playground 的校验与多语言翻译
- ws：WebSocket 管理与推送

## 快速开始

安装：

```bash
go get -v -u github.com/toujourser/utils@latest
```

在项目中像使用标准库一样引入，例如：

```go
import (
    "github.com/toujourser/utils/logger"
    "github.com/toujourser/utils/request"
)
```

## 常用模块与示例

以下示例仅展示最小用法，更多能力请查阅对应包源码与注释。

### 1. 日志 logger（zap）

```go
import "github.com/toujourser/utils/logger"

func init() {
    logger.Init(&logger.LogConf{
        Filename:   "/var/log/app.log",
        MaxSize:    128,
        MaxBackups: 7,
        MaxAge:     7,
        Compress:   true,
        LocalTime:  true,
        Level:      "info",
        JsonFormat: false,
        StdOut:     true,
        NewLog:     false,
    })
}

func demo() {
    logger.Infof("hello %s", "world")
    logger.Error("something wrong")
}
```

若偏好 logrus：

```go
import ulog "github.com/toujourser/utils/logrus"

func init() {
    ulog.Init("/var/log", "app.log", 7, 24) // maxAge(days), rotationTime(hours)
}

func demo() {
    ulog.WithField("trace_id", "abc").Info("hello")
}
```

### 2. HTTP 请求

```go
import "github.com/toujourser/utils/request"

resp, err := request.Request("GET", "https://httpbin.org/get", map[string]string{
    "User-Agent": "utils-demo",
}, nil)
// 使用 resp.Body 自行处理
```

### 3. 重试 retry

```go
import (
    "errors"
    "github.com/toujourser/utils/retry"
)

err := retry.Retry(func() error {
    // do something that may fail
    return errors.New("try again")
}, retry.RetryTimes(3), retry.RetryDuration(time.Second))
```

### 4. 缓存 storage/cache（内存与 Redis）

- 内存

```go
import "github.com/toujourser/utils/storage/cache"

c := cache.NewMemory()
_ = c.Set("k", "v", 60)
val, _ := c.Get("k") // "v"
```

- Redis（使用 go-redis）

```go
import (
    "github.com/go-redis/redis"
    "github.com/toujourser/utils/storage/cache"
)

cli := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
rc, _ := cache.NewRedis(cli, nil)
_ = rc.Set("k", "v", 60)
val, _ := rc.Get("k")
```

### 5. 队列 storage/queue（Redis Stream 与 NSQ）

- Redis Stream（robinjoseph08/redisqueue）

```go
import (
    "github.com/robinjoseph08/redisqueue/v2"
    "github.com/toujourser/utils/storage"
    q "github.com/toujourser/utils/storage/queue"
)

producerOpt := &redisqueue.ProducerOptions{ /* 填写 Redis 连接等 */ }
consumerOpt := &redisqueue.ConsumerOptions{ /* group/stream 等 */ }

r, _ := q.NewRedis(producerOpt, consumerOpt)

msg := &q.Message{}
msg.SetStream("demo")
msg.SetID("1")
msg.SetValues(map[string]interface{}{"hello": "world"})
_ = r.Append(msg)

r.Register("demo", func(m storage.Messager) error {
    // 消费逻辑
    return nil
})

r.Run()
```

- NSQ

```go
import (
    "github.com/nsqio/go-nsq"
    q "github.com/toujourser/utils/storage/queue"
)

nsqAddrs := []string{"127.0.0.1:4150"}
nsqCfg := nsq.NewConfig()
nsqQ, _ := q.NewNSQ(nsqAddrs, nsqCfg, "ch-")

// 生产
_ = nsqQ.Append(msg)

// 消费（topic 即 name）
nsqQ.Register("demo", func(m storage.Messager) error {
    return nil
})
```

### 6. 飞书 WebHook（im_hook）

```go
import hook "github.com/toujourser/utils/im_hook"

h := hook.NewFeishuHook("<secret>", "https://open.feishu.cn/open-apis/bot/v2/hook/xxx")
_ = h.Push(hook.FText, "标题或文本", "")
```

### 7. 加解密与哈希（encrypt）

```go
import b "github.com/toujourser/utils/encrypt/basic"

_ = b.Base64StdEncode("hi")
_ = b.Md5String("abc")
_ = b.HmacSha256("data", "key")
```

更多 AES/DES/RSA/Hash 等请查看 encrypt 子包。

### 8. 数据校验（validation）

```go
import (
    "errors"
    v "github.com/toujourser/utils/validation"
)

// validator 在包初始化时已注册中文或英文翻译（默认中文）
if err := errors.New("mock validation error"); err != nil {
    msg := v.Error(err) // 生成可读性较好的提示
    _ = msg
}
```

### 9. 集合操作（slice）与流式处理（stream）

```go
import s "github.com/toujourser/utils/slice"

nums := []int{1,2,3,4}
odd := s.Filter(nums, func(_ int, x int) bool { return x%2==1 }) // [1,3]

// 流式
import st "github.com/toujourser/utils/stream"
res := st.FromRange(1, 10, 1).Filter(func(x int) bool { return x%2==0 }).ToSlice()
```

### 10. 并发模式（concurrency）

```go
import (
    "context"
    c "github.com/toujourser/utils/concurrency"
)

ch := c.NewChannel[int]()
ctx, cancel := context.WithCancel(context.Background())
values := ch.Take(ctx, ch.Repeat(ctx, 1, 2, 3), 5)
for v := range values { _ = v }
cancel()
```

### 11. 邮件（mail）

```go
import m "github.com/toujourser/utils/mail"

_ = m.Send(&m.Options{
    MailHost: "smtp.example.com",
    MailPort: 465,
    MailUser: "noreply@example.com",
    MailPass: "password",
    MailTo:   "user1@example.com,user2@example.com",
    Subject:  "主题",
    Body:     "<b>内容</b>",
})
```

### 12. UUID / 雪花 / 随机

```go
import u "github.com/toujourser/utils/uuid"

id := u.MustString()            // UUID
sf := u.GetSnowflakeId()        // 雪花 ID
code := u.CreateValidateCode(6) // 6 位数字验证码
```

### 13. WebSocket（ws）

提供连接管理、分组广播、单发与全局广播、与 gin 集成的示例。典型用法：

- 使用 WebsocketManager 进行 Register/UnRegister
- 通过 Send/SendGroup/SendAll 推送消息
- WsClient/WsLogout 提供 gin handler

详情见 ws 包源码与注释。

### 14. Google Gemini 示例（gemini）

```go
// 参见 gemini/test/gemini.go
model := client.GenerativeModel("gemini-pro")
iter := model.GenerateContentStream(ctx, genai.Text("使用中文讲一个鬼故事"))
```

## 依赖与第三方库（节选）

- 日志：go.uber.org/zap、github.com/sirupsen/logrus、gopkg.in/natefinch/lumberjack.v2
- 缓存/队列：github.com/go-redis/redis、github.com/robinjoseph08/redisqueue/v2、github.com/nsqio/go-nsq
- 校验：github.com/go-playground/validator/v10、github.com/gin-gonic/gin/binding
- WebSocket：github.com/gorilla/websocket、github.com/gin-gonic/gin
- 邮件：gopkg.in/gomail.v2
- 结构体复制：github.com/jinzhu/copier、github.com/ulule/deepcopier
- AI：github.com/google/generative-ai-go

完整依赖版本请参考 go.mod。

## 适用场景

- 快速搭建中小型服务的基础能力（日志、缓存、队列、HTTP、校验等）
- 需要通用泛型集合工具提升日常业务开发效率
- 与现有框架（gin、beego 等）搭配，通过 utils 作为工具层

## 版本与兼容性

- Go 1.21 及以上
- 遵循语义化版本（SemVer），尽量保持向后兼容。涉及破坏性调整会在 release note 标注。

## 贡献指南

- 欢迎 Issue/PR，建议附带用例与文档片段
- 保持 API 简洁、注释明确，遵从已有代码风格
- 若引入新第三方库，请说明选择理由与替代方案

## 许可证

GPL，详见根目录 LICENSE。

## 致谢

- 感谢 JetBrains 为开源项目提供的免费授权
- 也感谢所有参与贡献与反馈的开发者
