package golang_generics

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleType(t *testing.T) {
	MultipleParameter("Masred", 100)
	MultipleParameter(100, "Masred")
}
