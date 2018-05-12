package main

import (
	"fmt"
	"github.com/sndnvaps/sm_ms_api"
)

func main() {
	resp, err := sm_ms_api.Clear()
	if err == nil {
		fmt.Printf("Code: %s\n", resp.Code)
		fmt.Printf("Msg: %s\n", resp.Msg)
	} else {
		fmt.Printf("Error msg = %s\n", err.Error())
	}
}