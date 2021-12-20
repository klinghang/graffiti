// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/common/matcher/v3/matcher.proto

package envoy_config_common_matcher_v3

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

// Validate checks the field values on Matcher with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Matcher) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOnNoMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MatcherValidationError{
				field:  "OnNoMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.MatcherType.(type) {

	case *Matcher_MatcherList_:

		if v, ok := interface{}(m.GetMatcherList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatcherValidationError{
					field:  "MatcherList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherTree_:

		if v, ok := interface{}(m.GetMatcherTree()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatcherValidationError{
					field:  "MatcherTree",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return MatcherValidationError{
			field:  "MatcherType",
			reason: "value is required",
		}

	}

	return nil
}

// MatcherValidationError is the validation error returned by Matcher.Validate
// if the designated constraints aren't met.
type MatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatcherValidationError) ErrorName() string { return "MatcherValidationError" }

// Error satisfies the builtin error interface
func (e MatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatcherValidationError{}

// Validate checks the field values on MatchPredicate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MatchPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *MatchPredicate_OrMatch:

		if v, ok := interface{}(m.GetOrMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "OrMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_AndMatch:

		if v, ok := interface{}(m.GetAndMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "AndMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_NotMatch:

		if v, ok := interface{}(m.GetNotMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "NotMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_AnyMatch:

		if m.GetAnyMatch() != true {
			return MatchPredicateValidationError{
				field:  "AnyMatch",
				reason: "value must equal true",
			}
		}

	case *MatchPredicate_HttpRequestHeadersMatch:

		if v, ok := interface{}(m.GetHttpRequestHeadersMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpRequestHeadersMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpRequestTrailersMatch:

		if v, ok := interface{}(m.GetHttpRequestTrailersMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpRequestTrailersMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpResponseHeadersMatch:

		if v, ok := interface{}(m.GetHttpResponseHeadersMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpResponseHeadersMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpResponseTrailersMatch:

		if v, ok := interface{}(m.GetHttpResponseTrailersMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpResponseTrailersMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpRequestGenericBodyMatch:

		if v, ok := interface{}(m.GetHttpRequestGenericBodyMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpRequestGenericBodyMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpResponseGenericBodyMatch:

		if v, ok := interface{}(m.GetHttpResponseGenericBodyMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "HttpResponseGenericBodyMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return MatchPredicateValidationError{
			field:  "Rule",
			reason: "value is required",
		}

	}

	return nil
}

// MatchPredicateValidationError is the validation error returned by
// MatchPredicate.Validate if the designated constraints aren't met.
type MatchPredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicateValidationError) ErrorName() string { return "MatchPredicateValidationError" }

// Error satisfies the builtin error interface
func (e MatchPredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicateValidationError{}

// Validate checks the field values on HttpHeadersMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpHeadersMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpHeadersMatchValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpHeadersMatchValidationError is the validation error returned by
// HttpHeadersMatch.Validate if the designated constraints aren't met.
type HttpHeadersMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpHeadersMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpHeadersMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpHeadersMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpHeadersMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpHeadersMatchValidationError) ErrorName() string { return "HttpHeadersMatchValidationError" }

// Error satisfies the builtin error interface
func (e HttpHeadersMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpHeadersMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpHeadersMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpHeadersMatchValidationError{}

// Validate checks the field values on HttpGenericBodyMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpGenericBodyMatch) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BytesLimit

	if len(m.GetPatterns()) < 1 {
		return HttpGenericBodyMatchValidationError{
			field:  "Patterns",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPatterns() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpGenericBodyMatchValidationError{
					field:  fmt.Sprintf("Patterns[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpGenericBodyMatchValidationError is the validation error returned by
// HttpGenericBodyMatch.Validate if the designated constraints aren't met.
type HttpGenericBodyMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpGenericBodyMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpGenericBodyMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpGenericBodyMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpGenericBodyMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpGenericBodyMatchValidationError) ErrorName() string {
	return "HttpGenericBodyMatchValidationError"
}

// Error satisfies the builtin error interface
func (e HttpGenericBodyMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpGenericBodyMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpGenericBodyMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpGenericBodyMatchValidationError{}

// Validate checks the field values on Matcher_OnMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Matcher_OnMatch) Validate() error {
	if m == nil {
		return nil
	}

	switch m.OnMatch.(type) {

	case *Matcher_OnMatch_Matcher:

		if v, ok := interface{}(m.GetMatcher()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_OnMatchValidationError{
					field:  "Matcher",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_OnMatch_Action:

		if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_OnMatchValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Matcher_OnMatchValidationError{
			field:  "OnMatch",
			reason: "value is required",
		}

	}

	return nil
}

// Matcher_OnMatchValidationError is the validation error returned by
// Matcher_OnMatch.Validate if the designated constraints aren't met.
type Matcher_OnMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_OnMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_OnMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_OnMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_OnMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_OnMatchValidationError) ErrorName() string { return "Matcher_OnMatchValidationError" }

// Error satisfies the builtin error interface
func (e Matcher_OnMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_OnMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_OnMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_OnMatchValidationError{}

// Validate checks the field values on Matcher_MatcherList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Matcher_MatcherList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetMatchers()) < 1 {
		return Matcher_MatcherListValidationError{
			field:  "Matchers",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetMatchers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherListValidationError{
					field:  fmt.Sprintf("Matchers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Matcher_MatcherListValidationError is the validation error returned by
// Matcher_MatcherList.Validate if the designated constraints aren't met.
type Matcher_MatcherListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherListValidationError) ErrorName() string {
	return "Matcher_MatcherListValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherListValidationError{}

// Validate checks the field values on Matcher_MatcherTree with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Matcher_MatcherTree) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetInput() == nil {
		return Matcher_MatcherTreeValidationError{
			field:  "Input",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Matcher_MatcherTreeValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.TreeType.(type) {

	case *Matcher_MatcherTree_ExactMatchMap:

		if v, ok := interface{}(m.GetExactMatchMap()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherTreeValidationError{
					field:  "ExactMatchMap",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherTree_PrefixMatchMap:

		if v, ok := interface{}(m.GetPrefixMatchMap()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherTreeValidationError{
					field:  "PrefixMatchMap",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherTree_CustomMatch:

		if v, ok := interface{}(m.GetCustomMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherTreeValidationError{
					field:  "CustomMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Matcher_MatcherTreeValidationError{
			field:  "TreeType",
			reason: "value is required",
		}

	}

	return nil
}

// Matcher_MatcherTreeValidationError is the validation error returned by
// Matcher_MatcherTree.Validate if the designated constraints aren't met.
type Matcher_MatcherTreeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherTreeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherTreeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherTreeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherTreeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherTreeValidationError) ErrorName() string {
	return "Matcher_MatcherTreeValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherTreeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherTree.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherTreeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherTreeValidationError{}

// Validate checks the field values on Matcher_MatcherList_Predicate with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Matcher_MatcherList_Predicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchType.(type) {

	case *Matcher_MatcherList_Predicate_SinglePredicate_:

		if v, ok := interface{}(m.GetSinglePredicate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_PredicateValidationError{
					field:  "SinglePredicate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherList_Predicate_OrMatcher:

		if v, ok := interface{}(m.GetOrMatcher()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_PredicateValidationError{
					field:  "OrMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherList_Predicate_AndMatcher:

		if v, ok := interface{}(m.GetAndMatcher()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_PredicateValidationError{
					field:  "AndMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Matcher_MatcherList_PredicateValidationError{
			field:  "MatchType",
			reason: "value is required",
		}

	}

	return nil
}

// Matcher_MatcherList_PredicateValidationError is the validation error
// returned by Matcher_MatcherList_Predicate.Validate if the designated
// constraints aren't met.
type Matcher_MatcherList_PredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherList_PredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherList_PredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherList_PredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherList_PredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherList_PredicateValidationError) ErrorName() string {
	return "Matcher_MatcherList_PredicateValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherList_PredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherList_Predicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherList_PredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherList_PredicateValidationError{}

// Validate checks the field values on Matcher_MatcherList_FieldMatcher with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *Matcher_MatcherList_FieldMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPredicate() == nil {
		return Matcher_MatcherList_FieldMatcherValidationError{
			field:  "Predicate",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPredicate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Matcher_MatcherList_FieldMatcherValidationError{
				field:  "Predicate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetOnMatch() == nil {
		return Matcher_MatcherList_FieldMatcherValidationError{
			field:  "OnMatch",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOnMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Matcher_MatcherList_FieldMatcherValidationError{
				field:  "OnMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Matcher_MatcherList_FieldMatcherValidationError is the validation error
// returned by Matcher_MatcherList_FieldMatcher.Validate if the designated
// constraints aren't met.
type Matcher_MatcherList_FieldMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherList_FieldMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherList_FieldMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherList_FieldMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherList_FieldMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherList_FieldMatcherValidationError) ErrorName() string {
	return "Matcher_MatcherList_FieldMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherList_FieldMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherList_FieldMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherList_FieldMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherList_FieldMatcherValidationError{}

// Validate checks the field values on
// Matcher_MatcherList_Predicate_SinglePredicate with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Matcher_MatcherList_Predicate_SinglePredicate) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetInput() == nil {
		return Matcher_MatcherList_Predicate_SinglePredicateValidationError{
			field:  "Input",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Matcher_MatcherList_Predicate_SinglePredicateValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Matcher.(type) {

	case *Matcher_MatcherList_Predicate_SinglePredicate_ValueMatch:

		if v, ok := interface{}(m.GetValueMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_Predicate_SinglePredicateValidationError{
					field:  "ValueMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Matcher_MatcherList_Predicate_SinglePredicate_CustomMatch:

		if v, ok := interface{}(m.GetCustomMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_Predicate_SinglePredicateValidationError{
					field:  "CustomMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Matcher_MatcherList_Predicate_SinglePredicateValidationError{
			field:  "Matcher",
			reason: "value is required",
		}

	}

	return nil
}

// Matcher_MatcherList_Predicate_SinglePredicateValidationError is the
// validation error returned by
// Matcher_MatcherList_Predicate_SinglePredicate.Validate if the designated
// constraints aren't met.
type Matcher_MatcherList_Predicate_SinglePredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) ErrorName() string {
	return "Matcher_MatcherList_Predicate_SinglePredicateValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherList_Predicate_SinglePredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherList_Predicate_SinglePredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherList_Predicate_SinglePredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherList_Predicate_SinglePredicateValidationError{}

// Validate checks the field values on
// Matcher_MatcherList_Predicate_PredicateList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Matcher_MatcherList_Predicate_PredicateList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPredicate()) < 2 {
		return Matcher_MatcherList_Predicate_PredicateListValidationError{
			field:  "Predicate",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetPredicate() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherList_Predicate_PredicateListValidationError{
					field:  fmt.Sprintf("Predicate[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Matcher_MatcherList_Predicate_PredicateListValidationError is the validation
// error returned by Matcher_MatcherList_Predicate_PredicateList.Validate if
// the designated constraints aren't met.
type Matcher_MatcherList_Predicate_PredicateListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) ErrorName() string {
	return "Matcher_MatcherList_Predicate_PredicateListValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherList_Predicate_PredicateListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherList_Predicate_PredicateList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherList_Predicate_PredicateListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherList_Predicate_PredicateListValidationError{}

// Validate checks the field values on Matcher_MatcherTree_MatchMap with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Matcher_MatcherTree_MatchMap) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetMap()) < 1 {
		return Matcher_MatcherTree_MatchMapValidationError{
			field:  "Map",
			reason: "value must contain at least 1 pair(s)",
		}
	}

	for key, val := range m.GetMap() {
		_ = val

		// no validation rules for Map[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Matcher_MatcherTree_MatchMapValidationError{
					field:  fmt.Sprintf("Map[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Matcher_MatcherTree_MatchMapValidationError is the validation error returned
// by Matcher_MatcherTree_MatchMap.Validate if the designated constraints
// aren't met.
type Matcher_MatcherTree_MatchMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Matcher_MatcherTree_MatchMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Matcher_MatcherTree_MatchMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Matcher_MatcherTree_MatchMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Matcher_MatcherTree_MatchMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Matcher_MatcherTree_MatchMapValidationError) ErrorName() string {
	return "Matcher_MatcherTree_MatchMapValidationError"
}

// Error satisfies the builtin error interface
func (e Matcher_MatcherTree_MatchMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatcher_MatcherTree_MatchMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Matcher_MatcherTree_MatchMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Matcher_MatcherTree_MatchMapValidationError{}

// Validate checks the field values on MatchPredicate_MatchSet with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_MatchSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 2 {
		return MatchPredicate_MatchSetValidationError{
			field:  "Rules",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicate_MatchSetValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MatchPredicate_MatchSetValidationError is the validation error returned by
// MatchPredicate_MatchSet.Validate if the designated constraints aren't met.
type MatchPredicate_MatchSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicate_MatchSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicate_MatchSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicate_MatchSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicate_MatchSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicate_MatchSetValidationError) ErrorName() string {
	return "MatchPredicate_MatchSetValidationError"
}

// Error satisfies the builtin error interface
func (e MatchPredicate_MatchSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_MatchSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicate_MatchSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicate_MatchSetValidationError{}

// Validate checks the field values on HttpGenericBodyMatch_GenericTextMatch
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *HttpGenericBodyMatch_GenericTextMatch) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *HttpGenericBodyMatch_GenericTextMatch_StringMatch:

		if utf8.RuneCountInString(m.GetStringMatch()) < 1 {
			return HttpGenericBodyMatch_GenericTextMatchValidationError{
				field:  "StringMatch",
				reason: "value length must be at least 1 runes",
			}
		}

	case *HttpGenericBodyMatch_GenericTextMatch_BinaryMatch:

		if len(m.GetBinaryMatch()) < 1 {
			return HttpGenericBodyMatch_GenericTextMatchValidationError{
				field:  "BinaryMatch",
				reason: "value length must be at least 1 bytes",
			}
		}

	default:
		return HttpGenericBodyMatch_GenericTextMatchValidationError{
			field:  "Rule",
			reason: "value is required",
		}

	}

	return nil
}

// HttpGenericBodyMatch_GenericTextMatchValidationError is the validation error
// returned by HttpGenericBodyMatch_GenericTextMatch.Validate if the
// designated constraints aren't met.
type HttpGenericBodyMatch_GenericTextMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) ErrorName() string {
	return "HttpGenericBodyMatch_GenericTextMatchValidationError"
}

// Error satisfies the builtin error interface
func (e HttpGenericBodyMatch_GenericTextMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpGenericBodyMatch_GenericTextMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpGenericBodyMatch_GenericTextMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpGenericBodyMatch_GenericTextMatchValidationError{}
