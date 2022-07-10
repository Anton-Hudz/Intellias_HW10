package main

import "fmt"

type Animals interface {
	typeForAnimal
	nameForAnimal
	weightForAnimal
	foodForAnimal
}

type typeForAnimal interface {
	getTypeOfAnimal() string
}

type nameForAnimal interface {
	getNameOfAnimal() string
}
type weightForAnimal interface {
	getAnimalWeight() int
}

type foodForAnimal interface {
	getFeedingNorm() int
}

func main() {
	allAnimalsInFarm := []Animals{
		dog{
			typeOfAnimal: "Dog",
			nameOfAnimal: "Dog11",
			animalWeight: 15,
		},
		cat{
			typeOfAnimal: "Cat",
			nameOfAnimal: "Cat11",
			animalWeight: 5,
		},
		cow{
			typeOfAnimal: "Cow",
			nameOfAnimal: "Cow11",
			animalWeight: 300,
		},
		cow{
			typeOfAnimal: "Cow",
			nameOfAnimal: "Cow22",
			animalWeight: 250,
		},
	}
	resultAllFoodForAllAnimals := countingAmountOfFeed(allAnimalsInFarm)
	fmt.Printf("All animals need %d kg of feed per month", resultAllFoodForAllAnimals)

}

func countingAmountOfFeed(allAnimalsInFarm []Animals) int {
	var result int
	for _, v := range allAnimalsInFarm {
		result += v.getFeedingNorm()
		fmt.Printf(
			"Animal type: %s, name: %s, weight: %d kg, need feed per month: %d kg\n",
			v.getTypeOfAnimal(),
			v.getNameOfAnimal(),
			v.getAnimalWeight(),
			v.getFeedingNorm(),
		)
	}
	return result

}

// Домашнє завдання по інтерфейсах:
// кожна тварина в залежності від свого типу споживає різку кількість
// кілограмів їжі на місяць
// -- собака - їсть 10кг/мес корму на кожні 5кг власної ваги
// -- кішка - 7кг/мес на кожен кілограм власної ваги
// -- корова - 25кг/мес на кожен кілограм власної ваги
// на фермі може бути різна кількість собак, кішок та корів
// кажен тип тварини має сам розраховувати для себе вагу корму
// написати динамічну функцію,
// 1) яка буде виводити в консоль для кожної тварини на фермі її назву,
// 2) вагу, та скільки їжі треба для того щоб її прогодувати
// 3) ця функція також моє повертати сумму кг корму на всю ферму для
// подальшого виводу у консоль
