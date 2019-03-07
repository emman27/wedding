package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emman27/wedding/pkg/rsvp"
	"github.com/emman27/wedding/pkg/rsvp/typeform"
	"github.com/sirupsen/logrus"
)

var typeformAPIKey = os.Getenv("TYPEFORM_API_KEY")

var logLevel = flag.String("log-level", "info", "The level to log at. Can be one of debug, info, warn or error")

func main() {
	flag.Parse()
	logLevel, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetLevel(logLevel)
	var counter rsvp.Counter = typeform.NewCounter(typeformAPIKey, logger.WithField("component", "typeform.Counter"))
	attendeesChannel := make(chan int)
	go func() {
		attendees, err := counter.GetNumberOfAttendees()
		if err != nil {
			panic(err)
		}
		attendeesChannel <- attendees
	}()
	responses, err := counter.GetNumberOfResponses()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of responses: %d\nNumber of attendees: %d\n", responses, <-attendeesChannel)
}
