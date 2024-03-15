/*
 * @Author: zcw
 * @Date: 2024-03-04 15:12:08
 * @LastEditTime: 2024-03-04 15:52:19
 * @Description: In User Settings Edit
 * @FilePath: \studyDemo\hello\hello.go
 */
package hello

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) resetPersonName(name string) {
	p.name = name
}

func PrintHello() {
	fmt.Println("Hello模块 => say hello")

	p1 := &person{name: "张三", age: 18}
	fmt.Println(*p1)

	p11 := p1
	p1.resetPersonName("李四")
	fmt.Println(p1, p11)

	p2 := person{"王五", 21}
	p3 := p2
	p2.resetPersonName("王小五")
	fmt.Println(p2, p3)
}
