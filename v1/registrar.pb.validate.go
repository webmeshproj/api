// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/registrar.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PublicKey

	// no validation rules for Alias

	if all {
		switch v := interface{}(m.GetExpiry()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterRequestValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterRequestValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterRequestValidationError{
				field:  "Expiry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterResponseMultiError, or nil if none found.
func (m *RegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RegisterResponseMultiError(errors)
	}

	return nil
}

// RegisterResponseMultiError is an error wrapping multiple validation errors
// returned by RegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type RegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterResponseMultiError) AllErrors() []error { return m }

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on LookupRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LookupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LookupRequestMultiError, or
// nil if none found.
func (m *LookupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for PublicKey

	// no validation rules for Alias

	if len(errors) > 0 {
		return LookupRequestMultiError(errors)
	}

	return nil
}

// LookupRequestMultiError is an error wrapping multiple validation errors
// returned by LookupRequest.ValidateAll() if the designated constraints
// aren't met.
type LookupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupRequestMultiError) AllErrors() []error { return m }

// LookupRequestValidationError is the validation error returned by
// LookupRequest.Validate if the designated constraints aren't met.
type LookupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupRequestValidationError) ErrorName() string { return "LookupRequestValidationError" }

// Error satisfies the builtin error interface
func (e LookupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupRequestValidationError{}

// Validate checks the field values on LookupResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LookupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LookupResponseMultiError,
// or nil if none found.
func (m *LookupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for PublicKey

	// no validation rules for Alias

	if len(errors) > 0 {
		return LookupResponseMultiError(errors)
	}

	return nil
}

// LookupResponseMultiError is an error wrapping multiple validation errors
// returned by LookupResponse.ValidateAll() if the designated constraints
// aren't met.
type LookupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupResponseMultiError) AllErrors() []error { return m }

// LookupResponseValidationError is the validation error returned by
// LookupResponse.Validate if the designated constraints aren't met.
type LookupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupResponseValidationError) ErrorName() string { return "LookupResponseValidationError" }

// Error satisfies the builtin error interface
func (e LookupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupResponseValidationError{}
