package main
import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("List the last 1 hour file you upload to https://sm.ms\n")
	sm_ms_api.ListUploadHistory()
}