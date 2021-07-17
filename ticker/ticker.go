package main

import (
	"log"
	"time"
)

const RetriesNumber = 10

func task() {
	ti := time.NewTicker(time.Second)
	retries := 0
	for {
		<-ti.C
		log.Println("exec task...")
		if retries < RetriesNumber {
			retries++
			log.Println("retry: ", retries)
			ti = time.NewTicker(time.Second)
			continue
		}
		retries = 0
		ti = time.NewTicker(time.Second * 10)
	}
}

func main() {
	task()
}
