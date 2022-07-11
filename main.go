package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Animals interface {
	typeForAnimal
	nameForAnimal
	weightForAnimal
	foodForAnimal
	edibleForAnimal
	minWeightForAnimal
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

type edibleForAnimal interface {
	getEdibleAnimal() bool
}

type minWeightForAnimal interface {
	getMinWeight() int
}

var wrongAnimalType = errors.New("wrong type of animal")
var animalUnderweight = errors.New("underweight animal")
var animalNotEdible = errors.New("this animal is not edible")

func main() {
	allAnimalsInFarm := []Animals{
		dog{
			typeOfAnimal: "cat",
			nameOfAnimal: "Dog11",
			animalWeight: 15,
		},
		cat{
			typeOfAnimal: "cat",
			nameOfAnimal: "Cat11",
			animalWeight: 5,
		},
		cow{
			typeOfAnimal: "cow",
			nameOfAnimal: "Cow11",
			animalWeight: 199,
		},
		cow{
			typeOfAnimal: "cow",
			nameOfAnimal: "Cow22",
			animalWeight: 250,
		},
	}
	resultAllFoodForAllAnimals, err := countingAmountOfFeed(allAnimalsInFarm)
	if err != nil {
		err = fmt.Errorf("Warning!!! Program stopped: %w", err)
		fmt.Println(err)
	} else {
		fmt.Printf("All animals need %d kg of feed per month", resultAllFoodForAllAnimals)
	}
}

func countingAmountOfFeed(allAnimalsInFarm []Animals) (int, error) {
	var result int
	for _, v := range allAnimalsInFarm {
		err := validateParametresOfAnimal(v)
		if err != nil {
			if errors.Is(err, wrongAnimalType) {
				err = fmt.Errorf("bag in validation: %w", err)
				fmt.Println(err)
				continue

			} else if errors.Is(err, animalUnderweight) {
				err = fmt.Errorf("bag in validation: %w", err)
				return 0, err

			} else if errors.Is(err, animalNotEdible) {
				err = fmt.Errorf("bag in validation: %w", err)
				fmt.Println(err)
				continue
			}
		}
		result += v.getFeedingNorm()
		fmt.Printf(
			"Animal type: %s, name: %s, weight: %d kg, need feed per month: %d kg\n",
			v.getTypeOfAnimal(),
			v.getNameOfAnimal(),
			v.getAnimalWeight(),
			v.getFeedingNorm(),
		)
	}
	return result, nil
}

//ioa - instance of animal
func validateParametresOfAnimal(ioa Animals) error {
	err := checkTypeOfAnimal(ioa)
	if err != nil {
		err = fmt.Errorf("specified animal type is '%s', but the correct type is'%s': %w", ioa.getTypeOfAnimal(),
			reflect.TypeOf(ioa).Name(), err)
		err = fmt.Errorf("for %s: %w", reflect.TypeOf(ioa).Name(), err)
		return err
	}
	err = checkAnimalWeight(ioa)
	if err != nil {
		err = fmt.Errorf("min weight for this animal is %d kg,\n but weight of this animal is %d kg: %w", ioa.getMinWeight(),
			ioa.getAnimalWeight(), err)
		err = fmt.Errorf("for %s: %w", reflect.TypeOf(ioa).Name(), err)
		return err
	}
	err = checkAnimalforEdible(ioa)
	if err != nil {
		err = fmt.Errorf("the edibility status is incorrectly indicated\n for '%s': %w", reflect.TypeOf(ioa).Name(), err)
		err = fmt.Errorf("for %s: %w", reflect.TypeOf(ioa).Name(), err)
		return err
	}
	return nil
}

func checkTypeOfAnimal(ioa Animals) error {
	if reflect.TypeOf(ioa).Name() != ioa.getTypeOfAnimal() {
		return wrongAnimalType
	}
	return nil
}

func checkAnimalWeight(ioa Animals) error {
	if ioa.getAnimalWeight() <= ioa.getMinWeight() {
		return animalUnderweight
	}
	return nil
}

func checkAnimalforEdible(ioa Animals) error {
	if reflect.TypeOf(ioa).Name() != "cow" && ioa.getEdibleAnimal() == true {
		return animalNotEdible
	}
	return nil
}

//1!!
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

// 2!!
// створити PR до ферми
// додати функцію, яка буде приймати параметр інтерфейсного типу і виконує в собі мати 3 етапів валідації.
// кожен етап валідації - окрема функція що має повертати помилку
// валідація 1 - чи співпадає тип тварини з тим як вона називаєтсья
// валідація 2 - чи важить тварина меньше норми
// валідація 3 - чи є тварина їстівною
// якщо сталася помилка -  обгорнути її  на кожному етапі, додавши інформації як у прикладі.
// викликати валідацію перед тим як рахувати загальну кількість корму для ферми
// якщо сталася помилка - вивести її і припинити виконання програми
// вкладеність помилки має бути:
// помилка конкретної валідації
// що саме пійшло не так
// сказати що тварина не пройшла валідацію
// утонити для якої тварини провалена перевірка
// помилки можуть призводити до різних дій, на ваш розсуд можна обрати одну помилку яка зупеняє програму, а дві інші просто виводяться і ігнорується тільки конкретна тварина
// приклад:
// для собаки по імені "Боб": провалена валідація: боб насправді кіт: тип тварини має бути псом але їм не є
