package gow

import (
	. "../winapi"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

type Form struct {
	hWnd HWND
	ViewInfo
}

var hProc HWND
var hBtn HWND
var f *Form

//Register WndClass 
func RegisterClass() {
	var wndProcPtr uintptr = syscall.NewCallback(formProc)
	hInst := GetModuleHandle(nil)
	if hInst == 0 {
		println("error")
		panic("GetModuleHandle")
	}
	hIcon := LoadIcon(0, (*uint16)(unsafe.Pointer(uintptr(IDI_APPLICATION))))
	if hIcon == 0 {
		println("error")
		panic("LoadIcon")
	}
	hCursor := LoadCursor(0, (*uint16)(unsafe.Pointer(uintptr(IDC_ARROW))))
	if hCursor == 0 {
		println("error")
		panic("LoadCursor")
	}
	var wc WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = wndProcPtr
	wc.HInstance = hInst
	wc.HIcon = hIcon
	wc.HCursor = hCursor
	wc.HbrBackground = COLOR_BTNFACE + 1
	wc.LpszClassName = toTEXT(f.name)
	if atom := RegisterClassEx(&wc); atom == 0 {
		println("error")
		panic("RegisterClassEx")
	}
	hProc = HWND(SetWindowLong(f.hWnd, GWL_WNDPROC, int32(wndProcPtr)))
}

//add an button for window
func addButton(x, y, w, h int32, hPar HWND, text string) (btnhWnd HWND) {
	println("add")
	btnhWnd = CreateWindowEx(
		WS_EX_TRANSPARENT,
		toTEXT("button"),
		toTEXT(text),
		WS_CHILD|WS_VISIBLE|BS_PUSHBUTTON,
		x,
		y,
		w,
		h,
		hPar,
		0,
		GetModuleHandle(nil),
		unsafe.Pointer(nil))
	return
}

/*
func Channal(b *chan bool) {
	f.isExit = b
}
*/
func formProc(hWnd HWND, uiMsg uint32, wp, lp uintptr) uintptr {
	println("called", uiMsg)
	switch uiMsg {
	case WM_CREATE:
		println("called ")
		hBtn = addButton(10, 10, 100, 40, hWnd, "确定")
		return 0
	case WM_COMMAND:
		m := CreateMessage()
		m.ShowWithText(strconv.Itoa(int(uiMsg)))
	case WM_DESTROY:
		os.Exit(0)
	}
	return DefWindowProc(hWnd, uiMsg, wp, lp)
}
func (f *Form) init() {
	println("init")
	f.name = "main"
	f.width = 600
	f.height = 400
}
func (f *Form) setClassName(s string) {
	if s != "" {
		f.name = s
	}
}

func CreateMainWindow(name string, strTitle string, v ViewInfo) Form {
	println("create sta")
	if f == nil {
		f = new(Form)
	}
	//x := (GetSystemMetrics(SM_CXSCREEN) - 600) >> 1
	//y := (GetSystemMetrics(SM_CXSCREEN) - 400) >> 1
	/*f.ViewInfo.view = View{
	0,
	0,
	0,
	WS_EX_CLIENTEDGE, WS_OVERLAPPEDWINDOW | WS_CLIPSIBLINGS,
	x, y,
	600, 400,
	name,
	"windowsView",
	nil}
	*/
	f.init()
	f.setClassName(name)
	println(f.name)
	RegisterClass()
	f.hWnd =
		CreateView(f.ViewInfo)
		/*	CreateWindowEx(
				WS_EX_CLIENTEDGE,
				toTEXT(f.name),
				toTEXT(strTitle),
				WS_OVERLAPPEDWINDOW|WS_CLIPSIBLINGS,
				,
				(GetSystemMetrics(SM_CYSCREEN)-f.intHeight)>>1,
				f.intWidth,
				f.intHeight,
				0,
				0,
				GetModuleHandle(nil),
				unsafe.Pointer(nil))
			println("create")
		*/
	println("create end")
	return *f
}
func MainLoop() int {
	var message MSG
	for {
		if GetMessage(&message, 0, 0, 0) == 0 {
			break
		}
		println("mainloop")
		TranslateMessage(&message)
		DispatchMessage(&message)
	}
	return int(message.WParam)
}
func (f *Form) Show() {
	ShowWindow(f.hWnd, SW_SHOW)
}

//dummy
/*package windows

import (
	. "../winapi"
	"os"
	"syscall"
	"unsafe"
)

type Window struct {
	hWnd      HWND
	intWidth  int32
	intHeight int32
}

var hProc HWND

func _TEXT(str string) *uint16 {
	return syscall.StringToUTF16Ptr(str)
}

func formProc(hWnd HWND, uiMsg uint32, wp, lp uintptr) uintptr {
	switch uiMsg {
	case WM_DESTROY:
		os.Exit(0)
	}
	return CallformProc(uintptr(hProc), hWnd, uiMsg, wp, lp)
}
func (w *Window) Init() {
	f.intWidth = 600
	f.intHeight = 400
}
func (w Window) NewWindow() HWND {
	f.Init()
	f.hWnd = CreateWindowEx(
		WS_EX_CLIENTEDGE,
		_TEXT("EDIT"),
		_TEXT("HELLO GUI"),
		WS_OVERLAPPEDWINDOW,
		(GetSystemMetrics(SM_CXSCREEN)-f.intWidth)>>1,
		(GetSystemMetrics(SM_CYSCREEN)-f.intHeight)>>1,
		f.intWidth,
		f.intHeight,
		0,
		0,
		GetModuleHandle(nil),
		unsafe.Pointer(nil))
	wproc := syscall.NewCallback(formProc)
	hProc = HWND(SetWindowLong(f.hWnd, GWL_WNDPROC, int32(wproc)))
	return f.hWnd
}
func (w *Window) Show() {
	var message MSG
	ShowWindow(f.hWnd, SW_SHOW)
	for {
		if GetMessage(&message, 0, 0, 0) == 0 {
			break
		}
		TranslateMessage(&message)
		DispatchMessage(&message)
	}
	os.Exit(int(message.WParam))
}
*/
