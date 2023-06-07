// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/template/v1/api/api.proto

package api

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

// Validate checks the field values on RegistrationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegistrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegistrationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegistrationRequestMultiError, or nil if none found.
func (m *RegistrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegistrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTest()) < 1 {
		err := RegistrationRequestValidationError{
			field:  "Test",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegistrationRequestMultiError(errors)
	}

	return nil
}

// RegistrationRequestMultiError is an error wrapping multiple validation
// errors returned by RegistrationRequest.ValidateAll() if the designated
// constraints aren't met.
type RegistrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistrationRequestMultiError) AllErrors() []error { return m }

// RegistrationRequestValidationError is the validation error returned by
// RegistrationRequest.Validate if the designated constraints aren't met.
type RegistrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistrationRequestValidationError) ErrorName() string {
	return "RegistrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegistrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistrationRequestValidationError{}

// Validate checks the field values on AttachmentRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AttachmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttachmentRequestMultiError, or nil if none found.
func (m *AttachmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AttachmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AttachmentRequestMultiError(errors)
	}

	return nil
}

// AttachmentRequestMultiError is an error wrapping multiple validation errors
// returned by AttachmentRequest.ValidateAll() if the designated constraints
// aren't met.
type AttachmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttachmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttachmentRequestMultiError) AllErrors() []error { return m }

// AttachmentRequestValidationError is the validation error returned by
// AttachmentRequest.Validate if the designated constraints aren't met.
type AttachmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttachmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttachmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttachmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttachmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttachmentRequestValidationError) ErrorName() string {
	return "AttachmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AttachmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttachmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttachmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttachmentRequestValidationError{}
