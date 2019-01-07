package lib

// Command can be executed by someone
type Command interface {
	Execute() error
}
