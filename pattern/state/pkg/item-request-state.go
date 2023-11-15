package pkg

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item dispence in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted monet is less. Please insert %d", i.vendingMachine.itemPrice)
	}

	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)

	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
