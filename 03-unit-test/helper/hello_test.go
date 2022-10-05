package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ujang",
			request: "Ujang",
		},
		{
			name:    "Masred",
			request: "Masred",
		},
		{
			name:    "Dedi",
			request: "Dedi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Masred", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Hello Masred")
		}
	})

	b.Run("Ujang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Hello Ujang")
		}
	})
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Hello World!")
	}
}

func TestHelloTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ujang",
			request:  "Ujang",
			expected: "Hello Ujang",
		},
		{
			name:     "Masred",
			request:  "Masred",
			expected: "Hello Masred",
		},
		{
			name:     "Dedi",
			request:  "Dedi",
			expected: "Hello Dedi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Masred", func(t *testing.T) {
		result := Hello("Masred")
		require.Equal(t, "Hello Masred", result, "Result must be 'Hello Masred'")
	})
	t.Run("Ujang", func(t *testing.T) {
		result := Hello("Ujang")
		require.Equal(t, "Hello Ujang", result, "Result must be 'Hello Ujang'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "Windows" {
		t.Skip("Can not run on Windows")
	}
}

func TestHelloAssert(t *testing.T) {
	result := Hello("Masred")
	assert.Equal(t, "Hello Masred", result, "Result must be 'Hello Masred'")
	fmt.Println("TestHello with Assert Done!")
}

func TestHello(t *testing.T) {
	result := Hello("World")

	if result != "Hello World" {
		// error
		t.Error("Result must be 'Hello World'")
	}

	fmt.Println("TestHello Done!")
}

func TestHelloMasred(t *testing.T) {
	result := Hello("Masred")

	if result != "Hello Masred" {
		// error
		t.Fatal("Result must be 'Hello Masred'")
	}

	fmt.Println("TestHelloMasred Done!")
}
