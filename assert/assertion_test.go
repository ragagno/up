package assert

import "testing"

type mockTest struct {
	T *testing.T

	helper bool
	failed bool
}

func (mock *mockTest) checkHelper() {
	mock.T.Helper()

	if !mock.helper {
		mock.T.Error("t.Helper() was not called")
	}

	mock.helper = false
}

func (mock *mockTest) checkSuccess(status bool) {
	mock.T.Helper()

	mock.checkHelper()

	if mock.failed {
		mock.T.Error("t.Error(...interface{}) was called when it needed not to be")
	} else if !status {
		mock.T.Error("false was returned even tho t.Error(...interface{}) was not called")
	}

	mock.failed = false
}

func (mock *mockTest) checkFailure(status bool) {
	mock.T.Helper()

	mock.checkHelper()

	if !mock.failed {
		mock.T.Error("t.Error(...interface{}) was not called when it needed to be")
	} else if status {
		mock.T.Error("true was returned even though t.Error(...interface{}) was called")
	}

	mock.failed = false
}

func (mock *mockTest) Helper() {
	mock.helper = true
}

func (mock *mockTest) Error(...interface{}) {
	mock.failed = true
}

func TestAssertion_True(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.True(true))

	mock.checkFailure(assert.True(false))
}

func TestAssertion_False(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.False(false))

	mock.checkFailure(assert.False(true))
}

func TestAssertion_Nil(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.Nil(nil))
	mock.checkSuccess(assert.Nil((chan struct{})(nil)))
	mock.checkSuccess(assert.Nil((func())(nil)))
	mock.checkSuccess(assert.Nil((interface{})(nil)))
	mock.checkSuccess(assert.Nil((*interface{})(nil)))
	mock.checkSuccess(assert.Nil((map[struct{}]struct{})(nil)))
	mock.checkSuccess(assert.Nil((*struct{})(nil)))
	mock.checkSuccess(assert.Nil(([]struct{})(nil)))

	mock.checkFailure(assert.Nil(struct{}{}))
}

func TestAssertion_NotNil(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.NotNil(struct{}{}))

	mock.checkFailure(assert.NotNil(nil))
	mock.checkFailure(assert.NotNil((chan struct{})(nil)))
	mock.checkFailure(assert.NotNil((func())(nil)))
	mock.checkFailure(assert.NotNil((interface{})(nil)))
	mock.checkFailure(assert.NotNil((*interface{})(nil)))
	mock.checkFailure(assert.NotNil((map[struct{}]struct{})(nil)))
	mock.checkFailure(assert.NotNil((*struct{})(nil)))
	mock.checkFailure(assert.NotNil(([]struct{})(nil)))
}

func TestAssertion_Equals(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.Equals(4, 4))

	mock.checkFailure(assert.Equals(4, 4.0))
	mock.checkFailure(assert.Equals(4, 5))
	mock.checkFailure(assert.Equals(4, "a"))
}

func TestAssertion_NotEquals(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.NotEquals(4, 4.0))
	mock.checkSuccess(assert.NotEquals(4, 5))
	mock.checkSuccess(assert.NotEquals(4, "a"))

	mock.checkFailure(assert.NotEquals(4, 4))
}

func TestAssertion_Length(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var ar = [4]int{1, 2, 4, 8}
	var sl = []int{1, 2, 4, 8}
	var ma = map[int]int{1: 2, 4: 8, 16: 32, 64: 128}
	var st = "abcd"

	mock.checkSuccess(assert.Length(ar, 4))
	mock.checkSuccess(assert.Length(sl, 4))
	mock.checkSuccess(assert.Length(ma, 4))
	mock.checkSuccess(assert.Length(st, 4))

	mock.checkFailure(assert.Length(ar, 50))
	mock.checkFailure(assert.Length(sl, 50))
	mock.checkFailure(assert.Length(ma, 50))
	mock.checkFailure(assert.Length(st, 50))
}

func TestAssertion_NotLength(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var ar = [4]int{1, 2, 4, 8}
	var sl = []int{1, 2, 4, 8}
	var ma = map[int]int{1: 2, 4: 8, 16: 32, 64: 128}
	var st = "abcd"

	mock.checkSuccess(assert.NotLength(ar, 50))
	mock.checkSuccess(assert.NotLength(sl, 50))
	mock.checkSuccess(assert.NotLength(ma, 50))
	mock.checkSuccess(assert.NotLength(st, 50))

	mock.checkFailure(assert.NotLength(ar, 4))
	mock.checkFailure(assert.NotLength(sl, 4))
	mock.checkFailure(assert.NotLength(ma, 4))
	mock.checkFailure(assert.NotLength(st, 4))
}

