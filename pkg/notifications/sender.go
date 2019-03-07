package notifications

// Sender sends notifications
type Sender interface {
	SendMessage(string) error
}
