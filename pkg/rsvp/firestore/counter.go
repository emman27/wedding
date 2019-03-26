package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"flag"
	"fmt"
	"github.com/emman27/wedding/pkg/rsvp"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

var projectID = flag.String("project-id", "tingscoregoh", "Firebase project ID")

type Counter struct {
	client *firestore.Client
	log    logrus.FieldLogger
}

func NewCounter(opts ...Option) (*Counter, error) {
	conf := &firebase.Config{ProjectID: *projectID}
	app, err := firebase.NewApp(context.TODO(), conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(context.TODO())
	if err != nil {
		return nil, err
	}
	c := &Counter{client: client, log: logrus.New().WithField("component", "firestore.Counter")}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

var _ rsvp.Counter = (*Counter)(nil)

func (c *Counter) GetNumberOfResponses() (int, error) {
	iter := c.client.Collection("rsvps").Where("numberOfAttendees", ">=", 0).Documents(context.TODO())
	i := 0
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		i++
	}
	return i, nil
}

func (c *Counter) GetNumberOfAttendees() (int, error) {
	iter := c.client.Collection("rsvps").Where("numberOfAttendees", ">=", 0).Documents(context.TODO())
	i := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		val, err := doc.DataAt("numberOfAttendees")
		if err != nil {
			return 0, err
		}
		if count, ok := val.(int64); ok {
			i += int(count)
		} else {
			return 0, fmt.Errorf("Couldn't convert %v to integer", val)
		}
	}
	return i, nil
}

type Option func(*Counter) error

func WithLogger(log logrus.FieldLogger) Option {
	return func(c *Counter) error {
		c.log = log
		return nil
	}
}
