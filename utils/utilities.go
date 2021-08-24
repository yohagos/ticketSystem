package utils

import (
	"log"
	"time"
)

func IsError(err error) {
	if err != nil {
		log.Println("error occurred : ", err)
	}
}

func CreateTimeStamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}
