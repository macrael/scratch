package scratch

import (
	"fmt"
	"testing"
)

type namedStruct struct {
	me  string
	you int
}

func TestDifferentLiterals(t *testing.T) {

	// So, it's a little surprising to me that you can write a table test literal without naming each struct

	tests := []struct {
		myField   string
		yourField int
	}{
		{"foo", 3},
		{"bar", 7},
	}

	fmt.Println("tests", tests)

	// 1. can you do that with named structs?

	namedTests := []namedStruct{
		{"foo", 3},
		{"bar", 7},
	}

	fmt.Println("named", namedTests)

	// 2. obviously, this will work
	explicitlyNamedTests := []namedStruct{
		namedStruct{"foo", 3},
		namedStruct{"bar", 7},
	}

	fmt.Println("explicitly named", explicitlyNamedTests)

	t.Fatal()
}
