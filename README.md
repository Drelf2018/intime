# intime
 
基于 `int64` 毫秒级时间戳实现的时间类，功能强大。

### 示例

使用 `intime.New` 获取一个对象，当参数为 `nil` 时返回当前时间。

允许的参数类型 `[]byte` `string` `int` `int8` `int16` `int32` `int64` `uint` `uint8` `uint16` `uint32` `uint64` `float32` `float64` `time.Duration` 以及 `MilliTimestamp` `NanoTimestamp` `UnixTimestamp` 接口。

```go
type UnixTimestamp interface {
	Unix() int64
}

type MilliTimestamp interface {
	UnixMilli() int64
}

type NanoTimestamp interface {
	UnixNano() int64
}
```

#### 测试

```go
func TestNew(t *testing.T) {
	d := intime.New(nil)
	t.Log(d.Time())

	d = intime.New("20240321", "format", "%Y%m%d")
	t.Log(d)
}
```

#### 控制台

```
=== RUN   TestNew
	intime_test.go:13: 2024-03-21 13:28:11.704 +0800 CST
    intime_test.go:16: 1710979200000
--- PASS: TestNew (0.01s)
```

### 高级

`Time` 对象有许多方法。

#### 测试

```go
type Post struct {
	Time intime.Time
}

func TestPost(t *testing.T) {
	post := Post{Time: intime.Now()}
	post.Time.Format()
	tmpl, err := template.New("example").Parse(`{{.Time.Format "%H:%M:%S"}}`)
	if err != nil {
		t.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, post)
	if err != nil {
		t.Fatal(err)
	}
}
```

#### 控制台

```
=== RUN   TestPost
13:28:11
--- PASS: TestPost (0.00s)
```