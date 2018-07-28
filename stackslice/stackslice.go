package stackslice

// StackSlice is to store stack
type StackSlice struct {
	items []int
}

// NewStackSlice is to instantiate NewStackSlice
func NewStackSlice() *StackSlice {
	return &StackSlice{}
}

func (stackslice *StackSlice) push(item int) {
	stackslice.items = append(stackslice.items, item)
}

func (stackslice *StackSlice) pop() int {
	lastIndex := len(stackslice.items) - 1
	lastItem := stackslice.items[lastIndex]
	stackslice.items = append(stackslice.items[:0], stackslice.items[lastIndex:]...)

	return lastItem
}

func (stackslice *StackSlice) isEmpty() bool {
	return len(stackslice.items) > 0
}
