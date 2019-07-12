// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/fault/v2/fault.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on FaultAbort with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultAbort) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultAbortValidationError{
				Field:  "Percentage",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.ErrorType.(type) {

	case *FaultAbort_HttpStatus:

		if val := m.GetHttpStatus(); val < 200 || val >= 600 {
			return FaultAbortValidationError{
				Field:  "HttpStatus",
				Reason: "value must be inside range [200, 600)",
			}
		}

	default:
		return FaultAbortValidationError{
			Field:  "ErrorType",
			Reason: "value is required",
		}

	}

	return nil
}

// FaultAbortValidationError is the validation error returned by
// FaultAbort.Validate if the designated constraints aren't met.
type FaultAbortValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultAbortValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultAbortValidationError{}

// Validate checks the field values on HTTPFault with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HTTPFault) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				Field:  "Delay",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				Field:  "Abort",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for UpstreamCluster

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HTTPFaultValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetMaxActiveFaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				Field:  "MaxActiveFaults",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HTTPFaultValidationError is the validation error returned by
// HTTPFault.Validate if the designated constraints aren't met.
type HTTPFaultValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HTTPFaultValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPFault.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HTTPFaultValidationError{}
