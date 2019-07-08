// ioprobe defines some IO reader and writer probes for testing
// special cases like generating IO errors.

package ioprobe

import "io"

// FailingReader always fails
type FailingReader struct{}

// NewFailingReader returns a reader, which fails all the time.
func NewFailingReader() io.Reader {
	return &FailingReader{}
}

// Read implements read functionality of FailingReader by returning
// io.ErrUnexpectedEOF all the time
func (r *FailingReader) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

// FailingWriter always fails
type FailingWriter struct{}

// NewFailingWriter returns a writer, which fails all the time.
func NewFailingWriter() io.Writer {
	return &FailingWriter{}
}

// Write implements write functionality of FailingWriter by returning
// io.ErrShortWrite all the time
func (w *FailingWriter) Write(p []byte) (int, error) {
	return 0, io.ErrShortWrite
}

// TimeoutReader wraps a reader, but it fails on the Nth read.
type TimeoutReader struct {
	r      io.Reader
	count  int
	failAt int
}

// NewTimeoutReader creates a new TimeoutReader
func NewTimeoutReader(r io.Reader, failAt int) io.Reader {
	return &TimeoutReader{r: r, count: 0, failAt: failAt}
}

func (r *TimeoutReader) Read(p []byte) (int, error) {
	r.count++
	if r.count >= r.failAt {
		return 0, io.ErrUnexpectedEOF
	}
	return r.r.Read(p)
}

// TimeoutWriter wraps a reader, but it fails on the Nth read.
type TimeoutWriter struct {
	w      io.Writer
	count  int
	failAt int
}

// NewTimeoutWriter creates a new TimeoutWriter
func NewTimeoutWriter(w io.Writer, failAt int) io.Writer {
	return &TimeoutWriter{w: w, count: 0, failAt: failAt}
}

func (w *TimeoutWriter) Write(p []byte) (int, error) {
	w.count++
	if w.count >= w.failAt {
		return 0, io.ErrUnexpectedEOF
	}
	return w.w.Write(p)
}
