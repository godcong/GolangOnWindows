package main

import (
	"fmt"
)

func printVarA() {
	//变量声明
	var i1 int
	var b1 bool

	//变量赋值
	i1 = 10
	b1 = false
	fmt.Println("int =", i1)
	fmt.Println("bool =", b1)

	//变量赋值+自动声明
	i2 := 20
	b2 := false
	fmt.Println("int =", i2)
	fmt.Println("bool =", b2)
	//声明初始化多个变量
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)
	/*
	int = 10
	bool = false
	int = 20
	bool = false
	1 2 3
	*/

}
func printVarB() {
	//简介的声明
	var (
		i1 int
		b1 bool
	)
	//赋值
	i1 = 10
	b1 = true
	fmt.Println("int =", i1)
	fmt.Println("bool =", b1)

	//多重赋值
	i2, i3 := 20, 30
	//多重赋值,废弃"_"
	_, i4 := 40, 41
	fmt.Println("int =", i1)
	fmt.Println("int =", i2)
	fmt.Println("int =", i3)
	fmt.Println("int =", i4)
	//赋值运算
	var i5 int
	var i5_32 int32
	i5 = 50
	fmt.Println("int i5=", i5)
	//i5_32 = i5 + i5  <--错误的赋值
	fmt.Println("int i5_32=", i5_32)
	i5_32 = i5_32 + 50
	fmt.Println("int i5_32=", i5_32)
	/*
		int = 10
		bool = true
		int = 10
		int = 20
		int = 30
		int = 41
		int i5= 50
		int i5_32= 0
		int i5_32= 50
	*/
}
func printVarC() {
	//枚举赋值1
	const (
		a = iota
		b = iota
	)

	fmt.Println("enum =", a)
	fmt.Println("enum =", b)
	//枚举赋值2
	const (
		a1 = 10
		b2 = iota
		c3
		d4
	)
	fmt.Println("enum =", a1)
	fmt.Println("enum =", b2)
	fmt.Println("enum =", c3)
	fmt.Println("enum =", d4)
	//各种常量声明
	const (
		a3        = 0
		b3 string = "0"
	)
	fmt.Println("a3 =", a3)
	fmt.Println("b3 =", b3)
	/*
		enum = 0
		enum = 1
		enum = 10
		enum = 1
		enum = 2
		enum = 3
		a3 = 0
		b3 = 0
	*/
}
func sampleVarD() {
	s := "Hello CongCong"
	fmt.Println(s)
	c := []byte(s)
	c[0] = 'C'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
	//	s3 := "Starting part"
	//	+"Ending part"

	s3 := "Starting part" +
		"Ending part"
	s4 := `Starting part
Ending part`
	fmt.Printf("%s\n", s3)
	fmt.Printf("%s\n", s4)

	/*
		Hello CongCong
		Cello CongCong
		Starting partEnding part
		Starting part
		Ending part
	*/
}
