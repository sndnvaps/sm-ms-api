package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	sm_ms_api "github.com/sndnvaps/sm-ms-api"
	"github.com/urfave/cli"
)

func Clear(c *cli.Context) error {
	resp, err := sm_ms_api.Clear()
	if err == nil {
		fmt.Printf("success: %v\n", resp.Success)
		fmt.Printf("Code: %s\n", resp.Code)
		fmt.Printf("message: %s\n", resp.Message)
	} else {
		fmt.Printf("success: %v\n", resp.Success)
		fmt.Printf("Code: %s\n", resp.Code)
		fmt.Printf("message: %s\n", resp.Message)
	}
	fmt.Printf("RequestId: %s\n", resp.RequestId)
	return nil
}

func Delete(c *cli.Context) error {
	hash := c.String("hash")
	resp, err := sm_ms_api.Delete(hash)
	if err == nil {
		fmt.Printf("success: %v\n", resp.Success)
		fmt.Printf("Code: %s\n", resp.Code)
		fmt.Printf("message: %s\n", resp.Message)
	} else {
		fmt.Printf("success: %v\n", resp.Success)
		fmt.Printf("Code: %s\n", resp.Code)
		fmt.Printf("message: %s\n", resp.Message)
	}
	fmt.Printf("RequestId: %s\n", resp.RequestId)
	return nil
}

// Test is file
func IsFile(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// Test is Dir
func IsDir(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// IsExists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func List(c *cli.Context) error {
	token := c.String("token")
	var history sm_ms_api.SliceMsgBody
	var err error
	if "" == token {
		history, err = sm_ms_api.ListHistory()
	} else {
		history, err = sm_ms_api.ListUserHistory(token)
	}
	if err == nil {
		num := len(history.Data)
		//fmt.Printf("len(history.Data) = %d ", len(history.Data))
		for i := 0; i < num; i++ {
			data := history.Data[i]
			var datainfo sm_ms_api.DataInfo
			err := mapstructure.Decode(data, &datainfo) //map[string]interface{} -> struct
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("----------------------------\n")
			fmt.Printf("Id = %d\n", datainfo.FileId)
			fmt.Printf("Filename: %s\n", datainfo.FileName)
			fmt.Printf("FileInfo: %d x %d\n", datainfo.Width, datainfo.Height)
			fmt.Printf("StoreName: %s\n", datainfo.StoreName)
			fmt.Printf("Size: %d\n", datainfo.Size)
			fmt.Printf("Path: %s\n", datainfo.Path)
			fmt.Printf("Hash: %s\n", datainfo.Hash)
			fmt.Printf("Url: %s\n", datainfo.Url)
			fmt.Printf("Delete url link: %s\n", datainfo.Delete)
			fmt.Printf("Page link: %s\n", datainfo.Page)
			fmt.Printf("----------------------------\n")
		}

	} else {
		fmt.Printf("ListUserHistory:error = [%s]", err.Error())
		return err
	}

	return nil
}

func Upload(c *cli.Context) error {

	path := c.Args().First()
	token := c.String("token")
	//	fmt.Printf("Debug in upload : path= [%s],token= [%s]", path, token)
	if IsDir(path) {
		files, err := os.ReadDir(path)
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
				msg, err := sm_ms_api.Upload(FullPath, token)
				if err != nil {
					return err
				} else {
					fmt.Printf("Upload %s\n", msg.Code)
					fmt.Printf("Msg = %s\n", msg.Message)

					data := msg.Data

					var datainfo sm_ms_api.DataInfo
					mapstructure.Decode(data, &datainfo) //map[string]interface{} -> struct
					fmt.Printf("----------------------------\n")
					fmt.Printf("Id = %d\n", datainfo.FileId)
					fmt.Printf("Filename: %s\n", datainfo.FileName)
					fmt.Printf("FileInfo: %d x %d\n", datainfo.Width, datainfo.Height)
					fmt.Printf("StoreName: %s\n", datainfo.StoreName)
					fmt.Printf("Size: %d\n", datainfo.Size)
					fmt.Printf("Path: %s\n", datainfo.Path)
					fmt.Printf("Hash: %s\n", datainfo.Hash)
					fmt.Printf("Url: %s\n", datainfo.Url)
					fmt.Printf("Delete url link: %s\n", datainfo.Delete)
					fmt.Printf("Page link: %s\n", datainfo.Page)
					fmt.Printf("RequestId = %s\n", msg.RequestId)
				}
			}
		}
	} else {
		msg, err := sm_ms_api.Upload(path, token)
		if err != nil {
			return err
		} else {
			fmt.Printf("Upload %s\n", msg.Code)
			fmt.Printf("Msg = %s\n", msg.Message)
			data := msg.Data
			var datainfo sm_ms_api.DataInfo
			//fmt.Println(data) //for debug
			mapstructure.Decode(data, &datainfo) //map[string]interface{} -> struct
			fmt.Printf("----------------------------\n")
			fmt.Printf("Id = %d\n", datainfo.FileId)
			fmt.Printf("Filename: %s\n", datainfo.FileName)
			fmt.Printf("FileInfo: %d x %d\n", datainfo.Width, datainfo.Height)
			fmt.Printf("StoreName: %s\n", datainfo.StoreName)
			fmt.Printf("Size: %d\n", datainfo.Size)
			fmt.Printf("Path: %s\n", datainfo.Path)
			fmt.Printf("Hash: %s\n", datainfo.Hash)
			fmt.Printf("Url: %s\n", datainfo.Url)
			fmt.Printf("Delete url link: %s\n", datainfo.Delete)
			fmt.Printf("Page link: %s\n", datainfo.Page)
			fmt.Printf("RequestId = %s\n", msg.RequestId)
		}
	}

	return nil
}

