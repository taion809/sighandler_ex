package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGHUP)

	go func() {
		log.Println("-->> Sighandler Engaged!")
		log.Println("-->> Send me your kill commands.")

		for trap := range sigchan {
			log.Println("!! Got ", trap)
			err := loadConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
}
