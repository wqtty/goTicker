package main

import (
	"time"

	"log"

	"github.com/wqtty/goTicker/ticker"
)

func main() {
	err := ticker.RegisterTicker(3*time.Second, func(t time.Time) {
		log.Println("3s triggered at:", t.String())
	})
	if err != nil {
		log.Println("timer.RegisterTicker failed, err:", err)
	}
	err = ticker.RegisterTicker(5*time.Second, func(t time.Time) {
		log.Println("5s triggered at:", t.String())
	})
	if err != nil {
		log.Println("timer.RegisterTicker failed, err:", err)
	}
	err = ticker.Start()
	if err != nil {
		log.Println("timer.Start returned an error:", err)
	}
}
