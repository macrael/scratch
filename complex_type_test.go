package scratch

import "testing"

type foo int

func (f foo) Hi() int {
	return 3
}

type foos []foo

// func (fs []foo) Hi() int { //This fails to compile:    invalid receiver type []foo ([]foo is not a defined type)
func (fs foos) Hi() int {
	return 33
}

func TestComplexTypeMethod(t *testing.T) {

	t.Fatal("NO")
}
