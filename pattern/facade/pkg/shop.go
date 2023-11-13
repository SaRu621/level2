package pkg

import (
	"errors"
	"fmt"
	"time"
)

// shop и есть фасад, на нем все завязано

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error { // sell -- функция-фасад над сложной бизнес-логикой объектов card, user, bank и тп.
	fmt.Println("*[Магазин] Запрос к пользователю для получения остатка на карте")

	time.Sleep(time.Millisecond * 500)

	err := user.Card.CheckBalance()

	if err != nil {
		return err
	}

	fmt.Printf("*[Магазин] Проверка - может ли [%s] купить товар \n", user.Name)

	time.Sleep(time.Millisecond * 500)

	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > user.GetBalance() {
			return errors.New("*[Магазин] Недостаточно средств для покупки товара")
		}

		fmt.Printf("*[Магазин] Товар [%s] - куплен \n", prod.Name)
	}

	return nil
}
