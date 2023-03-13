package common

import (
	"fmt"
	"os"
)

func HandleError(msg string, err error) {
	fmt.Println(msg + ":")
	fmt.Println(err)
}

func DebugMsg(msg string) {
	debug := os.Getenv("DEBUG_MODE")
	if debug == "1" {
		fmt.Println(msg)
	}
}
