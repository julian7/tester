package ioprobe_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/julian7/tester"
	"github.com/julian7/tester/ioprobe"
)

func TestFailingReader(t *testing.T) {
	r := ioprobe.NewFailingReader()
	buf := make([]byte, 32)
	_, readerror := r.Read(buf)
	if err := tester.AssertError(errors.New("unexpected EOF"), readerror); err != nil {
		t.Error(err)
	}
}

func TestFailingWriter(t *testing.T) {
	w := ioprobe.NewFailingWriter()
	buf := make([]byte, 32)
	_, writererror := w.Write(buf)
	if err := tester.AssertError(errors.New("short write"), writererror); err != nil {
		t.Error(err)
	}
}

func TestTimeoutReader(t *testing.T) {
	maxIterations := 500
	blockSize := 32
	expectedError := errors.New("unexpected EOF")
	tt := []struct {
		name   string
		failAt int
	}{
		{"one", 1},
		{"ten", 10},
		{"hundred", 100},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, maxIterations*blockSize)
			buf := make([]byte, blockSize)
			origReader := bytes.NewReader(data)
			reader := ioprobe.NewTimeoutReader(origReader, tc.failAt)
			for iteration := 0; iteration < maxIterations; iteration++ {
				_, err := reader.Read(buf)
				if iteration >= tc.failAt {
					if fail := tester.AssertError(expectedError, err); fail != nil {
						t.Error(fail)
					}
				}
			}
		})
	}
}

func TestTimeoutWriter(t *testing.T) {
	maxIterations := 500
	blockSize := 32
	expectedError := errors.New("unexpected EOF")
	tt := []struct {
		name   string
		failAt int
	}{
		{"one", 1},
		{"ten", 10},
		{"hundred", 100},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			buf := make([]byte, blockSize)
			origWriter := bytes.NewBuffer(nil)
			writer := ioprobe.NewTimeoutWriter(origWriter, tc.failAt)
			for iteration := 0; iteration < maxIterations; iteration++ {
				_, err := writer.Write(buf)
				if iteration >= tc.failAt {
					if fail := tester.AssertError(expectedError, err); fail != nil {
						t.Error(fail)
					}
				}
			}
		})
	}
}
