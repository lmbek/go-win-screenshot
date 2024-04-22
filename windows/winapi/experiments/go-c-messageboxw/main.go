package main

/*
#cgo LDFLAGS: -lole32 -luuid
#include <Windows.h>
*/
import "C"

func main() {
	// Initialize the Windows API
	C.CoInitialize(nil)
	defer C.CoUninitialize()

	// Your Go code interacting with the Windows API goes here

	// For example, let's show a message box
	C.MessageBox(nil, C.CString("Hello from Go!"), C.CString("Greeting"), C.MB_OK)
}
