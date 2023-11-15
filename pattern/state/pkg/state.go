package pkg

type State interface {
	AddItem(count int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}
