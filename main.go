package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	fmt.Println("App running here")

	// info log
	log.WithFields(log.Fields{
		"animal": "cat",
	}).Info("A cat appear")

	// warning log
	log.WithFields(log.Fields{
		"animal": "cat",
	}).Warn("A cat run")

	// another example
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}
