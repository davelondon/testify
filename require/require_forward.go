/*
* CODE GENERATED AUTOMATICALLY WITH github.com/davelondon/ktest/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package require

import (
	sync "sync"
	time "time"

	assert "github.com/davelondon/ktest/assert"
)

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	Condition(a.t, comp, msgAndArgs...)
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
//    a.Contains(["Hello", "World"], "World", "But ["Hello", "World"] does contain 'World'")
//    a.Contains({"Hello": "World"}, "Hello", "But {'Hello': 'World'} does contain 'Hello'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	Contains(a.t, s, contains, msgAndArgs...)
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Empty(obj)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) {
	Empty(a.t, object, msgAndArgs...)
}

// Equal asserts that two objects are equal.
//
//    a.Equal(123, 123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	Equal(a.t, expected, actual, msgAndArgs...)
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err, "An error was expected") {
// 	   assert.Equal(t, err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	EqualError(a.t, theError, errString, msgAndArgs...)
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValues(uint32(123), int32(123), "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	EqualValues(a.t, expected, actual, msgAndArgs...)
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Error(err, "An error was expected") {
// 	   assert.Equal(t, err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) {
	Error(a.t, err, msgAndArgs...)
}

// Exactly asserts that two objects are equal is value and type.
//
//    a.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	Exactly(a.t, expected, actual, msgAndArgs...)
}

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) {
	Fail(a.t, failureMessage, msgAndArgs...)
}

// FailNow fails test
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) {
	FailNow(a.t, failureMessage, msgAndArgs...)
}

// False asserts that the specified value is false.
//
//    a.False(myBool, "myBool should be false")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) {
	False(a.t, value, msgAndArgs...)
}

// HasError works with the kerr package to test for a specific error on the error stack
func (a *Assertions) HasError(theError error, expectedId string, msgAndArgs ...interface{}) {
	HasError(a.t, theError, expectedId, msgAndArgs...)
}

// HasError works with the kerr package to test for a specific error on the error stack, but the error is expected to come from an external package so tooling shouldn't expect to find the error thrown.
func (a *Assertions) HasErrorExternal(theError error, expectedId string, msgAndArgs ...interface{}) {
	HasErrorExternal(a.t, theError, expectedId, msgAndArgs...)
}

// Implements asserts that an object is implemented by the specified interface.
//
//    a.Implements((*MyInterface)(nil), new(MyObject), "MyObject")
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	Implements(a.t, interfaceObject, object, msgAndArgs...)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 a.InDelta(math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	InDelta(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...)
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...)
}

// IsError works with the kerr package to test for a specific error
func (a *Assertions) IsError(theError error, expectedId string, msgAndArgs ...interface{}) {
	IsError(a.t, theError, expectedId, msgAndArgs...)
}

// IsError works with the kerr package to test for a specific error, but the error is expected to come from an external package so tooling shouldn't expect to find the error thrown.
func (a *Assertions) IsErrorExternal(theError error, expectedId string, msgAndArgs ...interface{}) {
	IsErrorExternal(a.t, theError, expectedId, msgAndArgs...)
}

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	IsType(a.t, expectedType, object, msgAndArgs...)
}

// JSONEq asserts that two JSON strings are equivalent.
//
//  a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	JSONEq(a.t, expected, actual, msgAndArgs...)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    a.Len(mySlice, 3, "The size of slice is not 3")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	Len(a.t, object, length, msgAndArgs...)
}

// Nil asserts that the specified object is nil.
//
//    a.Nil(err, "err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) {
	Nil(a.t, object, msgAndArgs...)
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoError(err) {
// 	   assert.Equal(t, actualObj, expectedObj)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) {
	NoError(a.t, err, msgAndArgs...)
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
//    a.NotContains(["Hello", "World"], "Earth", "But ['Hello', 'World'] does NOT contain 'Earth'")
//    a.NotContains({"Hello": "World"}, "Earth", "But {'Hello': 'World'} does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	NotContains(a.t, s, contains, msgAndArgs...)
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmpty(obj) {
//    assert.Equal(t, "two", obj[1])
//  }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	NotEmpty(a.t, object, msgAndArgs...)
}

// NotEqual asserts that the specified values are NOT equal.
//
//    a.NotEqual(obj1, obj2, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	NotEqual(a.t, expected, actual, msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
//
//    a.NotNil(err, "err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) {
	NotNil(a.t, object, msgAndArgs...)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanics(func(){
//     RemainCalm()
//   }, "Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	NotPanics(a.t, f, msgAndArgs...)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  a.NotRegexp("^start", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	NotRegexp(a.t, rx, str, msgAndArgs...)
}

// NotZero asserts that i is not the zero value for its type and returns the truth.
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) {
	NotZero(a.t, i, msgAndArgs...)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panics(func(){
//     GoCrazy()
//   }, "Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	Panics(a.t, f, msgAndArgs...)
}

// Regexp asserts that a specified regexp matches a string.
//
//  a.Regexp(regexp.MustCompile("start"), "it's starting")
//  a.Regexp("start...$", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	Regexp(a.t, rx, str, msgAndArgs...)
}

// True asserts that the specified value is true.
//
//    a.True(myBool, "myBool should be true")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) {
	True(a.t, value, msgAndArgs...)
}

//
func (a *Assertions) WaitFor(c chan struct{}, shouldBeOpen bool, msgAndArgs ...interface{}) {
	WaitFor(a.t, c, shouldBeOpen, msgAndArgs...)
}

//
func (a *Assertions) WaitForError(c chan error, errorId string, msgAndArgs ...interface{}) {
	WaitForError(a.t, c, errorId, msgAndArgs...)
}

//
func (a *Assertions) WaitForGroup(wg *sync.WaitGroup, msgAndArgs ...interface{}) {
	WaitForGroup(a.t, wg, msgAndArgs...)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   a.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
}

// Zero asserts that i is the zero value for its type and returns the truth.
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) {
	Zero(a.t, i, msgAndArgs...)
}
