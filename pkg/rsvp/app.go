package rsvp

import (
	"fmt"
	"log"
	"os"

	"github.com/emman27/wedding/pkg/notifications"
	"github.com/sirupsen/logrus"
)

// NewApp starts up a new app
func NewApp(counter Counter, notificationSender notifications.Sender, options ...Option) *App {
	a := &App{counter: counter, notificationSender: notificationSender, logger: log.New(os.Stdout, "rsvp.App", log.LstdFlags)}
	for _, opt := range options {
		opt(a)
	}
	return a
}

// Option applies some additional configuration to the application
type Option func(*App)

// SetLogger option for an application
func SetLogger(logger logrus.StdLogger) Option {
	return func(a *App) {
		a.logger = logger
	}
}

// App tracks the number of RSVPs and sends some notifications
type App struct {
	counter            Counter
	notificationSender notifications.Sender
	logger             logrus.StdLogger
}

// Run the cycle of fetching and dispatching notifications
func (a *App) Run() error {
	attendeesChannel := make(chan int)
	errChannel := make(chan error)
	go func() {
		attendees, err := a.counter.GetNumberOfAttendees()
		if err != nil {
			errChannel <- err
		}
		attendeesChannel <- attendees
	}()
	responses, err := a.counter.GetNumberOfResponses()
	if err != nil {
		return err
	}
	var attendees int
	select {
	case result := <-attendeesChannel:
		attendees = result
	case err = <-errChannel:
		return err
	}
	a.notificationSender.SendMessage(fmt.Sprintf("Number of responses: %d\nNumber of attendees: %d\n", responses, attendees))
	return nil
}
