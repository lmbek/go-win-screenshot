package main

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	messageBoxProc   = user32.NewProc("MessageBoxW")
	getForegroundWin = user32.NewProc("GetForegroundWindow")
)

func main() {
	getForegroundWin.Call()
	messageBoxProc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Hello, Go DLL!"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Greetings"))),
		uintptr(0))
}
