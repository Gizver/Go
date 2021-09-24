/*
Построить таблицу с информацией по билетам для путешествия на Марс.
Для каждого билета случайным образом выбирается космическая станция: Space Adventures, SpaceX или Virgin Galactic.
Обеспечить ввод расстояния между Марсом и Землей (в км).
Скорость космического корабля будет выбрана случайным образом из диапазона от 16 до 30 км/c. Это определит продолжительность поездки на Марс,
а также цену билета. Более быстрые корабли дороже. Цены на билеты варьируются от $36 до $50 миллионов. Цена для поездки туда-обратно удваивается.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const secondsPerDay = 24 * 60 * 60

	fmt.Println("Введите расстояние до Марса, на момент вылета (км):")

	var distance int
	fmt.Scanln(&distance)

	fmt.Println("Spaceline        Days   Trip type     Price, m.")

	for count := 0; count < 10; count++ {

		speed := rand.Intn(15) + 16
		days := distance / speed / secondsPerDay
		price := 36 + (speed - 16)
		trip := "One-way"
		spaceline := ""
		if tripRout := rand.Intn(2); tripRout == 0 {
			trip = "Round-Trip"
			price = 2 * price
		}

		switch spacelineId := rand.Intn(3); spacelineId {
		case 0:
			spaceline = "SpaceX"
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "Virgin Galactic"
		}
		fmt.Printf("%-16v %-6v %-14v %6v $\n", spaceline, days, trip, price)
	}
}