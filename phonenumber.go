package phonenumber

/*
#cgo CPPFLAGS: -I/usr/local/include
#cgo CPPFLAGS: -Wall -Werror -Wno-deprecated-declarations
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -lphonenumber

#include <stdlib.h>
#include "phonenumber.h"
*/
import "C"

// PhoneNumber stores information about a number
type PhoneNumber struct {
	// Original number given to the function
	Number string `json:"number"`
	// Number in normalized format (E164)
	Normalized *string `json:"normalized"`
	// Any error given
	Error *Error `json:"error"`
	// If the number was valid in the first place
	Valid bool `json:"valid"`
	//Country code
	CountryCode C.int `json:"countryCode"`
	// Number Type ex : MOBILE
	NumberType string `json:"numberType"`
	// Region Code ex : US
	RegionCode string `json:"regionCode"`
}

func phoneNumberInfoFromCStruct(info *C.struct_phone_info) PhoneNumber {
	pn := PhoneNumber{}
	if info == nil {
		return pn
	}
	pn.Valid = info.valid != 0
	if info.error != 0 {
		pn.Error = NewError(mapNumberError[int(info.error)])
	}
	if info.number != nil {
		pn.Number = C.GoString(info.number)
	}
	if info.normalized != nil {
		normalized := C.GoString(info.normalized)
		pn.Normalized = &normalized
	}
	pn.CountryCode = info.countryCode
	pn.NumberType = mapNumberType[int(info.numberType)]
	if info.regionCode != nil {
		pn.RegionCode = C.GoString(info.regionCode)
	}

	return pn
}
