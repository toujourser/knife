## 常用的工具包集合

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/toujourser/utils)

### 背景
utils是基于常见设计模式进行模块化设计的Go包组件集，封装了常用的功能，使用简单，致力于进行快速的业务研发。

### 安装

```bash
$ go get -v -u github.com/toujourser/utils
```

### 用法
在项目中和使用标准库一样在`import`中直接引用即可。

### 功能列表
* `algorithm` 签名算法
* `code` 错误码转换
* `color` 终端日志颜色输出
* `conv` golang类型转换
* `ddm` 字符串隐藏
* `encrypt` 加解密算法
* `file` 文件操作，包括文件下载上传功能
* `im_hook` 飞书WebHook推送
* `logger` 基于Zap和Lumberjack二次封装的日志库
* `mail` 邮件发送
* `request` Http请求封装
* `storage` 文件存储
* `structure` 结构体映射
* `timeparse` 时间解析
* `uuid` UUID生成
* `validation` 数据验证

### 使用许可
[GPL](LICENSE) © TOUJOURSER 2022