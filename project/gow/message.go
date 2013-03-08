package gow

import (
	. "../winapi"
	"strconv"
	"syscall"
)

type Message struct {
	hWnd                 HWND
	strTitle, strMessage string
	uiStyle              uint32
}

const (
	MSG_OK                uint32 = MB_OK
	MSG_OKCANCEL          uint32 = MB_OKCANCEL
	MSG_ABORTRETRYIGNORE  uint32 = MB_ABORTRETRYIGNORE
	MSG_YESNOCANCEL       uint32 = MB_YESNOCANCEL
	MSG_YESNO             uint32 = MB_YESNO
	MSG_RETRYCANCEL       uint32 = MB_RETRYCANCEL
	MSG_CANCELTRYCONTINUE uint32 = MB_CANCELTRYCONTINUE
	MSG_ICONHAND          uint32 = MB_ICONHAND
	MSG_ICONQUESTION      uint32 = MB_ICONQUESTION
	MSG_ICONEXCLAMATION   uint32 = MB_ICONEXCLAMATION
	MSG_ICONASTERISK      uint32 = MB_ICONASTERISK
	MSG_USERICON          uint32 = MB_USERICON
	MSG_ICONWARNING       uint32 = MB_ICONWARNING
	MSG_ICONERROR         uint32 = MB_ICONERROR
	MSG_ICONINFORMATION   uint32 = MB_ICONINFORMATION
	MSG_ICONSTOP          uint32 = MB_ICONSTOP
	MSG_DEFBUTTON1        uint32 = MB_DEFBUTTON1
	MSG_DEFBUTTON2        uint32 = MB_DEFBUTTON2
	MSG_DEFBUTTON3        uint32 = MB_DEFBUTTON3
	MSG_DEFBUTTON4        uint32 = MB_DEFBUTTON4
)

var m Message

func init() {
	m.hWnd = 0
	m.strTitle = "Message"
	m.strMessage = ""
	m.uiStyle = MSG_OK
}

func (m Message) String() string {
	return "address:" +
		"title:" + m.strTitle + "\n" +
		"message:" + m.strMessage + "\n" +
		"style:" + strconv.Itoa(int(m.uiStyle)) + "\n"
}

func CreateMessage() Message {
	return m
}

func (m Message) CreateMessageCopy() Message {
	return m
}

func (m *Message) SetMessageInfo(hWnd HWND, message, title string, style uint32) {
	m.strTitle = title
	m.strMessage = message
	m.uiStyle = style
	m.hWnd = hWnd
}

func (m *Message) SetHandle(hWnd HWND) {
	m.hWnd = hWnd
}

func (m *Message) SetTitle(title string) {
	m.strTitle = title
}

func (m *Message) SetMessage(message string) {
	m.strMessage = message
}

func (m *Message) SetStyle(style uint32) {
	m.uiStyle = style
}

func (m *Message) ShowMessage(hWnd HWND, message, title string, style uint32) (result int32) {
	m.SetMessageInfo(hWnd, message, title, style)
	result = m.Show()
	return
}

func (m *Message) Show() (result int32) {
	strMsg := syscall.StringToUTF16Ptr(m.strMessage)
	strTit := syscall.StringToUTF16Ptr(m.strTitle)
	result = int32(MessageBox(m.hWnd, strMsg, strTit, m.uiStyle))
	return
}

func (m *Message) ShowWithText(arg ...interface{}) (result int32) {
	var s string
	for _, tmp := range arg {
		switch tmp.(type) {
		case int32, int:
			s += strconv.Itoa(tmp.(int))
		case string:
			s += tmp.(string)
		default:
		}
	}
	m.strMessage = s
	result = m.Show()
	return
}
