package saver

import (
	"time"
)

type Saver struct {
	database Database
}

func NewSaver(db Database) *Saver {
	return &Saver{database: db}
}

func (s *Saver) Save(rsvp RSVP) error {
	return s.database.SaveRSVP(rsvp)
}

type Database interface {
	SaveRSVP(RSVP) error
}

type RSVP interface {
	Name() string
	NumberOfPeople() int
	Date() time.Time
	Contact() string
}