func TestAssertion_Contains(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var ar = [4]int{1, 2, 4, 8}
	var sl = []int{1, 2, 4, 8}
	var ma = map[int]int{1: 2, 4: 8, 16: 32, 64: 128}
	var st = "abcd"

	mock.checkSuccess(assert.Contains(ar, 2))
	mock.checkSuccess(assert.Contains(sl, 2))
	mock.checkSuccess(assert.Contains(ma, 2))
	mock.checkSuccess(assert.Contains(st, "bc"))
	mock.checkSuccess(assert.Contains(st, byte('b')))
	mock.checkSuccess(assert.Contains(st, 'b'))

	mock.checkFailure(assert.Contains(ar, 3))
	mock.checkFailure(assert.Contains(sl, 3))
	mock.checkFailure(assert.Contains(ma, 3))
	mock.checkFailure(assert.Contains(st, "bd"))
	mock.checkFailure(assert.Contains(st, byte('e')))
	mock.checkFailure(assert.Contains(st, 'e'))
}

func TestAssertion_NotContains(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var ar = [4]int{1, 2, 4, 8}
	var sl = []int{1, 2, 4, 8}
	var ma = map[int]int{1: 2, 4: 8, 16: 32, 64: 128}
	var st = "abcd"

	mock.checkSuccess(assert.NotContains(ar, 3))
	mock.checkSuccess(assert.NotContains(sl, 3))
	mock.checkSuccess(assert.NotContains(ma, 3))
	mock.checkSuccess(assert.NotContains(st, "bd"))
	mock.checkSuccess(assert.NotContains(st, byte('e')))
	mock.checkSuccess(assert.NotContains(st, 'e'))

	mock.checkFailure(assert.NotContains(ar, 2))
	mock.checkFailure(assert.NotContains(sl, 2))
	mock.checkFailure(assert.NotContains(ma, 2))
	mock.checkFailure(assert.NotContains(st, "bc"))
	mock.checkFailure(assert.NotContains(st, byte('b')))
	mock.checkFailure(assert.NotContains(st, 'b'))
}

func TestAssertion_MapHasKey(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _map = map[string]int{
		"foo": 1,
		"bar": 2,
	}

	mock.checkSuccess(assert.HasKey(_map, "foo"))

	mock.checkFailure(assert.HasKey(_map, "unknown"))
}

func TestAssertion_NotMapHasKey(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _map = map[string]int{
		"foo": 1,
		"bar": 2,
	}

	mock.checkSuccess(assert.NotHasKey(_map, "unknown"))

	mock.checkFailure(assert.NotHasKey(_map, "foo"))
}

func TestAssertion_ContainsPair(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _map = map[string]int{
		"foo": 1,
		"bar": 2,
	}

	mock.checkSuccess(assert.ContainsPair(_map, "foo", 1))

	mock.checkFailure(assert.ContainsPair(_map, "bar", 1))
	mock.checkFailure(assert.ContainsPair(_map, "unknown", 8))
}

func TestAssertion_NotContainsPair(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _map = map[string]int{
		"foo": 1,
		"bar": 2,
	}

	mock.checkSuccess(assert.NotContainsPair(_map, "bar", 1))
	mock.checkSuccess(assert.NotContainsPair(_map, "unknown", 8))

	mock.checkFailure(assert.NotContainsPair(_map, "foo", 1))
}

func TestAssertion_Prefix(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _string = "foo bar"

	mock.checkSuccess(assert.Prefix(_string, "foo"))

	mock.checkFailure(assert.Prefix(_string, "unknown"))
}

func TestAssertion_NotPrefix(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _string = "foo bar"

	mock.checkSuccess(assert.NotPrefix(_string, "unknown"))

	mock.checkFailure(assert.NotPrefix(_string, "foo"))
}

func TestAssertion_Suffix(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _string = "foo bar"

	mock.checkSuccess(assert.Suffix(_string, "bar"))

	mock.checkFailure(assert.Suffix(_string, "unknown"))
}

func TestAssertion_NotSuffix(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	var _string = "foo bar"

	mock.checkSuccess(assert.NotSuffix(_string, "unknown"))

	mock.checkFailure(assert.NotSuffix(_string, "bar"))
}

func TestAssertion_Panics(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.Panics(func() { panic("error") }))
	mock.checkSuccess(assert.Panics(func() { panic(nil) }))

	mock.checkFailure(assert.Panics(func() {}))
}

func TestAssertion_NotPanics(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.NotPanics(func() {}))

	mock.checkFailure(assert.NotPanics(func() { panic("error") }))
	mock.checkFailure(assert.NotPanics(func() { panic(nil) }))
}

func TestAssertion_PanicValue(t *testing.T) {
	var mock = &mockTest{T: t}

	var assert = newAssertion(mock)

	mock.checkSuccess(assert.PanicValue(func() { panic("error") }, "error"))
	mock.checkSuccess(assert.PanicValue(func() { panic(nil) }, nil))

	mock.checkFailure(assert.PanicValue(func() {}, "error"))
	mock.checkFailure(assert.PanicValue(func() { panic(nil) }, "error"))
	mock.checkFailure(assert.PanicValue(func() { panic("error") }, nil))
	mock.checkFailure(assert.PanicValue(func() { panic("unknown") }, "error"))
}
