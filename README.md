# sm_ms_api
api for  https://sm.ms

# Golang 版本的 https://sm.ms/doc/v2 API接口

# 如何安装 

建议使用go mod 安装
支持go 1.16.x


# 测试代码

 [main.go](cmd/main.go)


# 函数原型

```go

  func Clear() (MsgBody, error)

  func Delete(delUrlLink string) string
  func GenToken(usr, pwd string) (MsgBody, error)

  func Upload(filename string,token string) (MsgBody, error)

  func ListUploadHistory() (SliceMsgBody, error)
  func ListUserHistory(token string, page int) (SliceMsgBody, error)
  func 

```

# 返回值定义结构体

```go
// Authorization, 用于验证用户信息，token
type Authorization struct {
	Token string `json:"token"`
}

// 提供 API Token，获得对应用户的基本信息.
type UserProfile struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	GroupExpire   string `json:"group_expire"`
	DiskUsage     string `json:"disk_usage"`
	EmailVerified int    `json:"email_verified"`
	DiskUsageRaw  int    `json:"disk_usage_raw"`
	DiskLimit     string `json:"disk_limit"`
	DiskLimitRaw  int    `json:"disk_limit_raw"`
}

//锁，用于 Upload() func
//var mutex = &sync.Mutex{}

// 用于 返回信息
type MsgBody struct {
	Success   bool                   `json:"success"`
	Code      string                 `json:"code"`
	Message   string                 `json:"Message"` //用于接收错误信息
	Data      map[string]interface{} `json:"data,omitempty"`
	RequestId string                 `json:"RequestID"`
}

// 用于 返回信息
type SliceMsgBody struct {
	Success   bool                     `json:"success"`
	Code      string                   `json:"code"`
	Message   string                   `json:"Message"` //用于接收错误信息
	Data      []map[string]interface{} `json:"data,omitempty"`
	RequestId string                   `json:"RequestID"`
}

// 用于获取上传图片的信息
type DataInfo struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FileName  string `json:"filename"`
	FileId    int    `json:"file_id,omitempty"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
	Page      string `json:"page"`
}
```

## License
#### [MIT](https://sndnvaps.mit-license.org/2018)
