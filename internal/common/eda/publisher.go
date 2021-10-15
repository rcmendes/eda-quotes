package eda

type CommandPublisher interface {
	Register(commandID string, handler CommandHandler)
	Notify(cmd Command)
}

type EventPublisher interface {
	Register(eventID string, handler EventHandler)
	Notify(evt Event)
}
