package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp" //NTP -- Network Time Protocol
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org") //pool.ntp.org -- сервер, который отправляет нам время с точностью до миллисекунд

	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()

	fmt.Println("Текущее время: ", currentTime.Format(time.RFC3339))
	fmt.Println("Точное время: ", ntpTime.Format(time.RFC3339))
}
