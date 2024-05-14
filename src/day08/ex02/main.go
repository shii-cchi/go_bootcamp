package main

// #cgo LDFLAGS: -framework Cocoa
// #include "window/application.h"
// #include "window/window.h"
import "C"

func main() {
	C.InitApplication()

	window := C.Window_Create(0, 0, 300, 200, C.CString("School 21"))
	C.Window_MakeKeyAndOrderFront(window)

	C.RunApplication()
}
