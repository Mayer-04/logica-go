package main

type Message struct {
	title, message, messageType string
	account                     string
	accountList                 []string
	token                       string
	tokenList                   []string
}

type MessageOption func(*Message)

func NewMessage(title, message, messageType string, opts ...MessageOption) *Message {
	msg := &Message{
		title:       title,
		message:     message,
		messageType: messageType,
	}

	for _, opt := range opts {
		opt(msg) // aplicar cada opción
	}

	return msg
}

func WithAccount(account string) MessageOption {
	return func(msg *Message) {
		msg.account = account
	}
}

func main() {

}
