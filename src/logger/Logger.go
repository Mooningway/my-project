package logger

import (
	"fmt"
	"log"
)

func Print(msg string, args ...interface{}) {
	out := fmt.Sprintf(msg, args...)
	log.Println(out)
}
