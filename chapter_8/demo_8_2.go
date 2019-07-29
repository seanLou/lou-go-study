package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 作为指针
	person := new(Person)
	person.firstName = "楼"
	person.lastName = "关杨"
	upPerson(person)
	fmt.Printf("我的名字叫 %s %s\n", person.firstName, person.lastName)

	// 字面量
	liLei := &Person{"li", "lei"}
	upPerson(liLei)
	fmt.Printf("我的名字叫 %s %s\n", liLei.firstName, liLei.lastName)

	wangYiyi := Person{"王", "一一"}
	upPerson(&wangYiyi)
	fmt.Printf("我的名字叫 %s %s\n", wangYiyi.firstName, wangYiyi.lastName)

	hanMeimei := Person{lastName: "meimei", firstName: "han"}
	upPerson(&hanMeimei)
	fmt.Printf("我的名字叫 %s %s\n", hanMeimei.firstName, hanMeimei.lastName)
}
