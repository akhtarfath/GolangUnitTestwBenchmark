package BelajarGolangUnitTest

import (
	"github.com/akhtarfath/GolangUnitTestwBenchmark/helper"
	"testing"
)

func BenchmarkAllHelloWorldWithSub(b *testing.B) { // b is *testing.B
	for i := 0; i < b.N; i++ { // b.N is number of iteration
		b.Run("John Cena", func(b *testing.B) {
			helper.HelloWorld("John Cena")
		})
		b.Run("Fathan", func(b *testing.B) {
			helper.HelloWorldFathan()
		})
	}
}

// BenchmarkHelloWorld is a function to benchmark HelloWorld function
func BenchmarkHelloWorld(b *testing.B) { // b is *testing.B
	for i := 0; i < b.N; i++ { // b.N is number of iteration
		helper.HelloWorld("John")
	}
}

// BenchmarkHelloWorldFathan is a function to benchmark HelloWorldFathan function
func BenchmarkHelloWorldFathan(b *testing.B) { // b is *testing.B for benchmark
	for i := 0; i < b.N; i++ { // b.N is number of iteration for benchmark. this is the default value
		helper.HelloWorldFathan()
	}
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct { // struct is a data type that can contain multiple data type
		Name    string // Name is name of the benchmark
		Request string // Request is request parameter
	}{ // this is a table data
		{
			Name:    "HelloWorldJohn",
			Request: "John",
		},
		{
			Name:    "HelloWorldFathan",
			Request: "Fathan",
		},
		{
			Name:    "HelloWorldJohnCena",
			Request: "John Cena",
		},
		{
			Name:    "HelloWorldAulia",
			Request: "Aulia",
		},
	}

	for _, benchmark := range benchmarks { // loop for each benchmark
		b.Run(benchmark.Name, func(b *testing.B) { // benchmark.Name is name of the benchmark
			for i := 0; i < b.N; i++ { // b.N is number of iteration
				helper.HelloWorld(benchmark.Request) // benchmark.Request is request parameter
			}
		})
	}
}
