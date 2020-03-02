# sm_ms_api
api for  https://sm.ms

# Golang 版本的 https://sm.ms/doc/v2 API接口

# 如何安装 

建议使用go mod 安装
支持go 1.13.x


# 测试代码

 [main.go](cmd/main.go)


# 函数原型

```go

  func Clear() (MsgBody, error)

  func Delete(delUrlLink string) string

  func Upload(filename string) (map[string]interface{}, error)

  func ListUploadHistory() (HistoryMsgBody, error)

```

# 返回值定义结构体

```go
//用于 Upload() func
type MsgBody struct {
	Code string   `json:"code"`
	Data DataInfo `json:"data,omitempty"`
	Msg  string   `json:"msg,omitempty"` //用于接收错误信息
}

//用于 ListUploadHistory() func
type HistoryMsgBody struct {
	Code string     `json:"code"`
	Data []DataInfo `json:"data"`
}

//用于获取上传图片的信息
type DataInfo struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FileName  string `json:"filename"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	TimeStamp int64  `json:"timestamp"`
	Ip        string `json:"ip"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
}
```

 
# 使用方法，请查看 [example](example) 目录下面的例子


## License
#### [MIT](https://sndnvaps.mit-license.org/2018)
