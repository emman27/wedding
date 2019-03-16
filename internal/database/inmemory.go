package database

import (
	"github.com/emman27/wedding/pkg/rsvp-saver"
	"github.com/sirupsen/logrus"
	"time"
)

func NewInMemoryDB() saver.Database {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	db := &InMemory{
		RSVPs:  []*RSVP{},
		Logger: logger.WithField("component", "database.InMemory"),
	}
	return db
}

type RSVP struct {
	name              string
	numberOfAttendees int
	contact           string
	timestamp         time.Time
}

func (r *RSVP) String() string {
	return r.name
}

type InMemory struct {
	RSVPs  []*RSVP
	Logger logrus.FieldLogger
}

func (db *InMemory) SaveRSVP(r saver.RSVP) error {
	db.RSVPs = append(db.RSVPs, &RSVP{
		name:              r.Name(),
		numberOfAttendees: r.NumberOfPeople(),
		contact:           r.Contact(),
		timestamp:         r.Date(),
	})
	db.Logger.Debugf("Current RSVPs: %v", db.RSVPs)
	return nil
}
