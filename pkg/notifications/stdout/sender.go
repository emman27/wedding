package stdout

import (
	"fmt"

	"github.com/emman27/wedding/pkg/notifications"
)

// Sender sends messages to stdout
type Sender struct{}

var _ notifications.Sender = (*Sender)(nil)

// NewSender initializes a new Sender
func NewSender() *Sender {
	return new(Sender)
}

func (s *Sender) SendMessage(msg string) error {
	fmt.Println(msg)
	return nil
}
