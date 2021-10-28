package calculator

type FloatStack []float64

// IsEmpty: check if FloatStack is empty
func (s *FloatStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the FloatStack
func (s *FloatStack) Push(str float64) {
	*s = append(*s, str) // Simply append the new value to the end of the FloatStack
}

// Remove and return top element of FloatStack. Return false if FloatStack is empty.
func (s *FloatStack) Pop() (float64, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the FloatStack by slicing it off.
		return element, true
	}
}
func (s *FloatStack) Top() (float64, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}
