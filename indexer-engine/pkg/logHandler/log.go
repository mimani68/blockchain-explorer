package logHandler

import (
	"encoding/json"
	"fmt"
)

const (
	INFO  = "INFO"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

func Log(level string, msg string, tag ...string) {
	if tag != nil {
		fmt.Println("[" + level + "][" + tag[0] + "] " + msg)
	} else {
		fmt.Println("[" + level + "]" + msg)
	}
}

func Debug(msg string, payload interface{}, tag ...string) {
	payloadString, _ := json.Marshal(payload)
	if tag != nil {
		fmt.Println("[DEBUG]["+tag[0]+"] "+msg+" data:", string(payloadString))
	} else {
		fmt.Println("[DEBUG]" + msg)
	}
}
