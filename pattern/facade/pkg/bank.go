package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("*[Банк] Получение остатка по карте %s\n", cardNumber)

	time.Sleep(time.Millisecond * 300)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue

		} else {
			if card.Balance <= 0 {
				return errors.New("*[Банк] Недостаточно средств")

			}
		}
	}

	fmt.Println("*[Банк] Остаток положительный")

	return nil
}
