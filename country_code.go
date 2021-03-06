package phonenumber

/*
#cgo CPPFLAGS: -I/usr/local/include
#cgo CPPFLAGS: -Wall -Werror
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -lphonenumber

#include <stdlib.h>
#include "phonenumber.h"
*/
import "C"

import (
	"unsafe"
)

// CountryCode returns the country code for the supplied number / ex : 1, 33, 31 .....
func CountryCode(number string) C.int {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	res := C.get_country_code(cNum)
	return res
}
