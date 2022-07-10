package main

const (
	feedingNormDog = 10 / 5
	feedingNormCat = 7
	feedingNormCow = 25
)

type dog struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
}

func (d dog) getTypeOfAnimal() string {
	return d.typeOfAnimal
}

func (d dog) getNameOfAnimal() string {
	return d.nameOfAnimal
}

func (d dog) getAnimalWeight() int {
	return d.animalWeight
}

func (d dog) getFeedingNorm() int {
	return d.feedingNorm
}

type cat struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
}

func (c cat) getTypeOfAnimal() string {
	return c.typeOfAnimal
}

func (c cat) getNameOfAnimal() string {
	return c.nameOfAnimal
}

func (c cat) getAnimalWeight() int {
	return c.animalWeight
}

func (c cat) getFeedingNorm() int {
	return c.feedingNorm
}

type cow struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
}

func (c cow) getTypeOfAnimal() string {
	return c.typeOfAnimal
}

func (c cow) getNameOfAnimal() string {
	return c.nameOfAnimal
}

func (c cow) getAnimalWeight() int {
	return c.animalWeight
}

func (c cow) getFeedingNorm() int {
	return c.feedingNorm
}
