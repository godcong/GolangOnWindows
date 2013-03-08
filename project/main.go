package main

import (
	G "./gow"
)

func main() {
	v := G.ViewInfo{}
	g := G.CreateMainWindow("", "测试", v)
	g.Show()
	G.MainLoop()

}
