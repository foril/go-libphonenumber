package phonenumber_test

import (
	"fmt"

	"github.com/foril/go-libphonenumber"
)

func ExampleCountryCodeFi() {
	number1 := "+358401231234"
	cc1 := phonenumber.CountryCode(number1)
	number2 := "+15417543010"
	cc2 := phonenumber.CountryCode(number2)
	fmt.Printf("%d and %d", cc1, cc2)
	// Output: 358 and 1
}
