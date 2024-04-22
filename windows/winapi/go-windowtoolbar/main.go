package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32DLL        = syscall.NewLazyDLL("user32.dll")
	kernel32DLL      = syscall.NewLazyDLL("kernel32.dll")
	createWindowEx   = user32DLL.NewProc("CreateWindowExW")
	defWindowProc    = user32DLL.NewProc("DefWindowProcW")
	dispatchMessage  = user32DLL.NewProc("DispatchMessageW")
	getMessage       = user32DLL.NewProc("GetMessageW")
	registerClassEx  = user32DLL.NewProc("RegisterClassExW")
	translateMessage = user32DLL.NewProc("TranslateMessage")
	destroyWindow    = user32DLL.NewProc("DestroyWindow")
	postQuitMessage  = user32DLL.NewProc("PostQuitMessage")
	getModuleHandleW = kernel32DLL.NewProc("GetModuleHandleW")

	loadIconW    = user32DLL.NewProc("LoadIconW")
	loadCursorW  = user32DLL.NewProc("LoadCursorW")
	showWindow   = user32DLL.NewProc("ShowWindow")
	updateWindow = user32DLL.NewProc("UpdateWindow")
)

const (
	cSW_SHOWNORMAL       = 1
	cWM_DESTROY          = 0x0002
	cWM_COMMAND          = 0x0111
	cBN_CLICKED          = 0x00B
	cWS_OVERLAPPED       = 0x00000000
	cWS_CAPTION          = 0x00C00000
	cWS_SYSMENU          = 0x00080000
	cWS_MINIMIZEBOX      = 0x00020000
	cWS_MAXIMIZEBOX      = 0x00010000
	cWS_OVERLAPPEDWINDOW = cWS_OVERLAPPED | cWS_CAPTION | cWS_SYSMENU | cWS_MINIMIZEBOX | cWS_MAXIMIZEBOX
	cIDI_APPLICATION     = 32512
	COLOR_BTNFACE        = 15
	cWS_TABSTOP          = 0x00010000
	cWS_VISIBLE          = 0x10000000
	cBS_DEFPUSHBUTTON    = 0x00000001
	cIDC_ARROW           = 32512
)

type HWND uintptr
type POINT struct {
	X, Y int32
}

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type wndClassEx struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}

type createStruct struct {
	CreateParams uintptr
	Instance     uintptr
}

func HIWORD(n uint32) uint16 {
	return uint16(n >> 16)
}

func main() {
	hInstance := getModuleHandle()

	className := syscall.StringToUTF16Ptr("SimpleWindowClass")
	wndClass := wndClassEx{
		Size:       uint32(unsafe.Sizeof(wndClassEx{})),
		WndProc:    syscall.NewCallback(windowProc),
		Instance:   hInstance,
		Icon:       loadIcon(0, cIDI_APPLICATION),
		Cursor:     loadCursor(0, cIDC_ARROW),
		Background: COLOR_BTNFACE + 1,
		ClassName:  className,
	}
	registerClassEx.Call(uintptr(unsafe.Pointer(&wndClass)))

	hwnd, _, _ := createWindowEx.Call(
		0, // Add WS_EX_TOOLWINDOW here
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Simple Window"))),
		cWS_OVERLAPPEDWINDOW,
		100, 100, 400, 300,
		0, 0, hInstance, 0)

	showWindow.Call(hwnd, cSW_SHOWNORMAL)
	updateWindow.Call(hwnd)

	var msg MSG
	for {
		var result uintptr
		result, _, err := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if result == 0 {
			break
		}
		if result == ^uintptr(0) && err != nil {
			break
		}

		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}

	destroyWindow.Call(hwnd)
	postQuitMessage.Call(0)
}

func windowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case cWM_DESTROY:
		postQuitMessage.Call(0)
		return 0
	case cWM_COMMAND:
		if lParam == 0 && HIWORD(uint32(wParam)) == cBN_CLICKED {
			fmt.Println("Button clicked!")
		}
	}
	result, _, _ := defWindowProc.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return result
}

func getModuleHandle() uintptr {
	hInstance, _, _ := getModuleHandleW.Call(0)
	return hInstance
}

func loadIcon(instance, iconName uintptr) uintptr {
	hIcon, _, _ := loadIconW.Call(instance, iconName)
	return hIcon
}

func loadCursor(instance, cursorName uintptr) uintptr {
	hCursor, _, _ := loadCursorW.Call(instance, cursorName)
	return hCursor
}
