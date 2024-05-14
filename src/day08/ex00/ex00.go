package ex00

import (
	"errors"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("error: empty slice")
	}

	if idx < 0 {
		return 0, errors.New("error: negative index")
	}

	if idx > len(arr)-1 {
		return 0, errors.New("error: index is out of bounds")
	}

	ptr := unsafe.Pointer(&arr[0])

	elemPtr := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(idx)*unsafe.Sizeof(arr[0])))

	return *elemPtr, nil
}
