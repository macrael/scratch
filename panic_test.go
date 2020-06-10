package scratch

import "testing"
import "fmt"

func notImplemented(anarg string) bool {
	panic("Not Implemented")
}

func TestNotImplemented(t *testing.T) {
	fa := notImplemented("foo")

	t.Fatal(fa)
}

func TestFormatString(t *testing.T) {

	fmt.Printf(`{"message": %q}`, "the message")

	t.Fatal()
}
