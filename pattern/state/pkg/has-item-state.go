package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)

		return fmt.Errorf("No item present")
	}

	fmt.Printf("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)

	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added", count)
	i.vendingMachine.incrementItemCount(count)

	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
