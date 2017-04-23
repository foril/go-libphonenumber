package phonenumber_test

import (
	"fmt"

	"github.com/foril/go-libphonenumber"
)

func ExampleRegionCode() {
	number1 := "+358401231234"
	cc1 := phonenumber.RegionCode(number1)
	number2 := "+15417543010"
	cc2 := phonenumber.RegionCode(number2)
	fmt.Printf("%s and %s", cc1, cc2)
	// Output: FI and US
}
