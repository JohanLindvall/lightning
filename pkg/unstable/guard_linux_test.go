//go:build linux && (amd64 || arm64)

package unstable

import (
	"testing"

	"golang.org/x/sys/unix"
)

// guardedPage maps three pages with the outer two inaccessible and returns the
// middle one: a slice placed flush against either end of it faults on any read
// past that end, which is how the tests below prove that no assembly body reads
// outside its buffer. The masked tail loads rely on fault suppression for the
// lanes they switch off, and the overlapping tails on reading only inside the
// slice; both are claims about memory that a normal Go allocation, always
// followed by more heap, can never falsify.
func guardedPage(t *testing.T) []byte {
	t.Helper()
	ps := unix.Getpagesize()
	mem, err := unix.Mmap(-1, 0, 3*ps, unix.PROT_NONE, unix.MAP_PRIVATE|unix.MAP_ANON)
	if err != nil {
		t.Skipf("mmap: %v", err)
	}
	t.Cleanup(func() { _ = unix.Munmap(mem) })
	mid := mem[ps : 2*ps : 2*ps]
	if err := unix.Mprotect(mid, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Skipf("mprotect: %v", err)
	}
	return mid
}
