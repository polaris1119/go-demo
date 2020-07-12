package singleton_test

import (
	"testing"

	"github.com/polaris1119/go-demo/singleton"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleton.New()
	}
}

func BenchmarkNew2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleton.New2()
	}
}

func BenchmarkNew3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleton.New3()
	}
}
