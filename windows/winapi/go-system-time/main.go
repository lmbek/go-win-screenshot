package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32DLL   = syscall.NewLazyDLL("kernel32.dll")
	getSystemTime = kernel32DLL.NewProc("GetSystemTime")
)

type systemtime struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

func getCurrentSystemTime() (time.Time, error) {
	var st systemtime
	getSystemTime.Call(uintptr(unsafe.Pointer(&st)))

	currentTime := time.Date(
		int(st.wYear),
		time.Month(st.wMonth),
		int(st.wDay),
		int(st.wHour),
		int(st.wMinute),
		int(st.wSecond),
		0,
		time.UTC,
	)

	return currentTime.Local(), nil
}

func main() {
	currentTime, err := getCurrentSystemTime()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current System Time:", currentTime.Format(time.RFC3339))
}
