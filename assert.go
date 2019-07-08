package tester

import "fmt"

// AssertError asserts expected error matches with received.
//
// Usage:
//
// _, someerr := SomeMethod(...)
// if err := tester.AssertError(tc.err, someerr); err != nil {
//     t.Error(err)
// }
func AssertError(expected, received error) error {
	if received == nil {
		if expected == nil {
			return nil
		}
		return fmt.Errorf("unexpected success. Expected error: %q", expected)
	}
	if expected == nil {
		return fmt.Errorf("unexpected error: %q", received.Error())
	}
	if expected.Error() != received.Error() {
		return fmt.Errorf(
			"expected error doesn't match.\nExpected: %q\nReceived: %q",
			expected.Error(),
			received.Error(),
		)
	}
	return nil
}
