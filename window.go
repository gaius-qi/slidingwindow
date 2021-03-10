package slidingwindow

import (
	"errors"
	"sync"

	"github.com/willf/bitset"
)

// Window the p2p streaming window.
type Window interface {
	// Start is the left position of the window.
	Start() uint

	// Size is the size of the window.
	Size() uint

	// Count is total count of events happened in the window.
	Count() uint

	// Len is the length of bitset.
	Len() uint

	// AddCount increments the accumulated count by n.
	AddCount(n uint)

	// IsFinished returns streaming window execution status.
	IsFinished() bool

	// WindowStatus describe the window status.
	Status() *WindowStatus
}

// WindowStatus describe the window status, which data are available.
type WindowStatus struct {
	mu    sync.RWMutex
	size  uint
	start uint
	bitset.BitSet
}

// NewWindow creates a new window.
func NewWindow(len uint, size uint) (Window, error) {
	if len <= size || len == 0 || size == 0 {
		return nil, errors.New("Len or size values are illegal")
	}

	return &WindowStatus{
		size:   size,
		BitSet: *bitset.New(len),
	}, nil
}

func (w *WindowStatus) Start() uint {
	return w.start
}

func (w *WindowStatus) Size() uint {
	return w.size
}

func (w *WindowStatus) Count() uint {
	return w.BitSet.Count()
}

func (w *WindowStatus) AddCount(n uint) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.Set(n)

	for w.Test(w.start) {
		w.start += 1
	}
}

func (w *WindowStatus) IsFinished() bool {
	return w.All()
}

func (w *WindowStatus) Status() *WindowStatus {
	return w
}
