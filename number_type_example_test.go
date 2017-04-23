package phonenumber_test

import (
	"fmt"

	"github.com/foril/go-libphonenumber"
)

func ExampleNumberType() {
	number1 := "+358401231234"
	cc1 := phonenumber.NumberType(number1)
	number2 := "+15417543010"
	cc2 := phonenumber.NumberType(number2)
	fmt.Printf("%s and %s", cc1, cc2)
	// Output: MOBILE and FIXED_LINE_OR_MOBILE

}
