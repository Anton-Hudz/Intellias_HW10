package main

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
	feedingNormDog := 10 / 5
	d.feedingNorm = feedingNormDog * d.animalWeight
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
	feedingNormCat := 7
	c.feedingNorm = feedingNormCat * c.animalWeight
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
	feedingNormCow := 25
	c.feedingNorm = feedingNormCow * c.animalWeight
	return c.feedingNorm
}
