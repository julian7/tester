package tester_test

import (
	"errors"
	"testing"

	"github.com/julian7/tester"
)

func TestAssertError(t *testing.T) {
	errA := errors.New("A")
	errB := errors.New("B")
	tt := []struct {
		name     string
		left     error
		right    error
		expected error
	}{
		{"no errors", nil, nil, nil},
		{"unexpected success", errA, nil, errors.New(`unexpected success. Expected error: "A"`)},
		{"different", errA, errB, errors.New("expected error doesn't match.\nExpected: \"A\"\nReceived: \"B\"")},
		{"unexpected error", nil, errA, errors.New(`unexpected error: "A"`)},
		{"matching", errA, errA, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			received := tester.AssertError(tc.left, tc.right)
			if received == nil && tc.expected == nil {
				return
			}
			if received == nil && tc.expected != nil {
				t.Errorf("expected error not received: %v", tc.expected)
				return
			}
			if received != nil && tc.expected == nil {
				t.Errorf("unexpected error received: %v", received)
				return
			}
			if received.Error() != tc.expected.Error() {
				t.Errorf(
					"received error doesn't match with expected:\nExpected: %v\nReceived: %v",
					tc.expected,
					received,
				)
				return
			}
		})
	}
}
