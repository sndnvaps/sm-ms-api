package main

import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("let's upload the file to https://sm.ms\n")
	msg, err := sm_ms_api.Upload("test.jpg")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Upload %s\n", msg.Code)
		if msg.Msg != "" {
			fmt.Printf("Msg = %s\n", msg.Msg)
		} else {
			fmt.Printf("Filename: %s\n", msg.Data.FileName)
			fmt.Printf("FileInfo: %d x %d\n", msg.Data.Width, msg.Data.Height)
			fmt.Printf("StoreName: %s\n", msg.Data.StoreName)
			fmt.Printf("Size: %d\n", msg.Data.Size)
			fmt.Printf("Path: %s\n", msg.Data.Path)
			fmt.Printf("Hash: %s\n", msg.Data.Hash)
			fmt.Printf("TimeStamp: %d\n", msg.Data.TimeStamp)
			fmt.Printf("Url: %s\n", msg.Data.Url)
			fmt.Printf("Delete: %s\n", msg.Data.Delete)
		}

	}

}
