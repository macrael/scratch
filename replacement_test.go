package scratch

import "testing"

import "strings"

func TestSometing(t *testing.T) {

	template := `
	{
		"here": "$here",
		"some": $some
	}
	`

	expected := `
	{
		"here": "foo",
		"some": 23
	}
	`

	replacer := strings.NewReplacer(
		"$here", "foo",
		"$some", "23",
	)

	actual := replacer.Replace(template)

	if expected != actual {
		t.Fatal("DIDN'tLREpP", actual)
	}

	t.Fatal("FO")
}
