package v1

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonErrorsWrapping(t *testing.T) {
	errors := []struct {
		name           string
		errorFunc      func(string, ...interface{}) error
		validationFunc func(error) bool
	}{
		{"Canceled", Canceled, IsCanceled},
		{"DeadlineExceeded", DeadlineExceeded, IsDeadlineExceeded},
		{"InvalidArgument", InvalidArgument, IsInvalidArgument},
		{"NotFound", NotFound, IsNotFound},
		{"AlreadyExists", AlreadyExists, IsAlreadyExists},
		{"PermissionDenied", PermissionDenied, IsPermissionDenied},
		{"PreconditionFailed", PreconditionFailed, IsPreconditionFailed},
		{"Unauthenticated", Unauthenticated, IsUnauthenticated},
		{"ResourceExhausted", ResourceExhausted, IsResourceExhausted},
		{"Unknown", Unknown, IsUnknown},
		{"Unavailable", Unavailable, IsUnavailable},
		{"Aborted", Aborted, IsAborted},
	}
	for idx, testCase := range errors {
		t.Run(testCase.name, func(t *testing.T) {
			e := testCase.errorFunc("%s error", testCase.name)
			wrapped := fmt.Errorf("Wraps: %w", e)
			wrapped2 := fmt.Errorf("Wraps another one: %w", wrapped)
			for idx2, _ := range errors {
				assert.Equal(t, idx == idx2, errors[idx2].validationFunc(e))
				assert.Equal(t, idx == idx2, errors[idx2].validationFunc(wrapped), "wrapped error %s is not detected as an error", testCase.name)
				assert.Equal(t, idx == idx2, errors[idx2].validationFunc(wrapped2), "wrapped error %s is not detected as an error", testCase.name)
			}
		})
	}
}
