package main
import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("let's upload the file to https://sm.ms\n")
	sm_ms_api.Upload("test.jpg")

}