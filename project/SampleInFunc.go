/*
Table 2.3.Go中的预定义函数
close new panic complex
closed make recover real
len append print imag
cap copy println
*/
/*
close和closed 用于channel通讯和关闭channel，参阅第7章了解更多。
len和cap 可用于不同的类型，len用于返回字符串、slice和数组的长度。参阅”array、slices
和map”小节了解更多关于slice、数组和函数cap的详细信息。
new 用于各种类型的内存分配。参阅”用new分配内存”在53页。
make 用于内建类型（map、slice和channel）的内存分配。参阅”用make分配内存”在53
页。
copy 用于复制slice。append用于追加slice。参阅本章的”slice”。
panic和recover 用于异常处理机制。参阅”恐慌（Panic）和恢复（Recover）”在31页了
解更多信息。
print和println 是底层打印函数，可以在不引入fmt 包的情况下使用。它们主要用于调
试。
complex、real和imag 全部用于处理复数。有了之前给的简单的例子，不用再进一步讨论
复数了。
*/
package main

//import "fmt"

func sampleFuncArrary() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	println("The first element is %d\n", arr[0])
	/*
		可以像这样声明一个数组：var a [3]int，如果不使用零来初始化它，则用复合声
		明：a := [3]int{1, 2, 3}也可以简写为a := [...]int{1, 2, 3}，Go会自动统计元素的
		个数。注意，所有项目必须都指定。因此，如果你使用多维数组，有一些内容你必须录入： 复 合 声 明 允 许 你
		直 接 将 值 赋 值
		给array 、slice 或
		者map
	*/
	a1 := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	//类似于：
	a2 := [2][2]int{[...]int{1, 2}, [...]int{3, 4}}
	/*
		当声明一个array时，你必须在方括号内输入些内容，数字或者三个点(...)。从release.2010-10-27[7]
		这个语法使用更加简单了。来自于发布记录：	
		array、slice和map的复合声明变得更加简单。使用复合声明的array、slice
		和map，元素复合声明的类型与外部一直，则可以省略。
		这表示上面的例子可以修改为：

	*/
	a3 := [][]int{{1, 2}, {3, 4}}
	for i, k := range a1[1] {
		println("a1-i", i)
		println("a1-k", k)
		/*
		a1-i 0
		a1-k 3
		a1-i 1
		a1-k 4
		*/
	}
	for i, _ := range a2 {
		println("a2", i)

	}
	for i, _ := range a3 {
		println("a3", i)

	}

}

