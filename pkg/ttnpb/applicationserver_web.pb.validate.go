// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

// ValidateFields checks the field values on ApplicationWebhookIdentifiers with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ApplicationWebhookIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationWebhookIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_ids":

			if v, ok := interface{}(&m.ApplicationIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookIdentifiersValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "webhook_id":

			if utf8.RuneCountInString(m.GetWebhookID()) > 36 {
				return ApplicationWebhookIdentifiersValidationError{
					field:  "webhook_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ApplicationWebhookIdentifiers_WebhookID_Pattern.MatchString(m.GetWebhookID()) {
				return ApplicationWebhookIdentifiersValidationError{
					field:  "webhook_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		default:
			return ApplicationWebhookIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationWebhookIdentifiersValidationError is the validation error
// returned by ApplicationWebhookIdentifiers.ValidateFields if the designated
// constraints aren't met.
type ApplicationWebhookIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationWebhookIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationWebhookIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationWebhookIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationWebhookIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationWebhookIdentifiersValidationError) ErrorName() string {
	return "ApplicationWebhookIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationWebhookIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationWebhookIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationWebhookIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationWebhookIdentifiersValidationError{}

var _ApplicationWebhookIdentifiers_WebhookID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on ApplicationWebhook with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationWebhook) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationWebhookFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.ApplicationWebhookIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "created_at":

			if v, ok := interface{}(&m.CreatedAt).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "updated_at":

			if v, ok := interface{}(&m.UpdatedAt).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "updated_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "base_url":

			if uri, err := url.Parse(m.GetBaseURL()); err != nil {
				return ApplicationWebhookValidationError{
					field:  "base_url",
					reason: "value must be a valid URI",
					cause:  err,
				}
			} else if !uri.IsAbs() {
				return ApplicationWebhookValidationError{
					field:  "base_url",
					reason: "value must be absolute",
				}
			}

		case "headers":
			// no validation rules for Headers
		case "format":
			// no validation rules for Format
		case "uplink_message":

			if v, ok := interface{}(m.GetUplinkMessage()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "uplink_message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "join_accept":

			if v, ok := interface{}(m.GetJoinAccept()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "join_accept",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "downlink_ack":

			if v, ok := interface{}(m.GetDownlinkAck()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "downlink_ack",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "downlink_nack":

			if v, ok := interface{}(m.GetDownlinkNack()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "downlink_nack",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "downlink_sent":

			if v, ok := interface{}(m.GetDownlinkSent()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "downlink_sent",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "downlink_failed":

			if v, ok := interface{}(m.GetDownlinkFailed()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "downlink_failed",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "downlink_queued":

			if v, ok := interface{}(m.GetDownlinkQueued()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "downlink_queued",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "location_solved":

			if v, ok := interface{}(m.GetLocationSolved()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ApplicationWebhookValidationError{
						field:  "location_solved",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return ApplicationWebhookValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationWebhookValidationError is the validation error returned by
// ApplicationWebhook.ValidateFields if the designated constraints aren't met.
type ApplicationWebhookValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationWebhookValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationWebhookValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationWebhookValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationWebhookValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationWebhookValidationError) ErrorName() string {
	return "ApplicationWebhookValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationWebhookValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationWebhook.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationWebhookValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationWebhookValidationError{}

// ValidateFields checks the field values on ApplicationWebhooks with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationWebhooks) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationWebhooksFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "webhooks":

			for idx, item := range m.GetWebhooks() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ApplicationWebhooksValidationError{
							field:  fmt.Sprintf("webhooks[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return ApplicationWebhooksValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationWebhooksValidationError is the validation error returned by
// ApplicationWebhooks.ValidateFields if the designated constraints aren't met.
type ApplicationWebhooksValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationWebhooksValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationWebhooksValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationWebhooksValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationWebhooksValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationWebhooksValidationError) ErrorName() string {
	return "ApplicationWebhooksValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationWebhooksValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationWebhooks.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationWebhooksValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationWebhooksValidationError{}

// ValidateFields checks the field values on ApplicationWebhookFormats with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationWebhookFormats) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationWebhookFormatsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "formats":
			// no validation rules for Formats
		default:
			return ApplicationWebhookFormatsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationWebhookFormatsValidationError is the validation error returned by
// ApplicationWebhookFormats.ValidateFields if the designated constraints
// aren't met.
type ApplicationWebhookFormatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationWebhookFormatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationWebhookFormatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationWebhookFormatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationWebhookFormatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationWebhookFormatsValidationError) ErrorName() string {
	return "ApplicationWebhookFormatsValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationWebhookFormatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationWebhookFormats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationWebhookFormatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationWebhookFormatsValidationError{}

// ValidateFields checks the field values on GetApplicationWebhookRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetApplicationWebhookRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetApplicationWebhookRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.ApplicationWebhookIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetApplicationWebhookRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetApplicationWebhookRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GetApplicationWebhookRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetApplicationWebhookRequestValidationError is the validation error returned
// by GetApplicationWebhookRequest.ValidateFields if the designated
// constraints aren't met.
type GetApplicationWebhookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplicationWebhookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplicationWebhookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApplicationWebhookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplicationWebhookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplicationWebhookRequestValidationError) ErrorName() string {
	return "GetApplicationWebhookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApplicationWebhookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplicationWebhookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplicationWebhookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplicationWebhookRequestValidationError{}

// ValidateFields checks the field values on ListApplicationWebhooksRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ListApplicationWebhooksRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListApplicationWebhooksRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_ids":

			if v, ok := interface{}(&m.ApplicationIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListApplicationWebhooksRequestValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListApplicationWebhooksRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return ListApplicationWebhooksRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListApplicationWebhooksRequestValidationError is the validation error
// returned by ListApplicationWebhooksRequest.ValidateFields if the designated
// constraints aren't met.
type ListApplicationWebhooksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplicationWebhooksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplicationWebhooksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplicationWebhooksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplicationWebhooksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplicationWebhooksRequestValidationError) ErrorName() string {
	return "ListApplicationWebhooksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplicationWebhooksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplicationWebhooksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplicationWebhooksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplicationWebhooksRequestValidationError{}

// ValidateFields checks the field values on SetApplicationWebhookRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *SetApplicationWebhookRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SetApplicationWebhookRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "webhook":

			if v, ok := interface{}(&m.ApplicationWebhook).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SetApplicationWebhookRequestValidationError{
						field:  "webhook",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SetApplicationWebhookRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return SetApplicationWebhookRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SetApplicationWebhookRequestValidationError is the validation error returned
// by SetApplicationWebhookRequest.ValidateFields if the designated
// constraints aren't met.
type SetApplicationWebhookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetApplicationWebhookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetApplicationWebhookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetApplicationWebhookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetApplicationWebhookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetApplicationWebhookRequestValidationError) ErrorName() string {
	return "SetApplicationWebhookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetApplicationWebhookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetApplicationWebhookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetApplicationWebhookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetApplicationWebhookRequestValidationError{}

// ValidateFields checks the field values on ApplicationWebhook_Message with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ApplicationWebhook_Message) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationWebhook_MessageFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "path":
			// no validation rules for Path
		default:
			return ApplicationWebhook_MessageValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationWebhook_MessageValidationError is the validation error returned
// by ApplicationWebhook_Message.ValidateFields if the designated constraints
// aren't met.
type ApplicationWebhook_MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationWebhook_MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationWebhook_MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationWebhook_MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationWebhook_MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationWebhook_MessageValidationError) ErrorName() string {
	return "ApplicationWebhook_MessageValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationWebhook_MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationWebhook_Message.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationWebhook_MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationWebhook_MessageValidationError{}
