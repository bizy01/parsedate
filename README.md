### 介绍

该项目是一个关于go语言的时间字符串parse解析包，主要用于解决go语言原生的时间字符串解析问题。

使用了goyacc词法解析工具对各种时间字符串进行语法定义，手动实现了lex词法扫描器

### 功能

- 时间字符串parse

### 支持的字符串

#### 日期格式
   - ANSIC       = "Mon Jan _2 15:04:05 2006"
   - UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
   - RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
   - RFC822      = "02 Jan 06 15:04 MST"
   - RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
   - RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
   - RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
   - RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
   - RFC3339     = "2006-01-02T15:04:05Z07:00"
   - RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
   - more than ...

### 使用
todo

### demo
todo

## 参考

- [dateparse](https://github.com/araddon/dateparse)
- [promql](https://github.com/prometheus/prometheus/tree/main/promql/parser)

