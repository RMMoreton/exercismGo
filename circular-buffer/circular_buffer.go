// Package circular solves an Exercism challenge.
package circular

import (
	"errors"
)

// testVersion is required for testing.
const testVersion = 3

// Buffer is a circular buffer.
type Buffer struct {
	buf       []byte
	size      int
	valid     int
	readHead  int
	writeHead int
}

// NewBuffer returns a new buffer of size n.
func NewBuffer(n int) *Buffer {
	return &Buffer{
		size:      n,
		buf:       make([]byte, n),
		valid:     0,
		readHead:  0,
		writeHead: 0,
	}
}

// ReadByte reads a byte from b, and returns that
// byte. In case of an error, err is non-nil.
func (b *Buffer) ReadByte() (byte, error) {
	if b.valid <= 0 {
		return 0, errors.New("no readable bytes avaliable")
	}
	res := b.buf[b.readHead]
	b.readHead = (b.readHead + 1) % b.size
	b.valid--
	return res, nil
}

// WriteByte writes a byte to b if possible. If not,
// then WriteByte returns an error.
func (b *Buffer) WriteByte(c byte) error {
	if b.valid == b.size {
		return errors.New("no empty slots for a write")
	}
	b.buf[b.writeHead] = c
	b.writeHead = (b.writeHead + 1) % b.size
	b.valid++
	return nil
}

// Overwrite writes a byte to b whether the byte at writeHead
// has been read or not.
func (b *Buffer) Overwrite(c byte) {
	b.buf[b.writeHead] = c
	b.writeHead = (b.writeHead + 1) % b.size
	if b.valid != b.size {
		b.valid++
	} else {
		b.readHead = (b.readHead + 1) % b.size
	}
}

// Reset resets the buffer to a clean state.
func (b *Buffer) Reset() {
	b.valid = 0
	b.readHead = 0
	b.writeHead = 0
}
