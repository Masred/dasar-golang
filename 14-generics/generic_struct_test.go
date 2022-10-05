package golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Masred",
		Second: "Ganteng",
	}
	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Masred",
		Second: "Ganteng",
	}

	assert.Equal(t, "Masred", data.ChangeFirst("Masred"))
	assert.Equal(t, "Hello Masred", data.SayHello("Masred"))
}
