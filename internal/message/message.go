package message

import "time"

type Message struct {
	From      string
	Content   string
	CreatedAt time.Time
}

func New(from string, content string) Message {
	return Message{
		From:      from,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (msg Message) Format() string {
	return "[" + msg.From + "] : " + msg.Content
}
