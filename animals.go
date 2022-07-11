package main

var (
	dogMinWeight   int = 5
	catMinWeight   int = 2
	cowMinWeight   int = 200
	feedingNormDog int = 10 / 5
	feedingNormCat int = 7
	feedingNormCow int = 25
)

type dog struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
	edibleAnimal bool
	minWeight    int
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
	d.feedingNorm = feedingNormDog * d.animalWeight
	return d.feedingNorm
}

func (d dog) getEdibleAnimal() bool {
	return false
}

func (d dog) getMinWeight() int {
	d.minWeight = dogMinWeight
	return d.minWeight
}

type cat struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
	edibleAnimal bool
	minWeight    int
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
	c.feedingNorm = feedingNormCat * c.animalWeight
	return c.feedingNorm
}

func (c cat) getEdibleAnimal() bool {
	return false
}

func (c cat) getMinWeight() int {
	c.minWeight = catMinWeight
	return c.minWeight
}

type cow struct {
	typeOfAnimal string
	nameOfAnimal string
	animalWeight int
	feedingNorm  int
	edibleAnimal bool
	minWeight    int
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
	c.feedingNorm = feedingNormCow * c.animalWeight
	return c.feedingNorm
}

func (c cow) getEdibleAnimal() bool {
	return true
}

func (c cow) getMinWeight() int {
	c.minWeight = cowMinWeight
	return c.minWeight
}
