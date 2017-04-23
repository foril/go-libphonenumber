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

var mapNumberType = map[int]string{
	0:  "FIXED_LINE",
	1:  "MOBILE",
	2:  "FIXED_LINE_OR_MOBILE",
	3:  "TOLL_FREE",
	4:  "PREMIUM_RATE",
	5:  "SHARED_COST",
	6:  "VOIP",
	7:  "PERSONAL_NUMBER",
	8:  "PAGER",
	9:  "UAN",
	10: "VOICEMAIL",
	11: "UNKNOWN",
}

// NumberType returns the number type for the supplied number / ex : FIXED_LINE, MOBILE .....
func NumberType(number string) string {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	res := mapNumberType[int(C.get_number_type(cNum))]
	return res
}
