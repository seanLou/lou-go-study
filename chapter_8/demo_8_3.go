package main

import "fmt"

// 链表结构体
type Node struct {
	data float64
	su   *Node
}

// 二叉树
type Tree struct {
	// 左节点
	le   *Tree
	data float64
	// 右节点
	ri *Tree
}

type number struct {
	f float32
}

// 类型别名
type nr number

func main() {
	n1 := number{5.0}
	n2 := nr{5.0}

	// var i float32 = n2
	// 编译错误
	// Cannot use 'n2' (type nr) as type float32 in assignment

	// var i = float32(n2)
	// 编译错误
	// Cannot convert expression of type nr to type float32

	var n3 = number(n2)

	fmt.Println(n1, n2, n3)
}
