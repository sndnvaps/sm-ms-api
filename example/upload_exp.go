package main

import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("let's upload the file to https://sm.ms\n")
	status, err := sm_ms_api.Upload("test.jpg")
	if err != nil {
		fmt.Println(err)
	} else {

		for _, val := range status {
			switch val.(type) {
			case sm_ms_api.ErrMsgBody:
				body := val.(sm_ms_api.ErrMsgBody)
				fmt.Printf("ErrCode = %s\n", body.Code)
				fmt.Printf("Msg = %s\n", body.Msg)
			case sm_ms_api.MsgBody:
				body := val.(sm_ms_api.MsgBody)
				fmt.Printf("Upload %s\n", body.Code)
				fmt.Printf("Filename: %s\n", body.Data.FileName)
				fmt.Printf("FileInfo: %d x %d\n", body.Data.Width, body.Data.Height)
				fmt.Printf("StoreName: %s\n", body.Data.StoreName)
				fmt.Printf("Size: %d\n", body.Data.Size)
				fmt.Printf("Path: %s\n", body.Data.Path)
				fmt.Printf("Hash: %s\n", body.Data.Hash)
				fmt.Printf("TimeStamp: %d\n", body.Data.TimeStamp)
				fmt.Printf("Url: %s\n", body.Data.Url)
				fmt.Printf("Delete: %s\n", body.Data.Delete)
			}
		}
	}

}
