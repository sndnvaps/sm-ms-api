package main

import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

//Test is file
func IsFile(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

//Test is Dir
func IsDir(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}

	return f.IsDir()
}

//IsExists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func List(c *cli.Context) error {
	history, err := sm_ms_api.ListUploadHistory()
	if err == nil {
		num := len(history.Data)
		for i := 0; i < num; i++ {
			fmt.Printf("----------------------------\n")
			fmt.Printf("Id = %d\n", i+1)
			fmt.Printf("Filename: %s\n", history.Data[i].FileName)
			fmt.Printf("FileInfo: %d x %d\n", history.Data[i].Width, history.Data[i].Height)
			fmt.Printf("StoreName: %s\n", history.Data[i].StoreName)
			fmt.Printf("Size: %d\n", history.Data[i].Size)
			fmt.Printf("Path: %s\n", history.Data[i].Path)
			fmt.Printf("Hash: %s\n", history.Data[i].Hash)
			fmt.Printf("TimeStamp: %d\n", history.Data[i].TimeStamp)
			fmt.Printf("Url: %s\n", history.Data[i].Url)
			fmt.Printf("Delete url link: %s\n", history.Data[i].Delete)
			fmt.Printf("----------------------------\n")
		}

	} else {
		return err
	}

	return nil
}

func Delete(c *cli.Context) error {
	delurl := c.Args().First() //获取第一个参数
	resp := sm_ms_api.Delete(delurl)
	fmt.Println(resp)
	return nil
}

func Upload(c *cli.Context) error {

	path := c.Args().First()

	if IsDir(path) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, file := range files {
			if IsPic, _ := sm_ms_api.CheckFileSuffix(file.Name()); IsPic == true {
				var FullPath string
				if strings.HasSuffix(path, "/") {
					FullPath = path + file.Name()
				} else {
					FullPath = path + "/" + file.Name()
				}
				msg, err := sm_ms_api.Upload(FullPath)
				if err != nil {
					return err
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
		}
	} else {
		msg, err := sm_ms_api.Upload(path)
		if err != nil {
			return err
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

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "sm_ms_tools"
	app.Compiled = time.Now()
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Jimes Yang",
			Email: "sndnvaps@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Jimes Yang<sndnvaps@gmail.com>"
	app.Usage = "A tool for sm.ms"
	app.Commands = []cli.Command{
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete the pic you upload to sm.ms ",
			Action:  Delete,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list the upload history you upload to sm.ms",
			Action:  List,
		},
		{
			Name:    "upload",
			Aliases: []string{"u", "up"},
			Usage:   "upload the file or the files in the folder to sm.ms",
			Action:  Upload,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