func sampleFuncSlice() {
	/*
		slice与array接近，但是在新的元素加入的时候可以增加长度。slice总是指向底层的一
		个array。slice是一个指向array的指针，这是其与array不同的地方；slice是引用类型，
		这意味着当赋值某个slice到另外一个变量，两个引用会指向同一个array。例如，如果一个
		函数需要一个slice参数，在其内对slice元素的修改也会体现在函数调用者中，这和传递底
		层的array指针类似。通过：
	*/
	s := make([]int, 10)
	_ = s
	/*
		创建了一个保存有10个元素的slice。需要注意的是底层的array并无不同。slice总是与一个
		固定长度的array成对出现。其影响slice的容量和长度。图2.1描述了下面的Go代码。首先
		创建了m个元素长度的array，元素类型int：var array[m]int
		然后对这个array创建slice：slice := array[0:n]
		然后现在有：
		len(slice)== n == cap(slice)== n
		len(array)== cap(array)== m
	*/
	/*
		给定一个array或者其他slice，一个新slice通过a[I:J]的方式创建。这会创建一个新
		的slice，指向a，从序号I开始，结束在序号J之前。长度为J - I。
		array[n:m] 从array 创建了一个slice，具有元素n to m-1
	*/
	a := [...]int{1, 2, 3, 4, 5} //.0 定义一个5个元素的array，序号从0到4；
	s1 := a[2:4]                 //.1 从序号2至3创建slice，它包含元素3, 4；
	s2 := a[1:5]                 //.2 从序号1至4创建，它包含元素2, 3, 4, 5；
	s3 := a[:]                   //.3 用array中的所有元素创建slice，这是a[0:len(a)]的简化写法；
	s4 := a[:4]                  //.4 从序号0至3创建，这是a[0:4]的简化写法，得到1, 2, 3, 4；
	s5 := s2[:]                  //.5 从slices2创建slice，注意s5仍然指向arraya。
	_ = a
	_, _, _, _, _ = s1, s2, s3, s4, s5
	/*
		在2.6列出的代码中，我们在第八行尝试做一些错误的事情，让一些东西超出范围（底
		层array的最大长度），然后得到了一个运行时错误。
		Listing 2.6. array和slice
		1 package main
		3 func main() {
		4 var array [100]int // Create array, index from 0 to 99
		5 slice := array[0:99] // Create slice, index from 0 to 98
		7 slice[98] = 'a' // OK
		8 slice[99] = 'a' // Error: "throw: index out of range"
		9 }
	*/
}
func sampleFuncSliceAppend() {
	/*
			如果你想要扩展slice，有一堆内建函数让你的日子更加好过一些：append和copy。来自
		于[5]：
		函数append向slices追加零值或其他x值，并且返回追加后的新的、与s有相
		同类型的slice。如果s没有足够的容量存储追加的值，append分配一个足够大
		的、新的slice来存放原有slice的元素和追加的值。因此，返回的slice可能指
		向不同的底层array。
	*/
	s0 := []int{0, 0}
	s1 := append(s0, 2)       //.0 追加一个元素，s1 == []int{0, 0, 2}；
	s2 := append(s1, 3, 5, 7) //.1 追加多个元素，s2 == []int{0, 0, 2, 3, 5, 7}；
	s3 := append(s2, s0...)   //.2 追加一个slice，s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}。注意这三个点！
	_, _, _, _ = s0, s1, s2, s3
	/*
		还有
		函数copy从源slicesrc复制元素到目标dst，并且返回复制的元素的个数。源
		和目标可能重叠。复制的数量是len(src)和len(dst)中的最小值。
	*/
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	n1 := copy(s, a[0:]) //n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	n2 := copy(s, s[2:]) //n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	_, _ = n1,n2
}
func sampleFuncMap() {
	/*
	许多语言都内建了类似的类型，例如Perl有哈希，Python有字典，而C++同样也有map
	（在lib中）。在Go中有map类型。map可以认为是一个用字符串做索引的数组（在其最简单
	的形式下）。下面定义了map类型，用于将string（月的缩写）转换为int——那个月的天
	数。一般定义map的方法是：map[<from type>]<to type>
	*/
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //逗号是必须的
	}
	for s, i := range monthdays {
		println(s, i)
	}
	/*
	留意，当只需要声明一个map的时候，使用make的形式：monthdays := make(map[string]
	int)
	当在map中索引（搜索）时，使用方括号，例如打印出12月的天数：fmt.Printf("%d\
	n", monthdays["Dec"])
	当对array、slice、string或者map循环遍历的时候，range会帮助你，每次调用，它都会
	返回一个键和对应的值。
	year := 0
	for _, days := range monthdays {  键没有使用，因此用_, days
	year += days
	}
	fmt.Printf("Numbers of days in a year: %d\n", year)
	向map增加元素，可以这样做：
	monthdays["Undecim"] = 30  添加一个月
	monthdays["Feb"] = 29  闰年时重写这个元素
	检查元素是否存在，可以使用下面的方式[26]：
	var value int
	var present bool
	value, present = monthdays["Jan"]  如果存在，present 则有值true
	 或者更接近Go 的方式
	v, ok := monthdays["Jan"]  “逗号ok”形式
	也可以从map中移除元素：
	monthdays["Mar"] = 0, false  删除"Mar" 吧，总是下雨
	看起来有点像把“逗号ok”形式反过来。
	*/
}
