package pkg

import "fmt"

type PC struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPC() Computer {
	return PC{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc PC) GetType() string {
	return pc.Type
}

func (pc PC) PrintDetails() {
	fmt.Printf("Type:[%s], Core:[%d], Mem:[%d], Monitor:[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Memory)
}
