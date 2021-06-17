### 介绍

该项目是一个关于go语言的时间字符串parse解析包，主要用于解决go语言原生的时间字符串解析问题。

使用了goyacc词法解析工具对各种时间字符串进行语法定义，手动实现了lex词法扫描器

### 功能

- 时间字符串parse

### 支持的字符串

#### 日期格式

```
    {in: "10:20:30", out: 1623896430},
	{in: "2021-1-1", out: 1609430400},
	{in: "2021-1-1 10:20:30", out: 1609467630},
	{in: "2021-01-01 10:20:30", out: 1609467630},
	{in: "oct 7, 1970 10:20:30", out: 24114030},
	{in: "oct. 7, 1970 10:20:30", out: 24114030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "7 oct 70", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "03 February 2013", out: 1359820800},
	{in: "2021年10月1日", out: 1633017600},
	{in: "03/31/2014", out: 1396195200},
	{in: "2021/04/02", out: 1617292800},
```

### 使用

```
type dateTest struct {
	in           string
	out          int64
	err          bool
}

var testInputs = []dateTest{
	{in: "10:20:30", out: 1623896430},
	{in: "2021-1-1", out: 1609430400},
	{in: "2021-1-1 10:20:30", out: 1609467630},
	{in: "2021-01-01 10:20:30", out: 1609467630},
	{in: "oct 7, 1970 10:20:30", out: 24114030},
	{in: "oct. 7, 1970 10:20:30", out: 24114030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "7 oct 70", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "03 February 2013", out: 1359820800},
	{in: "2021年10月1日", out: 1633017600},
	{in: "03/31/2014", out: 1396195200},
	{in: "2021/04/02", out: 1617292800},
}

func TestParse(t *testing.T) {
	for i, th := range testInputs {
		expected, err := Timestamp(th.in)
		if err != nil {
			t.Log(err)
			continue
		}

		require.Equal(t, th.out, expected, "%d: input %q", i, th.in)
	}
}
```

### demo
todo

## 参考

- [dateparse](https://github.com/araddon/dateparse)
- [promql](https://github.com/prometheus/prometheus/tree/main/promql/parser)