func ListUserProfile(c *cli.Context) error {
	token := c.String("token")

	msg, err := sm_ms_api.ListUserProfile(token)

	if err != nil {
		return err
	} else {
		fmt.Printf("Http Code %s\n", msg.Code)
		fmt.Printf("Msg = %s\n", msg.Message)

		data := msg.Data
		var usr_profile sm_ms_api.UserProfile
		mapstructure.Decode(data, &usr_profile) //map[string]interface{} -> struct
		fmt.Printf("Username: %s\n", usr_profile.Username)
		fmt.Printf("Email: %s\n", usr_profile.Email)
		fmt.Printf("Role: %s\n", usr_profile.Role)
		fmt.Printf("Group_expire: %s\n", usr_profile.GroupExpire)
		fmt.Printf("Disk_Usage: %s\n", usr_profile.DiskUsage)
		fmt.Printf("Disk_Usage_Raw: %d\n", usr_profile.DiskUsageRaw)
		fmt.Printf("DiskLimit: %s\n", usr_profile.DiskLimit)
		fmt.Printf("DiskLimitRaw: %d\n", usr_profile.DiskLimitRaw)
		fmt.Printf("RequestId = %s\n", msg.RequestId)
	}
	return nil
}

func Login(c *cli.Context) error {
	usr := c.String("username")
	pwd := c.String("password")

	msg, err := sm_ms_api.GenToken(usr, pwd)

	if err != nil {
		return err
	} else {
		fmt.Printf("Upload %s\n", msg.Code)
		fmt.Printf("Msg = %s\n", msg.Message)

		data := msg.Data

		var token sm_ms_api.Authorization
		mapstructure.Decode(data, &token) //map[string]interface{} -> struct

		fmt.Printf("Authorization: %s\n", token.Token)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "sm_ms_tools"
	app.Compiled = time.Now()
	app.Version = "2.1.1"
	app.Authors = []cli.Author{
		{
			Name:  "Jimes Yang",
			Email: "sndnvaps@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 - 2023 Jimes Yang<sndnvaps@gmail.com>"
	app.Usage = "A tool for sm.ms"
	app.Commands = []cli.Command{
		{
			Name:    "clear",
			Aliases: []string{"c", "cls"},
			Usage:   "Clear the history file list log you upload to sm.ms",
			Action:  Clear,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "hash",
					Usage: "hash_id of the sm.ms picture",
				},
			},
			Usage:  "delete the pic you upload to sm.ms ",
			Action: Delete,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list the upload history you upload to sm.ms",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "token",
					Usage: "with token can show the user's upload pic info(it can be empty)",
				},
			},
			Action: List,
		},
		{
			Name:    "listusrprofile",
			Aliases: []string{"lup"},
			Usage:   "list the user_profile of sm.ms with token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "token",
					Usage: "with token can show the user's Profile",
				},
			},
			Action: ListUserProfile,
		},
		{
			Name:    "upload",
			Aliases: []string{"u", "up"},
			Usage:   "upload the file or the files in the folder to sm.ms(with api token)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "token",
					Usage: "api token (it can't be empty)",
				},
			},
			Action: Upload,
		},
		{
			Name:  "login",
			Usage: "use the username,password to get the api token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "username, usr",
					Usage: "username of sm.ms",
				},
				cli.StringFlag{
					Name:  "password, pwd",
					Usage: "password of sm.ms",
				},
			},
			Action: Login,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
