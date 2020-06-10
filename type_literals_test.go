package scratch

import (
	"fmt"
	"testing"
)

type OneType struct {
	name string
}

func (t OneType) PrintName() {
	fmt.Println("NAME: ", t.name)
}

type TwoType struct {
	firstName  string
	secondName string
}

func (t TwoType) PrintName() {
	fmt.Println("NAME: ", t.firstName, t.secondName)
}

type nameable interface {
	PrintName()
}

func TestLiteralList(t *testing.T) {

	allTypes := []nameable{
		OneType{}, TwoType{},
	}

	for _, aType := range allTypes {
		fmt.Println("Instance: ", aType)
		aType.PrintName()
	}

	t.Fatal("Did it run?")
}
