package phonenumber

import (
	"encoding/json"
)

// Error represents errors when handling phone numbers
type Error string

var mapNumberError = map[int]string{
	0: "No parsing error",
	1: "Invalid country code",
	2: "Not a number",
	3: "Too short after International Direct Dial Prefixes",
	4: "Too short National Significant Number",
	5: "Too long National Significant Number",
	6: "Invalid number",
	7: "Invalid number for region code",
}

// Error implements error interface
func (e *Error) Error() string {
	return string(*e)
}

// NewError constructs a new error
func NewError(msg string) *Error {
	err := Error(msg)
	return &err
}

// MarshalJSON implements json.Marshaler
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(*e)
}

// UnmarshalJSON implements json.Unmarshaler
func (e *Error) UnmarshalJSON(data []byte) error {
	var errorStr string
	err := json.Unmarshal(data, &errorStr)
	if err != nil {
		return err
	}
	if errorStr != "" {
		*e = Error(errorStr)
	}
	return nil
}
