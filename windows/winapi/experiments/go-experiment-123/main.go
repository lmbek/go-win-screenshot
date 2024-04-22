package main

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	shell32          = syscall.MustLoadDLL("shell32.dll")
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	getModuleHandleW = kernel32.MustFindProc("GetModuleHandleW")
	loadIconW        = user32.MustFindProc("LoadIconW")
	loadMenuW        = user32.MustFindProc("LoadMenuW")
	shellNotifyIcon  = user32.MustFindProc("Shell_NotifyIconW")
	getMessageW      = user32.MustFindProc("GetMessageW")
	translateMessage = user32.MustFindProc("TranslateMessage")
	dispatchMessage  = user32.MustFindProc("DispatchMessageW")
	destroyMenu      = user32.MustFindProc("DestroyMenu")
	destroyIcon      = user32.MustFindProc("DestroyIcon")
)

const (
	WM_QUIT          = 0x0012
	WM_DESTROY       = 0x0002
	WM_USER          = 0x0400
	WM_APP           = 0x8000
	WM_SYSTRAY       = WM_APP + 1
	WM_NOTIFY        = 0x004E
	ID_TRAY_APP_ICON = 1
)

var hIcon, hMenu, hNotify syscall.Handle

func init() {
	// Load your icon and menu resources here
	hIcon = loadIcon("your_icon.ico")
	hMenu = loadMenu("your_menu.json")
}

func main() {
	hModule, _, _ := getModuleHandleW.Call(0)
	msg := &syscall.Msg{}
	for {
		getMessageW.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0)
		translateMessage.Call(uintptr(unsafe.Pointer(msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(msg)))
		if msg.Message == WM_QUIT {
			break
		}
	}
}

func loadIcon(iconPath string) syscall.Handle {
	ret, _, _ := loadIconW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(iconPath))))
	return syscall.Handle(ret)
}

func loadMenu(menuPath string) syscall.Handle {
	// Load your menu from a resource or file here
	// For simplicity, you can create a JSON file to define your menu
	// and parse it to build the menu dynamically
	return syscall.Handle(0)
}

func showSysTray() {
	nid := NOTIFYICONDATA{
		CbSize:           uint32(unsafe.Sizeof(NOTIFYICONDATA{})),
		HWnd:             syscall.Handle(0),
		UID:              ID_TRAY_APP_ICON,
		UCallbackMessage: WM_SYSTRAY,
		HIcon:            hIcon,
		UTip:             [128]uint16{'M', 'y', ' ', 'A', 'p', 'p', 0},
	}
	shellNotifyIcon.Call(uintptr(0), uintptr(unsafe.Pointer(&nid)))
}

func hideSysTray() {
	nid := NOTIFYICONDATA{
		CbSize: uint32(unsafe.Sizeof(NOTIFYICONDATA{})),
		HWnd:   syscall.Handle(0),
		UID:    ID_TRAY_APP_ICON,
	}
	shellNotifyIcon.Call(uintptr(2), uintptr(unsafe.Pointer(&nid)))
	destroyIcon.Call(uintptr(hIcon))
	destroyMenu.Call(uintptr(hMenu))
}
