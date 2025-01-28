// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gateway/v1/gateway.proto

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

// Validate checks the field values on IdempotentRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdempotentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdempotentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdempotentRequestMultiError, or nil if none found.
func (m *IdempotentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdempotentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IdempotentRequestMultiError(errors)
	}

	return nil
}

// IdempotentRequestMultiError is an error wrapping multiple validation errors
// returned by IdempotentRequest.ValidateAll() if the designated constraints
// aren't met.
type IdempotentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdempotentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdempotentRequestMultiError) AllErrors() []error { return m }

// IdempotentRequestValidationError is the validation error returned by
// IdempotentRequest.Validate if the designated constraints aren't met.
type IdempotentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdempotentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdempotentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdempotentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdempotentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdempotentRequestValidationError) ErrorName() string {
	return "IdempotentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IdempotentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdempotentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdempotentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdempotentRequestValidationError{}

// Validate checks the field values on IdempotentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IdempotentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdempotentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdempotentResponseMultiError, or nil if none found.
func (m *IdempotentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IdempotentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return IdempotentResponseMultiError(errors)
	}

	return nil
}

// IdempotentResponseMultiError is an error wrapping multiple validation errors
// returned by IdempotentResponse.ValidateAll() if the designated constraints
// aren't met.
type IdempotentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdempotentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdempotentResponseMultiError) AllErrors() []error { return m }

// IdempotentResponseValidationError is the validation error returned by
// IdempotentResponse.Validate if the designated constraints aren't met.
type IdempotentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdempotentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdempotentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdempotentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdempotentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdempotentResponseValidationError) ErrorName() string {
	return "IdempotentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IdempotentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdempotentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdempotentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdempotentResponseValidationError{}

// Validate checks the field values on GetVersionRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetVersionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVersionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVersionRequestMultiError, or nil if none found.
func (m *GetVersionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVersionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetVersionRequestMultiError(errors)
	}

	return nil
}

// GetVersionRequestMultiError is an error wrapping multiple validation errors
// returned by GetVersionRequest.ValidateAll() if the designated constraints
// aren't met.
type GetVersionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVersionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVersionRequestMultiError) AllErrors() []error { return m }

// GetVersionRequestValidationError is the validation error returned by
// GetVersionRequest.Validate if the designated constraints aren't met.
type GetVersionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVersionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVersionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVersionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVersionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVersionRequestValidationError) ErrorName() string {
	return "GetVersionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetVersionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVersionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVersionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVersionRequestValidationError{}

// Validate checks the field values on GetVersionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetVersionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVersionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVersionResponseMultiError, or nil if none found.
func (m *GetVersionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVersionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GitVersion

	// no validation rules for GitCommit

	// no validation rules for GitTreeState

	// no validation rules for BuildDate

	// no validation rules for GoVersion

	// no validation rules for Compiler

	// no validation rules for Platform

	if len(errors) > 0 {
		return GetVersionResponseMultiError(errors)
	}

	return nil
}

// GetVersionResponseMultiError is an error wrapping multiple validation errors
// returned by GetVersionResponse.ValidateAll() if the designated constraints
// aren't met.
type GetVersionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVersionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVersionResponseMultiError) AllErrors() []error { return m }

// GetVersionResponseValidationError is the validation error returned by
// GetVersionResponse.Validate if the designated constraints aren't met.
type GetVersionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVersionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVersionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVersionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVersionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVersionResponseValidationError) ErrorName() string {
	return "GetVersionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetVersionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVersionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVersionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVersionResponseValidationError{}
