package decorator

import (
	"testing"
)

func TestDecrator(t *testing.T) {
	americano := &Coffee{Type:Mocha, Name:"americano coffee"}
	Print(americano)
	milk := Milk{Deco:americano}
	print(milk.Coast())

	cho := Chocloate{&milk}
	print(cho.Coast())

}