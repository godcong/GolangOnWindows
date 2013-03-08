package gow

import (
	. "../winapi"
	"fmt"
	US "unsafe"
)

type View interface {
	SetVisible(bool)
	IsVisible() bool
}
type ViewStatus struct {
	x      int32
	y      int32
	width  int32
	height int32
}

type ViewInfo struct {
	view       View
	hWndParent HWND
	hMenu      HMENU
	hlnstnce   HINSTANCE
	exStyle    uint32
	style      uint32
	name       string
	title      string
	ipParam    *uintptr
	ViewStatus
}

func (v *ViewInfo) init() {
	println("view init sta")
	v.exStyle = WS_EX_CLIENTEDGE
	v.width = 600
	v.height = 400
	//	v.x = (GetSystemMetrics(SM_CXSCREEN) - 600) >> 1
	//	v.y = (GetSystemMetrics(SM_CXSCREEN) - 400) >> 1
	v.x = 0
	v.y = 0
	v.exStyle = WS_EX_CLIENTEDGE
	v.style = (WS_OVERLAPPEDWINDOW | WS_CLIPSIBLINGS)
	v.name = "main"
	v.title = "main"
}

func CreateView(v ViewInfo) HWND {
	fmt.Println(v)
	v.init()
	hHandle := CreateWindowEx(
		v.exStyle,
		toTEXT(v.name),
		toTEXT(v.title),
		v.style,
		v.x,
		v.y,
		v.width,
		v.height,
		v.hWndParent,
		v.hMenu,
		v.hlnstnce,
		US.Pointer(v.ipParam))
	return hHandle
}

/*
   函数功能：该函数创建一个具有扩展风格的重叠式窗口、弹出式窗口或子窗口，其他与     CreateWindow函数相同。关于创建窗口和其他参数的内容，请参看CreateWindowEx。

   函数原型：HWND CreateWindowEx（DWORD dwExStle，LPCTSTR IpClassName，LPCTSTR lpWindowName，DWORD dwStyle，int x，int y,int nWidth，int nHeight，HWND hWndParent，HMENUhMenu，HANDLE hlnstance，LPVOIDlpParam）；
   参数：
   ：指定窗口的扩展风格。该参数可以是下列值：
   WS_EX_ACCEPTFILES：指定以该风格创建的窗口接受一个拖拽文件。

   WS_EX_APPWINDOW：当窗口可见时，将一个顶层窗口放置到任务条上。

   WS_EX_CLIENTEDGE：指定窗口有一个带阴影的边界。

   WS_EX_CONTEXTHELP：在窗口的标题条包含一个问号标志。当用户点击了问号时，鼠标光标变为一个问号的指针、如果点击了一个子窗口，则子窗日接收到WM_HELP消息。子窗口应该将这个消息传递给父窗口过程，父窗口再通过HELP_WM_HELP命令调用WinHelp函数。这个Help应用程序显示一个包含子窗口帮助信息的弹出式窗口。 WS_EX_CONTEXTHELP不能与WS_MAXIMIZEBOX和WS_MINIMIZEBOX同时使用。

   WS_EX_CONTROLPARENT：允许用户使用Tab键在窗口的子窗口间搜索。

   WS_EX_DLGMODALFRAME：创建一个带双边的窗口；该窗口可以在dwStyle中指定WS_CAPTION风格来创建一个标题栏。

   WS_EX_LEFT：窗口具有左对齐属性，这是缺省设置的。

   WS_EX_LEFTSCROLLBAR：如果外壳语言是如Hebrew，Arabic，或其他支持reading order alignment的语言，则标题条（如果存在）则在客户区的左部分。若是其他语言，在该风格被忽略并且不作为错误处理。

   WS_EX_LTRREADING：窗口文本以LEFT到RIGHT（自左向右）属性的顺序显示。这是缺省设置的。

   WS_EX_MDICHILD：创建一个MD子窗口。

   WS_EX_NOPATARENTNOTIFY：指明以这个风格创建的窗口在被创建和销毁时不向父窗口发送WM_PARENTNOTFY消息。

   WS_EX_OVERLAPPED：WS_EX_CLIENTEDGE和WS_EX_WINDOWEDGE的组合。

   WS_EX_PALETTEWINDOW：WS_EX_WINDOWEDGE, WS_EX_TOOLWINDOW和WS_WX_TOPMOST风格的组合WS_EX_RIGHT:窗口具有普通的右对齐属性，这依赖于窗口类。只有在外壳语言是如Hebrew,Arabic或其他支持读顺序对齐（reading order alignment）的语言时该风格才有效，否则，忽略该标志并且不作为错误处理。

   WS_EX_RIGHTSCROLLBAR：垂直滚动条在窗口的右边界。这是缺省设置的。

   WS_EX_RTLREADING：如果外壳语言是如Hebrew，Arabic，或其他支持读顺序对齐（reading order alignment）的语言，则窗口文本是一自左向右）RIGHT到LEFT顺序的读出顺序。若是其他语言，在该风格被忽略并且不作为错误处理。

   WS_EX_STATICEDGE：为不接受用户输入的项创建一个3一维边界风格

   WS_EX_TOOLWIDOW：创建工具窗口，即窗口是一个游动的工具条。工具窗口的标题条比一般窗口的标题条短，并且窗口标题以小字体显示。工具窗口不在任务栏里显示，当用户按下alt＋Tab键时工具窗口不在对话框里显示。如果工具窗口有一个系统菜单，它的图标也不会显示在标题栏里，但是，可以通过点击鼠标右键或Alt＋Space来显示菜单。

   WS_EX_TOPMOST：指明以该风格创建的窗口应放置在所有非最高层窗口的上面并且停留在其L，即使窗口未被激活。使用函数SetWindowPos来设置和移去这个风格。

   WS_EX_TRANSPARENT：指定以这个风格创建的窗口在窗口下的同属窗口已重画时，该窗口才可以重画。

   由于其下的同属富日已被重画，该窗口是透明的。

   IpClassName:指向一个空结束的字符串或整型数atom。如果该参数是一个整型量，它是由此前调用theGlobaIAddAtom函数产生的全局量。这个小于OxCOOO的16位数必须是IpClassName参数字的低16位，该参数的高位必须是O。

   如果lpClassName是一个字符串，它指定了窗口的类名。这个类名可以是任何用函数RegisterClassEx注册的类名，或是任何预定义的控制类名。请看说明部分的列表。

   lpWindowName:指向一个指定窗口名的空结束的字符串指针。

   如果窗口风格指定了标题条，由lpWindowName指向的窗口标题将显示在标题条上。当使用CreateWindow

   函数来创建控制例如按钮，选择框和静态控制时，可使用lpWindowName来指定控制文本。

   dwStyle:指定创建窗口的风格。该参数可以是下列窗口风格的组合再加上说明部分的控制风格。

   x：参见CreateWindow。

   y：参见CreateWindow。

   nWidth:CreateWindow。

   nHeigth：参见CreateWindow。

   hWndParent：参见CreateWindow。

   hMenu：参见CreateWindow。

   hlnstance：参见CreateWindow。

   lpParam:参见CreateWindow。

   返回值：参见CreateWindow。

   备注:参见CreateWindow。

   速查：Windows NT：3.1以上版本；Windows：95以上版本；Windows CE：1.0以上版本；头文件：winuser.h；库文件：USer32.lib;Unicode：在Windows NT上实现为Unicode和ANSI两种版本。
*/
