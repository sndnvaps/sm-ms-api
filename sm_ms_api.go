package sm_ms_api

//sm.ms api for golang

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
	//"sync"
)

// API登录，返回token
// 用于 GenToken() func
type LoginBody struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

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

// Check file suffix , only support jpeg,jpg,png,gif,bmp
func CheckFileSuffix(filename string) (bool, string) {
	suffix := path.Ext(filename)
	if (suffix == ".jpeg") || (suffix == ".jpg") ||
		(suffix == ".png") || (suffix == ".gif") ||
		(suffix == ".bmp") {
		return true, suffix
	}

	return false, suffix
}

func GenToken(usr, pwd string) (MsgBody, error) {
	tmpurl := "https://sm.ms/api/v2/token"
	data := url.Values{}
	data.Set("username", usr)
	data.Set("password", pwd)

	var msg MsgBody
	msg = MsgBody{
		Code:    "error",
		Message: " Internal function error",
	}
	resp, err := http.PostForm(tmpurl, data)
	if err != nil {
		fmt.Println(err.Error())
		return msg, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		//		fmt.Printf("GetToken::Body -> [%s]", string(body))
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return msg, err
		}

	}
	return msg, nil
}

// doc link  https://sm.ms/doc/
func Upload(filename string, token string) (MsgBody, error) {
	fmt.Printf("upload func::token= [%s]", token)
	//mutex.Lock()
	tmpurl := "https://sm.ms/api/v2/upload"

	var msg MsgBody

	msg = MsgBody{
		Code:    "Error",
		Message: "Internal function error",
	}

	if isPic, suffix := CheckFileSuffix(filename); isPic != true {
		errmsg := fmt.Sprintf("File has an invalid extension %s\nsupport file ext is jpeg,jpg,png,gif,bmp\n", suffix)
		err := errors.New(errmsg)
		return msg, err
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	fn, _ := filepath.Abs(filename)
	formFile, err := writer.CreateFormFile("smfile", fn)

	if err != nil {
		return msg, err
	}

	//把文件读取并定稿表单
	srcFile, err := os.Open(fn)
	if err != nil {
		return msg, err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		return msg, err
	}

	//发送表单
	contentType := writer.FormDataContentType()
	writer.Close() //发送之前必须调用Close()以写入结尾行
	req, err := http.NewRequest("POST", tmpurl, &buf)
	//mutex.Unlock()
	if err != nil {
		return msg, err
	}
	req.Header.Set("Content-type", contentType)
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()

	//fmt.Printf("upload::StatusCode -> [%d]", resp.StatusCode)
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return msg, err
		}

	}

	return msg, nil
}

// 获得过去一小时内上传的文件列表
func ListHistory() (SliceMsgBody, error) {
	var msg SliceMsgBody
	resp, err := http.Get("https://sm.ms/api/v2/history")
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

}

// 提供 API Token，获得对应用户的所有上传图片信息.
func ListUserHistory(token string, page int) (SliceMsgBody, error) {
	var msg SliceMsgBody
	tmpurl := "https://sm.ms/api/v2/upload_history"
	req, err := http.NewRequest("POST", tmpurl, nil)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Authorization", token)
	if err != nil {
		return msg, err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

}

// 提供 API Token，获得对应用户的基本信息.
func ListUserProfile(token string) (MsgBody, error) {
	var msg MsgBody
	fmt.Printf("debug for ListUserProfile: token= [%s]", token)

	tmpurl := "https://sm.ms/api/v2/profile"
	req, err := http.NewRequest("POST", tmpurl, nil)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Authorization", token)

	if err != nil {
		fmt.Printf("debug in err1: msg = [%s],err info = [%s]", msg.Message, err.Error())
		return msg, err
	}
	//defer req.Body.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("debug in err1: msg = [%s],err info = [%s]", msg.Message, err.Error())
		return msg, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		fmt.Printf("deubug for lup: StatusCode = [%d]", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("deubug for lup: body=[%s]", string(body))
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

}

// 用指定的 删除地址来 删除已经上传的图片
func Delete(hash string) (MsgBody, error) {

	tmpurl := "https://sm.ms/api/v2/delete/" + hash
	var msg MsgBody
	resp, err := http.Get(tmpurl)
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}
}

func Clear() (MsgBody, error) {
	var msg MsgBody
	resp, err := http.Get("https://sm.ms/api/v2/clear")
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}
}
