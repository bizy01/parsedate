### 介绍

该项目是一个关于go语言的时间字符串parse解析包，主要用于解决go语言原生的时间字符串解析问题。

使用了goyacc词法解析工具对各种时间字符串进行语法定义，手动实现了lex词法扫描器

### 功能

- 时间字符串parse

### 支持的字符串

#### 日期格式

- oct 7, 1970
- May 17, 2012
- oct 7, '70
- Oct 7, '70
- Oct. 7, '70
- oct. 7, '70
- oct. 7, 1970
- 7 oct 70
- 7 oct 1970
- 7 September 1970
- September 17, 2012
- 03 Jul 2017
- 03-Jul-15
- 03-Jul 2015
- 3-Jul-15
- 3 Jan 2021
- 2014年04月08日
- 03/31/2014
- 3/31/2014
- 3/5/2014
- 08/08/71
- 8/8/71
- 2014/04/02
...

#### 时间格式

todo

### 使用
todo

### demo
todo

## 参考
[dateparse](https://github.com/araddon/dateparse)
[promql](https://github.com/prometheus/prometheus/tree/main/promql/parser)

