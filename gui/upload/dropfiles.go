// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/sndnvaps/sm_ms_api"
)

func IsDir(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func upload(paths string) string {
	var tmp string
	//检查当前路径是否为 目录
	if IsDir(paths) {
		return "This need a image file not a folder!"
	}

	if IsPic, _ := sm_ms_api.CheckFileSuffix(paths); IsPic == true {
		msg, err := sm_ms_api.Upload(paths)
		if err != nil {
			return err.Error()
		} else {
			tmp = fmt.Sprintf("Upload %s\r\n", msg.Code)
			if msg.Msg != "" {
				tmp += fmt.Sprintf("Msg = %s\r\n", msg.Msg)
			} else {
				tmp += fmt.Sprintf("Filename: %s\r\n", msg.Data.FileName)
				tmp += fmt.Sprintf("FileInfo: %d x %d\r\n", msg.Data.Width, msg.Data.Height)
				tmp += fmt.Sprintf("StoreName: %s\r\n", msg.Data.StoreName)
				tmp += fmt.Sprintf("Size: %d\r\n", msg.Data.Size)
				tmp += fmt.Sprintf("Path: %s\r\n", msg.Data.Path)
				tmp += fmt.Sprintf("Hash: %s\r\n", msg.Data.Hash)
				tmp += fmt.Sprintf("TimeStamp: %d\r\n", msg.Data.TimeStamp)
				tmp += fmt.Sprintf("Url: %s\r\n", msg.Data.Url)
				tmp += fmt.Sprintf("Delete: %s\r\n", msg.Data.Delete)
			}
		}
	}

	return tmp

}
func main() {
	var textEdit *walk.TextEdit
	var outputTextEdit *walk.TextEdit
	MainWindow{
		Title:   "upload file to sm.ms",
		MinSize: Size{320, 240},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			textEdit.SetHeight(40)     //设置本文本框的高度为10
			textEdit.SetText(files[0]) //我们现在只取第一个文件值
		},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &textEdit,
						ReadOnly: true,
						Text:     "Drop files here, from windows explorer...",
					},
					TextEdit{
						AssignTo: &outputTextEdit,
						ReadOnly: false,
					},
					PushButton{
						MaxSize: Size{40, 30},

						Text: "Upload image to sm.ms ",
						OnClicked: func() {
							outputTextEdit.SetHeight(150)
							out := upload(textEdit.Text())
							outputTextEdit.SetText(out)
						},
					},
				},
			},
		},
	}.Run()
}
