// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/webrtc.proto

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

// Validate checks the field values on StartDataChannelRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StartDataChannelRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartDataChannelRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StartDataChannelRequestMultiError, or nil if none found.
func (m *StartDataChannelRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StartDataChannelRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeID

	// no validation rules for Proto

	// no validation rules for Dst

	// no validation rules for Port

	// no validation rules for Answer

	// no validation rules for Candidate

	if len(errors) > 0 {
		return StartDataChannelRequestMultiError(errors)
	}

	return nil
}

// StartDataChannelRequestMultiError is an error wrapping multiple validation
// errors returned by StartDataChannelRequest.ValidateAll() if the designated
// constraints aren't met.
type StartDataChannelRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartDataChannelRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartDataChannelRequestMultiError) AllErrors() []error { return m }

// StartDataChannelRequestValidationError is the validation error returned by
// StartDataChannelRequest.Validate if the designated constraints aren't met.
type StartDataChannelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartDataChannelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartDataChannelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartDataChannelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartDataChannelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartDataChannelRequestValidationError) ErrorName() string {
	return "StartDataChannelRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StartDataChannelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartDataChannelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartDataChannelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartDataChannelRequestValidationError{}

// Validate checks the field values on DataChannelOffer with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DataChannelOffer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataChannelOffer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DataChannelOfferMultiError, or nil if none found.
func (m *DataChannelOffer) ValidateAll() error {
	return m.validate(true)
}

func (m *DataChannelOffer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Offer

	// no validation rules for Candidate

	if len(errors) > 0 {
		return DataChannelOfferMultiError(errors)
	}

	return nil
}

// DataChannelOfferMultiError is an error wrapping multiple validation errors
// returned by DataChannelOffer.ValidateAll() if the designated constraints
// aren't met.
type DataChannelOfferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataChannelOfferMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataChannelOfferMultiError) AllErrors() []error { return m }

// DataChannelOfferValidationError is the validation error returned by
// DataChannelOffer.Validate if the designated constraints aren't met.
type DataChannelOfferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataChannelOfferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataChannelOfferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataChannelOfferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataChannelOfferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataChannelOfferValidationError) ErrorName() string { return "DataChannelOfferValidationError" }

// Error satisfies the builtin error interface
func (e DataChannelOfferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataChannelOffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataChannelOfferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataChannelOfferValidationError{}
