package phonenumber_test

import (
	"testing"

	"github.com/foril/go-libphonenumber"
)

func BenchmarkIsPossibleNumberPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.IsPossibleNumber("+358401231234", "FI")
	}
}

func BenchmarkIsPossibleNumberPositiveHardCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.IsPossibleNumber("040-123 4321", "FI")
	}
}

func BenchmarkIsPossibleNumberNegativeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.IsPossibleNumber("asbdgjwoprg", "FI")
	}
}

func BenchmarkParseEasyPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.Parse("+358401231234", "FI")
	}
}

func BenchmarkParseHardPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.Parse("040-123 4321", "FI")
	}
}

func BenchmarkParseNegativeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonenumber.Parse("aeoighowHG", "FI")
	}
}
