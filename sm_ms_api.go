package sm_ms_api

//sm.ms api for golang

import (
	"bytes"
	"encoding/json"
	"github.com/antchfx/htmlquery"
    	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"io/ioutil"
	"path/filepath"

)

//所有错误返回
type  ErrMsgBody struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
}

//用于 Upload() func
type MsgBody struct {
	Code string `json:"code"`
	Data DataInfo `json:"data"`
}

//用于 ListUploadHistory() func
type HistoryMsgBody struct {
	Code string `json:"code"`
	Data []DataInfo `json:"data"`
}
type DataInfo struct {
	Width int `json:"width"`
	Height int `json:"height"`
	FileName string `json:"filename"`
	StoreName string `json:"storename"`
	Size int `json:"size"`
	Path string `json:"path"`
	Hash string `json:"hash"`
	TimeStamp int64 `json:"timestamp"`
	Ip string `json:"ip"`
	Url string `json:"url"`
	Delete string `json:"delete"`
}

//doc link  https://sm.ms/doc/
func Upload(filename string) (map[string]interface{}, error) {
	url := "https://sm.ms/api/upload"
	status := make(map[string]interface{}) //因为返回值有两个类型， 一个为 ErrMsgBody, 一个为MsgBody

	var errmsg ErrMsgBody
	var  msg MsgBody

	buf := new (bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fn , _ := filepath.Abs(filename)
	formFile , err := writer.CreateFormFile("smfile",fn)

	if err != nil {
		return status,err
	}

	//把文件读取并定稿表单
	srcFile, err := os.Open(fn)
	if err != nil {
		return status,err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		return status,err
	}

	writer.WriteField("ssl","0")
	
	//发送表单
	contentType := writer.FormDataContentType()
	writer.Close() //发送之前必须调用Close()以写入结尾行
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		return status,err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body , _ := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &msg); err == nil {
			status["msg"] = msg
			return status,nil
		} else {
			return status, err
		}
	} else {
		body , _ := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &errmsg); err == nil {
			status["msg"] = errmsg
			return status, nil
		}
	}
	return status, nil

}

//用指定的 删除地址来 删除已经上传的图片
func Delete(delUrlLink string) string {


	url := delUrlLink
	doc, _ := htmlquery.LoadURL(url)
	resp_msg := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='container']"))
	fmt.Printf("Delete [%s] from  https://sm.ms\n",url)
	return resp_msg
}

//获得过去一小时内上传的文件列表
func ListUploadHistory() (HistoryMsgBody, error) {
	var msg  HistoryMsgBody
	resp, err := http.Get("https://sm.ms/api/list")
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

	return msg, nil
}

