package pkg

import "fmt"

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case NotebookType:
		return NewNotebook()
	case PersonalComputerType:
		return NewPC()
	case ServerType:
		return NewServer()
	default:
		fmt.Printf("Несуществующий тип объекта: %s\n", typeName)
		return nil
	}
}
