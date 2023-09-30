// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/raft.proto

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

// Validate checks the field values on RaftLogEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RaftLogEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaftLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RaftLogEntryMultiError, or
// nil if none found.
func (m *RaftLogEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *RaftLogEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Key

	// no validation rules for Value

	if all {
		switch v := interface{}(m.GetTtl()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RaftLogEntryValidationError{
					field:  "Ttl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RaftLogEntryValidationError{
					field:  "Ttl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTtl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RaftLogEntryValidationError{
				field:  "Ttl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RaftLogEntryMultiError(errors)
	}

	return nil
}

// RaftLogEntryMultiError is an error wrapping multiple validation errors
// returned by RaftLogEntry.ValidateAll() if the designated constraints aren't met.
type RaftLogEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaftLogEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaftLogEntryMultiError) AllErrors() []error { return m }

// RaftLogEntryValidationError is the validation error returned by
// RaftLogEntry.Validate if the designated constraints aren't met.
type RaftLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaftLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaftLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaftLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaftLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaftLogEntryValidationError) ErrorName() string { return "RaftLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e RaftLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaftLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaftLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaftLogEntryValidationError{}

// Validate checks the field values on RaftApplyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RaftApplyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaftApplyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RaftApplyResponseMultiError, or nil if none found.
func (m *RaftApplyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RaftApplyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Time

	// no validation rules for Error

	if len(errors) > 0 {
		return RaftApplyResponseMultiError(errors)
	}

	return nil
}

// RaftApplyResponseMultiError is an error wrapping multiple validation errors
// returned by RaftApplyResponse.ValidateAll() if the designated constraints
// aren't met.
type RaftApplyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaftApplyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaftApplyResponseMultiError) AllErrors() []error { return m }

// RaftApplyResponseValidationError is the validation error returned by
// RaftApplyResponse.Validate if the designated constraints aren't met.
type RaftApplyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaftApplyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaftApplyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaftApplyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaftApplyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaftApplyResponseValidationError) ErrorName() string {
	return "RaftApplyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RaftApplyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaftApplyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaftApplyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaftApplyResponseValidationError{}

// Validate checks the field values on RaftSnapshot with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RaftSnapshot) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaftSnapshot with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RaftSnapshotMultiError, or
// nil if none found.
func (m *RaftSnapshot) ValidateAll() error {
	return m.validate(true)
}

func (m *RaftSnapshot) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetKv() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RaftSnapshotValidationError{
						field:  fmt.Sprintf("Kv[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RaftSnapshotValidationError{
						field:  fmt.Sprintf("Kv[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RaftSnapshotValidationError{
					field:  fmt.Sprintf("Kv[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RaftSnapshotMultiError(errors)
	}

	return nil
}

// RaftSnapshotMultiError is an error wrapping multiple validation errors
// returned by RaftSnapshot.ValidateAll() if the designated constraints aren't met.
type RaftSnapshotMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaftSnapshotMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaftSnapshotMultiError) AllErrors() []error { return m }

// RaftSnapshotValidationError is the validation error returned by
// RaftSnapshot.Validate if the designated constraints aren't met.
type RaftSnapshotValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaftSnapshotValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaftSnapshotValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaftSnapshotValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaftSnapshotValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaftSnapshotValidationError) ErrorName() string { return "RaftSnapshotValidationError" }

// Error satisfies the builtin error interface
func (e RaftSnapshotValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaftSnapshot.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaftSnapshotValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaftSnapshotValidationError{}

// Validate checks the field values on RaftDataItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RaftDataItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaftDataItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RaftDataItemMultiError, or
// nil if none found.
func (m *RaftDataItem) ValidateAll() error {
	return m.validate(true)
}

func (m *RaftDataItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if all {
		switch v := interface{}(m.GetTtl()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RaftDataItemValidationError{
					field:  "Ttl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RaftDataItemValidationError{
					field:  "Ttl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTtl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RaftDataItemValidationError{
				field:  "Ttl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RaftDataItemMultiError(errors)
	}

	return nil
}

// RaftDataItemMultiError is an error wrapping multiple validation errors
// returned by RaftDataItem.ValidateAll() if the designated constraints aren't met.
type RaftDataItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaftDataItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaftDataItemMultiError) AllErrors() []error { return m }

// RaftDataItemValidationError is the validation error returned by
// RaftDataItem.Validate if the designated constraints aren't met.
type RaftDataItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaftDataItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaftDataItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaftDataItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaftDataItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaftDataItemValidationError) ErrorName() string { return "RaftDataItemValidationError" }

// Error satisfies the builtin error interface
func (e RaftDataItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaftDataItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaftDataItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaftDataItemValidationError{}
