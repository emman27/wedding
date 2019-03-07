package main

import (
	"flag"
	"os"

	"github.com/emman27/wedding/pkg/notifications/stdout"
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
	var app = rsvp.NewApp(counter, stdout.NewSender(), rsvp.SetLogger(logger.WithField("component", "rsvp.App")))
	if err = app.Run(); err != nil {
		panic(err)
	}
}
