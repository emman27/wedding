package telegram

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/emman27/wedding/pkg/notifications"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// My personal chat details, set as default for convenience
const emmanChatID = 207051227

const messagePrefix = `Hello! I'm your friendly chatbot to help with keeping track of the number of invites to your wedding!

Here are your summary statistics so far!
%s`

var token = os.Getenv("TELEGRAM_BOT_TOKEN")

var recipientChatID = flag.Int64("telegram-chat-id", emmanChatID, "Unique chat ID to send messages to")

// Telegram constants
const (
	BaseURL = "https://api.telegram.org/bot%s"
)

// New telegram notification client
func New(opts ...Option) (*Sender, error) {
	s := &Sender{
		logger: logrus.New(),
	}
	for _, opt := range opts {
		opt(s)
	}
	s.logger = s.logger.WithField("component", "telegram.Sender")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		s.logger.Errorf("Failed to initialize telegram client with token %s", token)
		return nil, err
	}
	s.bot = bot
	self, err := s.bot.GetMe()
	if err != nil {
		s.logger.Errorf("Unable to get bot details")
		return nil, err
	}
	s.logger = s.logger.WithField("bot_name", self.UserName)
	s.logger.Infof("Initialized telegram client with token %s", strings.Repeat("*", len(token)))
	return s, nil
}

// Sender sends telegram messages
type Sender struct {
	bot    *tgbotapi.BotAPI
	logger logrus.FieldLogger
}

// SendMessage to the configured recipient
func (s Sender) SendMessage(msg string) error {
	m := tgbotapi.NewMessage(*recipientChatID, fmt.Sprintf(messagePrefix, msg))
	_, err := s.bot.Send(m)
	return err
}

var _ notifications.Sender = (*Sender)(nil)

// Option to configure the telegram notification sender
type Option func(*Sender)

// SetLogger to use for this notification system
func SetLogger(logger logrus.FieldLogger) Option {
	return func(s *Sender) {
		s.logger = logger
	}
}
