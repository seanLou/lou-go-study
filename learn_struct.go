package main

import (
	"fmt"
	"os"
)

const (
	ONE   = iota //0
	TWO          //1
	THREE        //2
	FOUR         //3
	FIVE         //4
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func init() {
	fmt.Println("init 函数在包初始化时自动执行，优先级比main函数高，一个源文件有且只有一个init函数。init函数不能被手动调用")
}

func main() {
	const PI = 3.1415926
	const PI2 float32 = 3.1415926

	const SUCCESS = true

	const HELLO = "hello"
	const WORLD string = " world"

	fmt.Println("PI 的值是:", PI)
	fmt.Println("PI2 的值是:", PI2)
	fmt.Println("SUCCESS 的值是:", SUCCESS)
	fmt.Println(HELLO, WORLD)

	a := 27 // 等同于 var a = 27
	c := 100

	c = a
	fmt.Println("把 a 赋值给 c, c = ", c)
	c += a
	fmt.Println(" c = c + a, c = ", c)

	c -= a
	fmt.Println("c = c - a, c = ", c)

	c *= a
	fmt.Println("c = c * a, c = ", c)

	c /= a
	fmt.Println("c = c / a, c = ", c)

	fmt.Println("打印枚举的值：", ONE, TWO, THREE, FOUR, FIVE)

	fmt.Println("以下是 变量定义 的学习》》》》》")
	var count int = 1
	var flag bool = false
	var str string = "hi"
	var d, f int8
	d = 9
	f = 12
	fmt.Println(count, flag, str, d, f)

	var j, k, l int
	j = 1
	k = 2
	l = 3
	fmt.Println(j, k, l)
	var x, y, z int = 788, 298, 2334
	fmt.Println(x, y, z)

	var a1, b1, c1 = "这是a1", 10, true

	fmt.Println("a1:", a1, ",b1:", b1, ",c1:", c1)

	fmt.Println(HOME, USER, GOROOT) // 获取系统相关常量

	fmt.Println("以下是 变量交换 的学习》》》》》")
	a = 1
	b := 2
	fmt.Println("a:", a, ", b:", b)
	a, b = b, a // 交换 a, b的值
	fmt.Println("交换后 a:", a, ", b:", b)

	fmt.Println("以下是 匿名函数 的学习》》》》》")

	className, _, teacher := getClass() // 获取班级名，老师名，人数不使用, 用匿名函数替代 _
	fmt.Println(className, "班主任是", teacher)

	fmt.Println("以下是 字符串 的学习》》》》》")

	// str := "你好" // 会产生编译错误
	str = "你好"
	fmt.Println(str)
	t := str
	str += ", Go 语言" // 字符串可以拼接， 但原字符无法修改
	fmt.Println(str)
	fmt.Println(t) // 原字符串

	fmt.Println("以下是 非解释字符串 的学习》》》》》")

	str = `小楼昨夜又东风 \n
下面是啥忘记了
	换行` // 反单引号可以换行, 但是引号内的所有字符都会被输出

	fmt.Println(str)

	str = "你好啊\nGo\t语言" // 双引号不能换行，会解析转义字符

	fmt.Println(str)

	fmt.Println("以下是 字符串操作 的学习》》》》》")

	str1 := "abcd 你"

	fmt.Println("字符串 str1 长度：", len(str1))
	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
	}

	str2 := str1[len(str1)-1]
	fmt.Println("最后一个字符：", str2)

	// len(str1) = 8, 截取字符串时 不包含最后一个字节, s[8] 会发生下标越界
	fmt.Println("打印所有字符》", str1[0:8])

	fmt.Println("以下是 字符串比较 的学习》》》》》")

	str = "你"
	str1 = "好"

	fmt.Println("str:", str, ", str1:", str1)
	// 字符串中比较大小是根据第一个字节来判断的
	if str < str1 {
		fmt.Println(str[0], str1[0])
		fmt.Println(str[1], str1[1])
		fmt.Println(str[2], str1[2])
	}

	fmt.Println("以下是 字符串遍历 的学习》》》》》")

	str = "我是中国人" // 汉字占3个字节

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n输出单字节值")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d = %v\n", i, str[i])
	}

	for _, v := range str {
		fmt.Printf("%c", v)
	}

	fmt.Println("\n以下是 字符串修改 的学习》》》》》")

	str = "hello 世界!"
	bytee := []byte(str) // 转换为 []byte, 自动复制数据
	bytee[5] = ','
	fmt.Printf("原字符串：%s\n", str)
	fmt.Printf("修改后的字符串：%s\n", bytee)

	r := []rune(str) // 转换为 []rune, 自动复制数据
	r[6] = '杭'
	r[7] = '州'

	fmt.Printf("修改后的字符串：%s\n", string(r)) // 转换为 字符串, 自动复制数据

	var arrays = [7]int{0, 1, 2, 3, 4, 5, 6}
	slice1 := arrays[0:5]
	slice2 := arrays[2:]
	slice3 := arrays[:5]
	fmt.Printf("数组arrays 下标 0 长度 4 切片 %v, slice1 len: %d \n", slice1, len(slice1))
	fmt.Printf("数组arrays 下标 2 开始 切片 %v, slice2 len: %d \n", slice2, len(slice2))
	fmt.Printf("数组arrays 下标 0 开始 长度 5 切片： %v, slice3 len: %d \n", slice3, len(slice3))

	slice1[2] = 22
	fmt.Printf("修改 slice1-下标2 元素后, slice1: %v\n", slice1)
	slice3[4] = 16
	fmt.Printf("修改 slice3-下标4元素后, slice2: %v\n", slice2)
	slice2[4] = -5
	fmt.Printf("修改 slice2-下标4元素后, slice3: %v\n", slice3)
	fmt.Printf("arrays: %v\n", arrays)

	makeSlice1 := make([]int, 10)
	makeSlice2 := make([]int, 10, 15)
	fmt.Println("makeSlice1 len:", len(makeSlice1), ", cap:", cap(makeSlice1))
	fmt.Println("makeSlice2 len:", len(makeSlice2), ", cap:", cap(makeSlice2))

	appendSlice := []int{10, 20, 30, 40}
	// 追加一个元素 底层数组容量不足，容量小于 1000 时, 容量成倍增加; 大于 1000 时， 1.25倍
	appendSlice2 := append(appendSlice, 50)
	fmt.Printf("appendSlice len:%d, cap:%d, appendSlice2 len:%d, cap:%d", len(appendSlice), cap(appendSlice),
		len(appendSlice2), cap(appendSlice2))

}

func getClass() (className string, studentNum int, teacher string) {
	return "103班", 50, "楼老师"
}
