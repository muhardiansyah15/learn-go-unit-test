package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hardi",
			request:  "Hardi",
			expected: "Hello Hardi",
		},
		{
			name:     "Hardian",
			request:  "Hardian",
			expected: "Hello Hardi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Hardi", func(t *testing.T) {
		result := HelloWorld("Hardi")
		require.Equal(t, "Hello Hardi", result, "Result must be Hello Hardi")
	})

	t.Run("Hardian", func(t *testing.T) {
		result := HelloWorld("Hardi")
		require.Equal(t, "Hello Hardian", result, "Result must be Hello Hardi")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// After
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can  not run")
	}

	result := HelloWorld("Hardi")
	require.Equal(t, "Hello Hardian", result, "Result must be Hello Hardi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Hardi")
	require.Equal(t, "Hello Hardian", result, "Result must be Hello Hardi")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Hardi")
	assert.Equal(t, "Hello Hardian", result, "Result must be Hello Hardi")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hardi")

	if result != "Hello Hardia" {
		t.Fatal("Result mus be Hello Hardi")
	}

	// if result != "Hello Hardi" {
	// 	t.Fail()
	// }

	fmt.Print("TestHelloWorld Done")

	// if result != "Hello Hardi" {
	// 	panic("I dont know")
	// }
}

func TestHelloWorldSecond(t *testing.T) {
	result := HelloWorld("Febri Zummiati")

	if result != "Hello Febri Zummiati" {
		t.Fail()
	}

	// if result != "Hello Hardi" {
	// 	panic("I dont know")
	// }
}
