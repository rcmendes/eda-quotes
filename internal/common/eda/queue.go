package eda

type CommandQueue interface {
	Register(commandID string, handler CommandHandler)
	Publish(cmd Command, result chan interface{}, err chan error)
	PublishAndForget(cmd Command)
}

type EventQueue interface {
	Register(eventID string, handler EventHandler)
	Publish(evt Event)
}
