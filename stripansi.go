// package stripansi provides utilities for removing ANSI escape sequences using regular expressions.
package stripansi

import (
	"io"
	"regexp"
	"sync"
)

var re *regexp.Regexp = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

// Regexp returns a copy of the underlying [regexp.Regexp].
func Regexp() *regexp.Regexp {
	return re.Copy()
}

// Bytes removes ANSI escape sequences from the byte slice.
func Bytes(b []byte) []byte {
	return re.ReplaceAll(b, nil)
}

// String removes ANSI escape sequences from the string.
func String(s string) string {
	return re.ReplaceAllString(s, "")
}

// Writer wraps an [io.Writer] and removes ANSI escape sequences from its output.
type Writer struct {
	w  io.Writer
	mu sync.Mutex
}

// NewWriter creates a new [Writer].
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Write removes ANSI escape sequences and writes to the underlying writer.
func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.w.Write(Bytes(p))
}

// Reader wraps an [io.Reader] and removes ANSI escape sequences from its output.
type Reader struct {
	r  io.Reader
	mu sync.Mutex
}

// NewReader creates a new [Reader].
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// Read reads from the underlying reader and removes ANSI escape sequences.
func (r *Reader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	n, err = r.r.Read(p)
	if err != nil {
		return n, err
	}

	cleaned := Bytes(p[:n])
	copy(p, cleaned)

	return len(cleaned), nil
}
