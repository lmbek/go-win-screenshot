package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	MB_OK           = 0x00000000
	MB_OKCANCEL     = 0x00000001
	MB_YESNO        = 0x00000004
	MB_SYSTEMMODAL  = 0x00001000
	MB_ICONQUESTION = 0x00000020
	MB_ICONWARNING  = 0x00000030
	MB_ICONASTERISK = 0x00000040
)

func main() {
	// Load the DLL file
	dllHandle, err := syscall.LoadLibrary("user32.dll")
	if err != nil {
		fmt.Println("Failed to load DLL:", err)
		return
	}
	defer syscall.FreeLibrary(dllHandle)

	// Get the address of the MessageBoxW function
	messageBoxProcAddr, err := syscall.GetProcAddress(dllHandle, "MessageBoxW")
	if err != nil {
		fmt.Println("Failed to find function:", err)
		return
	}

	lpCaption, _ := syscall.UTF16PtrFromString("This is a custom warning")
	lpText, _ := syscall.UTF16PtrFromString("As you are aware\n\nThis is a native solution\n\nThank you for trying it")

	// Call the MessageBoxW function
	_, _, _ = syscall.SyscallN(messageBoxProcAddr,
		0,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		MB_OK|MB_ICONWARNING, // Let the window TOPMOST.
	)
}
