package circular

import (
	"fmt"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.

type Buffer []byte

func NewBuffer(size int) *Buffer {
	return nil
}
func (b *Buffer) ReadByte() (byte, error) {
	fmt.Println(b)
	return 'a', nil
}
func (b *Buffer) WriteByte(c byte) error {
	fmt.Println("byte is :", c)
	fmt.Println(b)
	return nil
}
func (b *Buffer) Overwrite(c byte) {
	fmt.Println("byte is :", c)
}
func (b *Buffer) Reset() {

}

// put buffer in an empty state
//
// We chose the above API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.
