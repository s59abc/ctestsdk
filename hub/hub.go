package hub

import "fmt"

// Represent Command that should be executed
type Command interface {
	Execute()
}

func Start(commandChan chan Command) {
	for {
		cmd := <-commandChan
		cmd.Execute()
	}
}

type Connection struct {
	ID string
}

type ConnectCmd struct {
	connection Connection
}

func (cc ConnectCmd) Execute() {
	fmt.Println(cc.connection.ID)
}

type DisconnectCmd struct {
	connection Connection
}

func (cc DisconnectCmd) Execute() {
	fmt.Println(cc.connection.ID)
}
