package iteration

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q want %q", repeated, expected)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("b", 7)
	fmt.Println(repeatedString)

	// Output: bbbbbbb
}