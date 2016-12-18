package main

import (
	"errors"
)

// TODO: do this with sliceExample, refactor all, check errors

// Arrays: a numbered sequence of elements of a single type with a fixed length.
// slices: can never be longer than array but can be smaller. x := make([]float64, 5)
// x := make([]float64, 5, 10) being 10 the capacity of array and 5 the length of slice.
// x := arr[0:5]
// arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5]
// and arr[:] is the same as arr[0:len(arr)].
// a slice is a piece of an array
// [0:5]: inclusive, exclusive

type array struct {
	length        int
	zerothElement *int
}

// this is the array
var array [256]int

// this is the slice
type slice struct {
	length        int
	capacity      int
	zerothElement *int
}

var sliceExample = slice{
	Length:        50,
	Capacity:      256,       //equal to the length of the underlying array, minus the index in the array of the first element of the slice.
	ZerothElement: &array[0], //pointing to the array
}

func isInSlice(s []int, x int) (int, bool) {
	i, pos := 1
	for i <= len(s) && s[i] < x {
		i++
	}
	if i > len(s) || s[i] > x {
		pos = -1
		return pos, false
	}
	return pos, true
}

func isEmpty(s []int) bool {
	if s == nil || len(s) == nil {
		return true
	}
	return false
}

//reads each element
func read(s []int) (int, error) {
	if isEmpty {
		e := errors.New("Empty slice")
	} else {
		for _, val := range s {
			v := val
			e = nil
		}
	}
	return v, e
}

//writes elements
func write(n, pos int, s []int) ([]int, error) {
	if isEmpty {
		e := errors.New("Empty slice")
	} else {
		if s[pos] == 0 || s[pos] == nil {
			s[pos] = n
			e = nil
		}
	}
	return s, e
}

// how to make something optional? how to define the type?
func makeSlice(types string, length, capacity int) []int {
	if length < capacity {
		var slice2 []int // how to pass the type here?
		slice2 = slice2[:length]
		return slice2
	}
	return nil
}

func appendToSlice(original []int, position, value int, order bool) ([]int, error) {
	if !order {
		// make slice of original length plus 1
		target := make([]int, len(original)+1)

		// copy the contents of original[0:position] to target
		copy(target, original[:position])

		// replace but is in position with value
		target[position] = value

		// copy contents of original[position:len(arr)] to target[position+1:len(target)]
		copy(target[position+1:], original[position:])
	} else {
		err := errors.New("There is no space")
		return err, _
	}
	return target, nil
}

func deleteInSlice(original []int, d int, order bool) ([]int, error) {
	var target []int
	if !isEmpty {
		if order {
			target = original[:d+copy(original[d:], original[d+1:])]
			return _, target
		} else if !order {
			target = target[:len(target)-1]
			return _, target
		} else {
			return errors.New("element not found"), original
		}
	}
	return target, errors.New("slice is Empty")
}

func updateInSlice(original []int, u, x int, order bool) ([]int, error) {
	if !isEmpty {
		if !order {
			for _, i := range original {
				if original[i] == x {
					e := nil
					original[i] = u
				} else {
					e = errors.New("Element not found"), original

				}
			}
		}
	}
	return original, e
}

//NEW vs MAKE

//Let's talk about new first. It's a built-in function that allocates memory,
// but unlike its namesakes in some other languages it does not initialize the memory,
// it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T
// and returns its address, a value of type *T. In Go terminology, it returns a pointer
// to a newly allocated zero value of type T.

// The built-in function make(T, args) serves a purpose different from new(T). It creates
// slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T
// (not *T). The reason for the distinction is that these three types represent, under the covers
// , references to data structures that must be initialized before use. A slice, for example, is a
// three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity,
// and until those items are initialized, the slice is nil. For slices, maps, and channels, make initializes
// the internal data structure and prepares the value for use
