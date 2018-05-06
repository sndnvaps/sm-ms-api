package main
import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("List the last 1 hour file you upload to https://sm.ms\n")
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
			fmt.Printf("Path: %s\n",history.Data[i].Path)
			fmt.Printf("Hash: %s\n", history.Data[i].Hash)
			fmt.Printf("TimeStamp: %d\n", history.Data[i].TimeStamp)
			fmt.Printf("Url: %s\n",history.Data[i].Url)
			fmt.Printf("Delete url link: %s\n", history.Data[i].Delete)
			fmt.Printf("----------------------------\n")
		} 

	} else {
		fmt.Println(err)
	}
}