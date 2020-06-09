package internal

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Initlog(logLevel string) {
	parsedLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(parsedLevel)
	log.SetFormatter(&log.TextFormatter{ ForceColors: true })
}
