package eda

type Command interface {
	CommandID() string
}

type CommandHandler interface {
	Handle(command Command) (result interface{}, err error)
}
