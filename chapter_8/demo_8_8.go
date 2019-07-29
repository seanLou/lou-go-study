package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	// tags
	hasStock bool   `是否有存货`
	name     string `商品名称`
	price    int    `商品价格`
}

func main() {
	tagType := TagType{true, "IPhone X", 8888}
	fieldSize := reflect.TypeOf(tagType).NumField()
	for i := 0; i < fieldSize; i++ {
		refTag(tagType, i)
	}
}

/**
在一个变量上调用 reflect.TypeOf() 可以获取变量的正确类型
如果变量是一个结构体类型，就可以通过Field来索引结构体的字段，然后就可以使用Tag属性了
*/
func refTag(tagType TagType, index int) {
	structField := reflect.TypeOf(tagType).Field(index)
	fmt.Printf("第 %d 个字段, 字段名 %s, Tag %v\n", index, structField.Name, structField.Tag)
}
