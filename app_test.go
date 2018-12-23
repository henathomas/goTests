package spacecount

import (
	"fmt"
	"testing"
)

func TestValidateNumber(t *testing.T) {

	total := GetNumber()
	fmt.Println(total.Number)
	if total.Number != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total.Number, 6)
	}
	fmt.Println(total.Number)
}
