package libs

import (
	"log"
	"os"
)

func GetWorkPath() string {
	wp, err := os.Getwd()
	if err != nil {
		log.Print("cannot get work path")
	}
	return wp
}
