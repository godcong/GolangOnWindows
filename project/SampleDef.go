/* Go保留字
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var
*/
/*
表格2.2列出了Go中所有的保留字。在下面的段落和章节中会介绍它们。其中有一些已经遇
到过了。
• var和const参阅”变量、类型和保留字”在3页；
• package和import已经有过短暂的接触，在”Hello World”部分。在第4章对其有详细
的描述。
其他都有对应的介绍和章节：
• func用于定义函数和方法；
• return用于从函数返回，func和return参阅第3章了解详细信息；
• go用于并行（第7章）；
• select用于选择不同类型的通讯，参阅第7章；
• interface参阅第6章；
• struct用于抽象数据类型，参阅第5章；
• type同样参阅第5章。
*/
package main

import (
	"fmt"
	"os"
)

type stringArrary struct {
	str  string
	size int
}

var i1 int

func sampleDefPlus() {
	if i1 >= 50 {
		fmt.Println("over", i1)
	} else {
		fmt.Println("under", i1)
	}
}

//func函数定义
func runDef() int {
	//var,const赋值

	const (
		i2 = iota
	)
	fmt.Println(i1, i2)

	i1 += 10
	//返回
	return i1
}

func runDefRet() {
	file, err := os.Open("maintest")
	if err != nil {
		println("it is nil")
		file, err = os.Create("maintest")
		s := file.Name()
		println(s)
	} else if err == nil {
		s := file.Name()
		b := make([]byte, 1024)
		file.Read(b)
		println("not nil")
		println(s)
		s = string(b)
		println(s)
	}
	i, _ := file.WriteString("我要写一些东西进去")
	file.Close()
	println(i)
	_, err = os.Open("maintest")
	if err != nil {
		println("it is nil!!!")
	} else if err == nil {
		println("not nil")
	}

	//下面的判断是成立的		
	if true && true {
		println("true")
	}
	if !false {
		println("true")
	}

}

func runDefGoto() (int, int) {
	//Go有goto语句——明智的使用它。用goto跳转到一定是当前函数内定义的标签。
	i := 0
Here:
	i++
	if i > 50 {
		return i, 10
	}
	println(i)
	goto Here
}

func runDefUserType() {
	var sa stringArrary
	sa.str = "hello"
	sa.size = len(sa.str)
	fmt.Println(sa.str)
	fmt.Println(sa.size)
}

func runDefSwitch(i int) {
	switch {
	case i > 0 || i == 0:
		println("bigger")
	case i < 0:
		println("smaller")
		fallthrough
	default:
		println("output default!!!")
	}

	switch i {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8: // 空的case 体
	case 9:
		println("当i == 0 时，f 不会被调用！")
	}
}
func runDefFor() {
	//Go的for循环有三种形式，只有其中的一种使用分号。
	//for init; condition; post { }  和C 的for 一样
	//for condition { }  和while 一样
	//for { }  和C 的for(;;) 一样（死循环）
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i //  sum = sum + i 的简化写法
	} //  i 实例在循环结束会消失
}

func runDefRange() {
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		// 对k 和v 做想做的事情
		println(k)
		println(v)
	}
}
