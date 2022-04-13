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
)

// Validate checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Application) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApplicationMultiError, or
// nil if none found.
func (m *Application) ValidateAll() error {
	return m.validate(true)
}

func (m *Application) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for Name

	// no validation rules for VisualId

	// no validation rules for Status

	// no validation rules for Description

	// no validation rules for Version

	// no validation rules for ComponentStatus

	if all {
		switch v := interface{}(m.GetInstance()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "Instance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "Instance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInstance()).(interface{ Validate() error }); ok {
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

	for key, val := range m.GetComponentIngresses() {
		_ = val

		// no validation rules for ComponentIngresses[key]

		if all {
			switch v := interface{}(val).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("ComponentIngresses[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("ComponentIngresses[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationValidationError{
					field:  fmt.Sprintf("ComponentIngresses[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ApplicationMultiError(errors)
	}
	return nil
}

// ApplicationMultiError is an error wrapping multiple validation errors
// returned by Application.ValidateAll() if the designated constraints aren't met.
type ApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationMultiError) AllErrors() []error { return m }

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

// Validate checks the field values on IngressInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IngressInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IngressInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IngressInfoMultiError, or
// nil if none found.
func (m *IngressInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *IngressInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return IngressInfoMultiError(errors)
	}
	return nil
}

// IngressInfoMultiError is an error wrapping multiple validation errors
// returned by IngressInfo.ValidateAll() if the designated constraints aren't met.
type IngressInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IngressInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IngressInfoMultiError) AllErrors() []error { return m }

// IngressInfoValidationError is the validation error returned by
// IngressInfo.Validate if the designated constraints aren't met.
type IngressInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IngressInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IngressInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IngressInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IngressInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IngressInfoValidationError) ErrorName() string { return "IngressInfoValidationError" }

// Error satisfies the builtin error interface
func (e IngressInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIngressInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IngressInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IngressInfoValidationError{}

// Validate checks the field values on IngressList with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IngressList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IngressList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IngressListMultiError, or
// nil if none found.
func (m *IngressList) ValidateAll() error {
	return m.validate(true)
}

func (m *IngressList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ParentComponentName

	for idx, item := range m.GetIngresses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IngressListValidationError{
						field:  fmt.Sprintf("Ingresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IngressListValidationError{
						field:  fmt.Sprintf("Ingresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IngressListValidationError{
					field:  fmt.Sprintf("Ingresses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IngressListMultiError(errors)
	}
	return nil
}

// IngressListMultiError is an error wrapping multiple validation errors
// returned by IngressList.ValidateAll() if the designated constraints aren't met.
type IngressListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IngressListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IngressListMultiError) AllErrors() []error { return m }

// IngressListValidationError is the validation error returned by
// IngressList.Validate if the designated constraints aren't met.
type IngressListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IngressListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IngressListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IngressListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IngressListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IngressListValidationError) ErrorName() string { return "IngressListValidationError" }

// Error satisfies the builtin error interface
func (e IngressListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIngressList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IngressListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IngressListValidationError{}

// Validate checks the field values on AppInfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppInfoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppInfoRequestMultiError,
// or nil if none found.
func (m *AppInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EnvironmentQualifiedName

	// no validation rules for AccountId

	// no validation rules for EnvironmentId

	// no validation rules for ApplicationName

	if len(errors) > 0 {
		return AppInfoRequestMultiError(errors)
	}
	return nil
}

// AppInfoRequestMultiError is an error wrapping multiple validation errors
// returned by AppInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type AppInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppInfoRequestMultiError) AllErrors() []error { return m }

// AppInfoRequestValidationError is the validation error returned by
// AppInfoRequest.Validate if the designated constraints aren't met.
type AppInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppInfoRequestValidationError) ErrorName() string { return "AppInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppInfoRequestValidationError{}

// Validate checks the field values on AppSummaryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AppSummaryListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppSummaryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppSummaryListResponseMultiError, or nil if none found.
func (m *AppSummaryListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AppSummaryListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AppSummaryListResponseValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AppSummaryListResponseValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AppSummaryListResponseValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for From

	// no validation rules for To

	if len(errors) > 0 {
		return AppSummaryListResponseMultiError(errors)
	}
	return nil
}

// AppSummaryListResponseMultiError is an error wrapping multiple validation
// errors returned by AppSummaryListResponse.ValidateAll() if the designated
// constraints aren't met.
type AppSummaryListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppSummaryListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppSummaryListResponseMultiError) AllErrors() []error { return m }

// AppSummaryListResponseValidationError is the validation error returned by
// AppSummaryListResponse.Validate if the designated constraints aren't met.
type AppSummaryListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppSummaryListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppSummaryListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppSummaryListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppSummaryListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppSummaryListResponseValidationError) ErrorName() string {
	return "AppSummaryListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AppSummaryListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppSummaryListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppSummaryListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppSummaryListResponseValidationError{}

// Validate checks the field values on AppSummary with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppSummary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppSummary with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppSummaryMultiError, or
// nil if none found.
func (m *AppSummary) ValidateAll() error {
	return m.validate(true)
}

func (m *AppSummary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for Name

	// no validation rules for VisualId

	// no validation rules for Status

	// no validation rules for StatusName

	// no validation rules for ComponentStatus

	// no validation rules for ComponentStatusName

	if len(errors) > 0 {
		return AppSummaryMultiError(errors)
	}
	return nil
}

// AppSummaryMultiError is an error wrapping multiple validation errors
// returned by AppSummary.ValidateAll() if the designated constraints aren't met.
type AppSummaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppSummaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppSummaryMultiError) AllErrors() []error { return m }

// AppSummaryValidationError is the validation error returned by
// AppSummary.Validate if the designated constraints aren't met.
type AppSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppSummaryValidationError) ErrorName() string { return "AppSummaryValidationError" }

// Error satisfies the builtin error interface
func (e AppSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppSummaryValidationError{}

// Validate checks the field values on DeployApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeployApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeployApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeployApplicationRequestMultiError, or nil if none found.
func (m *DeployApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeployApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetApplicationData()) < 1 {
		err := DeployApplicationRequestValidationError{
			field:  "ApplicationData",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTargetEnvironmentQualifiedName()) < 1 {
		err := DeployApplicationRequestValidationError{
			field:  "TargetEnvironmentQualifiedName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeployApplicationRequestMultiError(errors)
	}
	return nil
}

// DeployApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by DeployApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type DeployApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeployApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeployApplicationRequestMultiError) AllErrors() []error { return m }

// DeployApplicationRequestValidationError is the validation error returned by
// DeployApplicationRequest.Validate if the designated constraints aren't met.
type DeployApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeployApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeployApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeployApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeployApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeployApplicationRequestValidationError) ErrorName() string {
	return "DeployApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeployApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeployApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeployApplicationRequestValidationError{}

// Validate checks the field values on RemoveApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveApplicationRequestMultiError, or nil if none found.
func (m *RemoveApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TargetEnvironmentQualifiedName

	// no validation rules for AccountId

	// no validation rules for EnvironmentId

	if utf8.RuneCountInString(m.GetApplicationName()) < 1 {
		err := RemoveApplicationRequestValidationError{
			field:  "ApplicationName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveApplicationRequestMultiError(errors)
	}
	return nil
}

// RemoveApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by RemoveApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type RemoveApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveApplicationRequestMultiError) AllErrors() []error { return m }

// RemoveApplicationRequestValidationError is the validation error returned by
// RemoveApplicationRequest.Validate if the designated constraints aren't met.
type RemoveApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveApplicationRequestValidationError) ErrorName() string {
	return "RemoveApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveApplicationRequestValidationError{}
