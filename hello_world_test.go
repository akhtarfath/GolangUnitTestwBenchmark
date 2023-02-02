package BelajarGolangUnitTest

import (
	"fmt"
	"github.com/akhtarfath/GolangUnitTestwBenchmark/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// in unit test, don't using panic, use t.Error / t.Fail / t.FailNow / t.Fatal
func TestHelloWorld(t *testing.T) { // Test function must start with Test
	result := helper.HelloWorld("John")
	if result != "Hello John" { // If result is not equal to "Hello John"
		t.Error("Result must be Hello John") // t.Error will show error message
	}
}

func TestHelloWorldFathanWithRequire(t *testing.T) {
	result := helper.HelloWorldFathan()
	// t is *testing.T
	// "Hello Fathan" is expected result
	// result is actual result
	// "Result must be Hello Fathan" is error message
	// require will stop the test if the result is not equal to expected result
	require.Equal(t, "Hello Fathan", result, "Result must be Hello Fathan")
	fmt.Println("TestHelloWorldFathan with require.Equal Done")
}

func TestHelloWorldFathanWithAssert(t *testing.T) {
	result := helper.HelloWorldFathan()
	// t is *testing.T
	// "Hello Fathan" is expected result
	// result is actual result
	// "Result must be Hello Fathan" is error message
	// assert will continue the test if the result is not equal to expected result
	assert.Equal(t, "Hello Fathan", result, "Result must be Hello Fathan")
	fmt.Println("TestHelloWorldFathan with assert.Equal Done")
}

// t.Skip will skip the test
func TestHelloWorldFathanWithSkipTest(t *testing.T) {
	if runtime.GOOS == "darwin" { // Skip test if OS is Mac OS
		t.Skip("Can not run on Mac OS")
	} else if runtime.GOOS == "linux" { // Skip test if OS is Linux
		t.Skip("Can not run on Linux")
	}

	result := helper.HelloWorldFathan()
	require.Equal(t, "Hello Fathan", result, "Result must be Hello Fathan")
}

func TestSubTest(t *testing.T) {
	// John is a sub test of TestSubTest
	t.Run("John", func(t *testing.T) { // t.Run will run sub test
		result := helper.HelloWorld("John")
		assert.Equal(t, "Hello John", result, "Result must be Hello John")
	})
	t.Run("John2", TestHelloWorld) // TestHelloWorld is a function that already exist
	// Fathan is a sub test of TestSubTest
	t.Run("Fathan", func(t *testing.T) { // t.Run will run sub test
		result := helper.HelloWorldFathan()
		require.Equal(t, "Hello Fathan", result, "Result must be Hello Fathan")
	})
	t.Run("Fathan2", TestHelloWorldFathanWithRequire) // TestHelloWorldFathanWithRequire is a function that already exist
}

// Table Test is a test that use table data
func TestHelloWorldTable(t *testing.T) {
	tests := []struct { // struct for table test data
		name, request, expected, message string // name of the test, request parameter, expected result
	}{
		{
			name:     "HelloWorld(John)", // name of the test
			request:  "John",             // request parameter
			expected: "Hello John",       // expected result
			message:  "Result must be Hello John",
		},
		{
			name:     "HelloWorld(Fathan)", // name of the test
			request:  "Fathan",             // request parameter
			expected: "Hello Fathan",       // expected result
			message:  "Result must be Hello Fathan",
		},
	}

	// _ is index, test is data, test is a struct
	for _, test := range tests { // loop for each test data
		t.Run(test.name, func(t *testing.T) { // test.name is name of the test
			result := helper.HelloWorld(test.request)                               // test.request is request parameter
			assert.Equal(t, test.expected, result, "Result must be "+test.expected) // test.expected is expected result
			assert.Equal(t, test.expected, result, test.message)                    // test.expected is expected result
		})
	}
}

// test Main will run before and after unit test in this package
func TestMain(m *testing.M) { // *testing.M is a type of testing package
	// before unit test
	fmt.Println("Before Unit Test")
	// run all unit test
	m.Run() // Run all unit test just once
	// after unit test
	fmt.Println("After Unit Test")
}
