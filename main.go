package myhim

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("INFO- %v", message)
}

func LogWarning(message string) {
	log.Printf("ERROR- %v", message)
}
