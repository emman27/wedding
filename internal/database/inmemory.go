package database

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/emman27/wedding/pkg/rsvp-saver"
)

type InMemory struct {
}

func (db *InMemory) SaveRSVP(saver.RSVP) error {
	return errors.New("not implemented")
}
