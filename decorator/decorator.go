package decorator

import (
	"fmt"
)

type DecType int

const (
	_ DecType = iota
	Americano
	Mocha
)


type Decorator interface {
	Coast() float64
}

type Coffee struct {
	Type DecType
	Name string
}

func Print(dec Decorator) {
	fmt.Println(dec.Coast())
}

func (self *Coffee) Coast() float64 {
	switch self.Type {
	case Americano:
		return 13.0
	case Mocha:
		return 15.0
	default:
		return 10.0

	}
}


type Milk struct {
	Deco Decorator
}

func (self *Milk) Coast() float64 {
	return 5.0 + self.Deco.Coast()
}

type Chocloate struct {
	Deco Decorator
}

func (self *Chocloate) Coast() float64 {
	return 3.0 + self.Deco.Coast()
}





