package eda

type CommandPublisher interface {
	Register(commandID string, handler CommandHandler)
	Publish(cmd Command)
}

type EventPublisher interface {
	Register(eventID string, handler EventHandler)
	Publish(evt Event)
}
