package cash

import (
	"fmt"
	"testing"
)

func TestVisitHappened(t *testing.T) {
	err := VisitHappened()
	fmt.Println(err)
}

func TestVisitGetNumber(t *testing.T) {
	number, err := VisitGetNumber()
	fmt.Println(number, err)
	t.Error(err)
	t.Fail()
}
