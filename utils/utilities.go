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
	return currentTime.Format(time.RFC850)
}
