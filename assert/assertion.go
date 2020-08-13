package assert

import (
	"fmt"
	"strings"
	"testing"
)

// test represents the functions to be used by Assertion.
type test interface {
	Helper()
	Error(...interface{})
}

// Assertion is a wrapper around testing.T.
type Assertion struct {
	t test
}

// NewAssertion constructs and returns the wrapper.
func NewAssertion(t *testing.T) *Assertion {
	return newAssertion(t)
}

// newAssertion constructs and returns the wrapper.
func newAssertion(t test) *Assertion {
	if t == nil {
		panic("nil test")
	}

	return &Assertion{t}
}

// buildMessage builds an error message.
func buildMessage(message string, args ...interface{}) string {
	if len(args) == 0 {
		return message
	} else {
		if format, ok := args[0].(string); !ok {
			if len(args) == 1 {
				return format
			} else {
				var formatCount = strings.Count(format, "%") - 2*strings.Count(format, "%%")

				if formatCount != len(args)-1 {
					panic("invalid format")
				}

				return fmt.Sprintf(format, args[1:])
			}
		} else {
			panic("args[0] is not a string")
		}
	}
}

// True checks if actual is true.
func (assert *Assertion) True(actual bool, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage("expected: true, actual: false", args)

	if actual {
		return true
	}

	assert.t.Error(message)

	return false
}

// False checks if actual is false.
func (assert *Assertion) False(actual bool, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage("expected: false, actual: true", args)

	if !actual {
		return true
	}

	assert.t.Error(message)

	return false
}

// Nil checks if actual is nil.
func (assert *Assertion) Nil(actual interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("expected: nil, actual: %v", actual), args)

	if isNil(actual) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotNil checks if actual is not nil
func (assert *Assertion) NotNil(actual interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage("unexpected: nil, actual: nil", args)

	if !isNil(actual) {
		return true
	}

	assert.t.Error(message)

	return false
}

// Equals checks if expected and actual are equal.
func (assert *Assertion) Equals(expected interface{}, actual interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("expected: %v, actual: %v", expected, actual), args)

	if expected == actual {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotEquals checks if unexpected and actual are not equal.
func (assert *Assertion) NotEquals(unexpected interface{}, actual interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("unexpected: %v, actual: %v", unexpected, actual), args)

	if unexpected != actual {
		return true
	}

	assert.t.Error(message)

	return false

}

// Length checks if collection is of length expected.
func (assert *Assertion) Length(collection interface{}, expected int, args ...interface{}) bool {
	assert.t.Helper()

	var actual = getLength(collection)

	var message = buildMessage(fmt.Sprintf("expected length: %v, actual: %v", expected, actual), args)

	if expected == actual {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotLength checks if collection is not of length unexpected.
func (assert *Assertion) NotLength(collection interface{}, unexpected int, args ...interface{}) bool {
	assert.t.Helper()

	var actual = getLength(collection)

	var message = buildMessage(fmt.Sprintf("unexpected length: %v, actual: %v", unexpected, actual), args)

	if unexpected != actual {
		return true
	}

	assert.t.Error(message)

	return false
}

// Contains checks if collection contains element.
func (assert *Assertion) Contains(collection interface{}, element interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("collection %v does not contain %v", collection, element), args)

	if contains(collection, element) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotContains checks if collection does not contain element.
func (assert *Assertion) NotContains(collection interface{}, element interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("collection %v contains %v", collection, element), args)

	if !contains(collection, element) {
		return true
	}

	assert.t.Error(message)

	return false
}

// HasKey checks if _map has key as a key.
func (assert *Assertion) HasKey(_map interface{}, key interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("map %v does not contain key %v", _map, key), args)

	if hasKey(_map, key) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotHasKey checks if _map doesn't have key as a key.
func (assert *Assertion) NotHasKey(_map interface{}, key interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("map %v contains key %v", _map, key), args)

	if !hasKey(_map, key) {
		return true
	}

	assert.t.Error(message)

	return false
}

// ContainsPair checks if _map contains key -> value.
func (assert *Assertion) ContainsPair(_map interface{}, key interface{}, value interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("map %v does not contain the pair (%v, %v)", _map, key, value), args)

	if containsPair(_map, key, value) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotContainsPair checks if _map does not contain key -> value.
func (assert *Assertion) NotContainsPair(_map interface{}, key interface{}, value interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("map %v contains the pair (%v, %v)", _map, key, value), args)

	if !containsPair(_map, key, value) {
		return true
	}

	assert.t.Error(message)

	return false
}

// Prefix checks if _string has prefix as prefix.
func (assert *Assertion) Prefix(_string string, prefix string, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("string \"%v\" does not begin with \"%v\"", _string, prefix), args)

	if strings.HasPrefix(_string, prefix) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotPrefix checks if _string does not have prefix as prefix.
func (assert *Assertion) NotPrefix(_string string, prefix string, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("string \"%v\" begins with \"%v\"", _string, prefix), args)

	if !strings.HasPrefix(_string, prefix) {
		return true
	}

	assert.t.Error(message)

	return false
}

// Suffix checks if _string has suffix as suffix.
func (assert *Assertion) Suffix(_string string, suffix string, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("string \"%v\" does not end with \"%v\"", _string, suffix), args)

	if strings.HasSuffix(_string, suffix) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotSuffix checks if _string does not have suffix as suffix.
func (assert *Assertion) NotSuffix(_string string, suffix string, args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("string \"%v\" ends with \"%v\"", _string, suffix), args)

	if !strings.HasSuffix(_string, suffix) {
		return true
	}

	assert.t.Error(message)

	return false
}

// Panics checks if function panics.
func (assert *Assertion) Panics(function func(), args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("function did not panic"), args)

	if panics(function) {
		return true
	}

	assert.t.Error(message)

	return false
}

// NotPanics checks if function does not panic.
func (assert *Assertion) NotPanics(function func(), args ...interface{}) bool {
	assert.t.Helper()

	var message = buildMessage(fmt.Sprintf("function panicked"), args)

	if !panics(function) {
		return true
	}

	assert.t.Error(message)

	return false
}

// PanicValue checks if function panics with the provided value.
func (assert *Assertion) PanicValue(function func(), expected interface{}, args ...interface{}) bool {
	assert.t.Helper()

	var messageValue = buildMessage(fmt.Sprintf("function did not panic with value %v", expected), args)

	var value = panicsWithValue(function, expected)

	if !value {
		assert.t.Error(messageValue)
	} else {
		return true
	}

	return false
}
