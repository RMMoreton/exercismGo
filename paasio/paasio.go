// Package paasio solves an Exercism challenge.
package paasio

import (
	"io"
	"sync"
)

// Required for testing.
const testVersion = 3

// A readCounter keeps track of how many bytes and read operations have been
// performed on it, and is thread-safe. Read order across multiple threads
// is no specified.
type readCounter struct {
	nBytes int64
	nOps   int
	b      io.Reader
	m      sync.Mutex
}

// A writeCounter keeps track of how many bytes have been written and write
// operations have been performed on it.
type writeCounter struct {
	nBytes int64
	nOps   int
	b      io.Writer
	m      sync.Mutex
}

// A readWriteCounter is both a readCounter and a writeCounter.
type readWriteCounter struct {
	r *readCounter
	w *writeCounter
}

// NewReadCounter returns a new readCounter.
func NewReadCounter(b io.Reader) ReadCounter {
	r := new(readCounter)
	r.b = b
	return r
}

// Read reads up to len(p) bytes from the underlying buffer of the reader.
func (r *readCounter) Read(p []byte) (int, error) {
	r.m.Lock()
	r.nOps++
	r.m.Unlock()
	n, err := r.b.Read(p)
	if n != 0 {
		r.m.Lock()
		r.nBytes += int64(n)
		r.m.Unlock()
	}
	return n, err
}

// ReadCount returns statistics on the number of reads that have been
// performed on r.
func (r *readCounter) ReadCount() (int64, int) {
	// Just to be safe I guess.
	r.m.Lock()
	nB, nO := r.nBytes, r.nOps
	r.m.Unlock()
	return nB, nO
}

// NewWriteCounter returns a new writeCounter.
func NewWriteCounter(b io.Writer) WriteCounter {
	w := new(writeCounter)
	w.b = b
	return w
}

// Write writes len(p) bytes to w's underlying writer, and records how many
// bytes were written.
func (w *writeCounter) Write(p []byte) (int, error) {
	w.m.Lock()
	w.nOps++
	w.m.Unlock()
	n, err := w.b.Write(p)
	if n != 0 {
		w.m.Lock()
		w.nBytes += int64(n)
		w.m.Unlock()
	}
	return n, err
}

// WriteCount returns statistics on the number of reads that have been
// performed on w.
func (w *writeCounter) WriteCount() (int64, int) {
	w.m.Lock()
	nB, nO := w.nBytes, w.nOps
	w.m.Unlock()
	return nB, nO
}

// NewReadWriter returns a new readWriteCounter.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	r := new(readCounter)
	w := new(writeCounter)
	r.b = rw
	w.b = rw
	rwCounter := new(readWriteCounter)
	rwCounter.r = r
	rwCounter.w = w
	return rwCounter
}

// Read reads from the readWriteCounter's reader.
func (rw *readWriteCounter) Read(p []byte) (int, error) {
	return rw.r.Read(p)
}

// Write writes to the readWriteCounter's writer.
func (rw *readWriteCounter) Write(p []byte) (int, error) {
	return rw.w.Write(p)
}

// ReadCount returns rw's read stats.
func (rw *readWriteCounter) ReadCount() (int64, int) {
	return rw.r.ReadCount()
}

// WriteCount returns rw's write stats.
func (rw *readWriteCounter) WriteCount() (int64, int) {
	return rw.w.WriteCount()
}
