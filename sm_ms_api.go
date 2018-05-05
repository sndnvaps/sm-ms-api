package sm_ms_api

//sm.ms api for golang

import (
	"bytes"
	"encoding/json"
	"github.com/antchfx/htmlquery"
    	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"io/ioutil"
	"path/filepath"

)

type  ErrMsgBody struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
}

type MsgBody struct {
	Code string `json:"code"`
	Data DataInfo `json:"data"`
}

type HistorymsgBody struct {
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
func Upload(filename string)  {
	url := "https://sm.ms/api/upload"

	var errmsg ErrMsgBody
	var  msg MsgBody

	buf := new (bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fn , _ := filepath.Abs(filename)
	formFile , err := writer.CreateFormFile("smfile",fn)

	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	//把文件读取并定稿表单
	srcFile, err := os.Open(fn)
	if err != nil {
		log.Fatalf("Open source file failed: %s\n", err)
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file failed: %s\n", err)
	}

	writer.WriteField("ssl","0")
	
	//发送表单
	contentType := writer.FormDataContentType()
	writer.Close() //发送之前必须调用Close()以写入结尾行
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n",err)
	} 
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body , _ := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &msg); err == nil {
			fmt.Printf("Upload %s\n", msg.Code)
			fmt.Printf("Filename: %s\n", msg.Data.FileName)
			fmt.Printf("FileInfo: %d x %d\n", msg.Data.Width, msg.Data.Height)
			fmt.Printf("StoreName: %s\n", msg.Data.StoreName)
			fmt.Printf("Size: %d\n", msg.Data.Size)
			fmt.Printf("Path: %s\n",msg.Data.Path)
			fmt.Printf("Hash: %s\n", msg.Data.Hash)
			fmt.Printf("TimeStamp: %d\n", msg.Data.TimeStamp)
			fmt.Printf("Url: %s\n",msg.Data.Url)
			fmt.Printf("Delete url link: %s\n", msg.Data.Delete)
		} else {
			fmt.Println(err)
		}
	} else {
		body , _ := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &errmsg); err == nil {
			fmt.Printf("ErrCode = %s\n", errmsg.Code)
			fmt.Printf("Msg = %s\n", errmsg.Msg)
		}
	}

}

//用指定的 删除地址来 删除已经上传的图片
func Delete(delUrlLink string) {


	url := delUrlLink
	doc, _ := htmlquery.LoadURL(url)
	resp_msg := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='container']"))
	fmt.Printf("Delete [%s] from  https://sm.ms\n",url)
	fmt.Printf("%s\n", resp_msg)
}

//获得过去一小时内上传的文件列表
func ListUploadHistory() {
	var  msg HistorymsgBody
	resp, err := http.Get("https://sm.ms/api/list")
	if err != nil {
		log.Fatalf("List the last 1 hour file you upload to https://sm.ms err : %s\n",err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		num := len(msg.Data)
		for i := 0; i < num; i++ {
			fmt.Printf("----------------------------\n")
			fmt.Printf("Id = %d\n", i+1)
			fmt.Printf("Filename: %s\n", msg.Data[i].FileName)
			fmt.Printf("FileInfo: %d x %d\n", msg.Data[i].Width, msg.Data[i].Height)
			fmt.Printf("StoreName: %s\n", msg.Data[i].StoreName)
			fmt.Printf("Size: %d\n", msg.Data[i].Size)
			fmt.Printf("Path: %s\n",msg.Data[i].Path)
			fmt.Printf("Hash: %s\n", msg.Data[i].Hash)
			fmt.Printf("TimeStamp: %d\n", msg.Data[i].TimeStamp)
			fmt.Printf("Url: %s\n",msg.Data[i].Url)
			fmt.Printf("Delete url link: %s\n", msg.Data[i].Delete)
			fmt.Printf("----------------------------\n")
		}
	}
}

