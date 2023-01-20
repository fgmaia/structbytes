package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// type countEntry struct {
// uid uint64
// key []byte
// }

type countEntry []byte

func countEntrySize(key []byte) int {
	return 8 /*uint64, 8bytes*8bits=64bits*/ + 4 /*memory alignment?*/ + len(key)
}
func marshalCountEntry(dst, key []byte, uid uint64) {
	binary.BigEndian.PutUint64(dst[0:8], uid)

	binary.BigEndian.PutUint32(dst[8:12], uint32(len(key)))
	n := copy(dst[12:], key)
	fmt.Println(n)
	//x.AssertTrue(len(dst) == n+12)
}

func (ci countEntry) Uid() uint64 {
	return binary.BigEndian.Uint64(ci[0:8])
}

func (ci countEntry) Key() []byte {
	sz := binary.BigEndian.Uint32(ci[8:12])
	return ci[12 : 12+sz]
}

func (ci countEntry) less(oe countEntry) bool {
	lk, rk := ci.Key(), oe.Key()
	if cmp := bytes.Compare(lk, rk); cmp != 0 {
		return cmp < 0
	}
	return ci.Uid() < oe.Uid()
}

func main() {

}
