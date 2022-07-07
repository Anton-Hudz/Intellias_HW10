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

func (d dog) gettypeOfAnimal() string {
	return d.typeOfAnimal
}

func (d dog) getnameOfAnimal() string {
	return d.nameOfAnimal
}

func (d dog) getanimalWeight() int {
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

func (c cat) gettypeOfAnimal() string {
	return c.typeOfAnimal
}

func (c cat) getnameOfAnimal() string {
	return c.nameOfAnimal
}

func (c cat) getanimalWeight() int {
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

func (c cow) gettypeOfAnimal() string {
	return c.typeOfAnimal
}

func (c cow) getnameOfAnimal() string {
	return c.nameOfAnimal
}

func (c cow) getanimalWeight() int {
	return c.animalWeight
}

func (c cow) getFeedingNorm() int {
	return c.feedingNorm
}
