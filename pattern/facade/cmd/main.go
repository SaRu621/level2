package main

import (
	"facade/pkg"
	"fmt"
)

var (
	// банк
	bank = pkg.Bank{
		Name:  "Bank",
		Cards: []pkg.Card{},
	}

	//карточки

	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}

	// пользователи

	user1 = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
	}

	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}

	// продукт

	prod = pkg.Product{
		Name:  "Cheese",
		Price: 150,
	}

	// магазин

	shop = pkg.Shop{
		Name: "Пятерочка",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {

	fmt.Println("*[Банк] Выпуск карт")

	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)

	err := shop.Sell(user1, prod.Name)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("[%s]", user2.Name)

	err = shop.Sell(user2, prod.Name)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

}
