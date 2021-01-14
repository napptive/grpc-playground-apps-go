// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: playground-apps/entities.proto

package grpc_playground_apps_go

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Application) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AppId

	// no validation rules for Name

	// no validation rules for VisualId

	// no validation rules for Status

	// no validation rules for Description

	// no validation rules for ComponentStatus

	if v, ok := interface{}(m.GetInstance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "Instance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StatusName

	// no validation rules for ComponentStatusName

	return nil
}

// ApplicationValidationError is the validation error returned by
// Application.Validate if the designated constraints aren't met.
type ApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationValidationError) ErrorName() string { return "ApplicationValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationValidationError{}

// Validate checks the field values on AppListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AppListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AppListResponseValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for From

	// no validation rules for To

	return nil
}

// AppListResponseValidationError is the validation error returned by
// AppListResponse.Validate if the designated constraints aren't met.
type AppListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppListResponseValidationError) ErrorName() string { return "AppListResponseValidationError" }

// Error satisfies the builtin error interface
func (e AppListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppListResponseValidationError{}
