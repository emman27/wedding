// Package rsvp deals with how our RSVP rates are going
package rsvp

// Counter exposes a mechanism to track the number of people who have responded to the invites
type Counter interface {
	GetNumberOfResponses() (int, error)
	GetNumberOfAttendees() (int, error)
}
