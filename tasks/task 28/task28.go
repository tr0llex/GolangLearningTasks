package main

import "fmt"

type client struct{}

type card interface {
	pay()
}

type sber struct{}

type troyka struct{}

type currencyAdapter struct {
	troykaCard *troyka
}

func (c *client) payByCard(visa card) {
	fmt.Println("Клиент рассчитывается")
	visa.pay()
}

func (s *sber) pay() {
	fmt.Println("Оплачиваем рублями")
}

func (troyka *troyka) pay() {
	fmt.Println("Оплачиваем долларами")
}

func (a *currencyAdapter) pay() {
	fmt.Println("Конвертация валюты из рублей в доллары")
	a.troykaCard.pay()
}

func main() {
	client := &client{}
	sber := &sber{}

	client.payByCard(sber)

	troykaCard := &troyka{}
	troykaCardAdapter := &currencyAdapter{
		troykaCard: troykaCard,
	}

	client.payByCard(troykaCardAdapter)
}
