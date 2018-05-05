package main
import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	fmt.Printf("Delete file from  https://sm.ms\n")
	sm_ms_api.Delete("https://sm.ms/delete/vJ82UtopYcnrZxq")
}