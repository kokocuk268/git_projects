package main

import "fmt"

func main() {
	var (
		firstPlace  = "-"
		secondPlace = "-"
		thirdPlace  = "-"
		fourthPlace = "-"
		fifthPlace  = "-"
	)
	for {
		var inputStr string
		var inputInt int
		fmt.Scanln(&inputStr, &inputInt)

		//! ----- ПРОВЕРКА НЕКОРРЕКТНОГО ВВОДА Пустота -----

		if (inputInt == 0 && inputStr == "") {
			continue
		}

		//! ----- КОМАНДЫ -----

		if inputStr == "очередь" {
			fmt.Println("1. " + firstPlace)
			fmt.Println("2. " + secondPlace)
			fmt.Println("3. " + thirdPlace)
			fmt.Println("4. " + fourthPlace)
			fmt.Println("5. " + fifthPlace)
			continue
		}
		if inputStr == "конец" {
			fmt.Println("1. " + firstPlace)
			fmt.Println("2. " + secondPlace)
			fmt.Println("3. " + thirdPlace)
			fmt.Println("4. " + fourthPlace)
			fmt.Println("5. " + fifthPlace)
			break
		}

		if inputStr == "количество" {
			var zanato int
			var free int

			if firstPlace != "-" {
				zanato++
			} else {
				free++
			}
			if secondPlace != "-" {
				zanato++
			} else {
				free++
			}
			if thirdPlace != "-" {
				zanato++
			} else {
				free++
			}
			if fourthPlace != "-" {
				zanato++
			} else {
				free++
			}
			if fifthPlace != "-" {
				zanato++
			} else {
				free++
			}
			fmt.Printf("Осталось свободных мест: %d\n", free)
			fmt.Printf("Всего человек в очереди: %d\n", zanato)

			continue
		}

		//! ----- ПРОВЕРКА НЕКОРРЕКТНОГО ВВОДА -----

		if (inputInt >= 6 || inputInt <= 0) {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", inputInt)
			continue
		}

		//! ----- ПРОВЕРКА ПЕРЕПОЛНЕНИЯ -----

		if firstPlace != "-" && secondPlace != "-" && thirdPlace != "-" && fourthPlace != "-" && fifthPlace != "-" {

			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", inputInt)
			continue
		}

		//! ----- ЗАПИСЬ -----

		switch inputInt {
		case 1:
			if firstPlace != "-" {
				fmt.Println("Запись на место номер 1 невозможна: место уже занято")
				continue
			}
			firstPlace = inputStr
		case 2:
			if secondPlace != "-" {
				fmt.Println("Запись на место номер 2 невозможна: место уже занято")
				continue
			}
			secondPlace = inputStr
		case 3:
			if thirdPlace != "-" {
				fmt.Println("Запись на место номер 3 невозможна: место уже занято")
				continue
			}
			thirdPlace = inputStr
		case 4:
			if fourthPlace != "-" {
				fmt.Println("Запись на место номер 4 невозможна: место уже занято")
				continue
			}
			fourthPlace = inputStr
		case 5:
			if fifthPlace != "-" {
				fmt.Println("Запись на место номер 5 невозможна: место уже занято")
				continue
			}
			fifthPlace = inputStr
		}
	}
}
