package hello

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func MyShaper() {

	sq1 := &Square{side: 5.0}
	area := sq1.Area()
	fmt.Println("面积：", area)
}

type Animal interface {
	Eat()
}

type Dog struct {
	name string
}

func (d *Dog) Eat() {
	fmt.Println("狗吃肉")
}

func (d *Dog) SetName(name string) {
	d.name = name
}

type Cat struct {
	name string
}

func (c *Cat) Eat() {
	fmt.Println("猫吃鱼")
}

func (c *Cat) SetName(name string) {
	c.name = name
}
