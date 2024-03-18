/*
 * @Author: zcw
 * @Date: 2024-03-04 14:53:15
 * @LastEditTime: 2024-03-15 14:38:29
 * @Description: In User Settings Edit
 * @FilePath: \studyDemo\main.go
 */
package main

import (
	"fmt"
	controldemo "studyDemo/public/controlDemo"
	stringdemo "studyDemo/public/stringDemo"
	timedemo "studyDemo/public/timeDemo"
)

func main() {

	stringdemo.StringTest()

	timedemo.TimeTest()

	controldemo.ControlTest()

	/*fmt.Println("go project")

	hello.PrintHello()
	hello.MyShaper()

	d := &hello.Dog{}
	d.SetName("小汪汪")

	c := &hello.Cat{}
	c.SetName("小喵喵")

	animal := []hello.Animal{}

	for _, v := range animal {
		v.Eat()
	}

	var animal1 hello.Animal
	animal1 = c

	if a, ok := animal1.(*hello.Cat); ok {
		fmt.Println("断言：", ok, a)
	}

	password := []byte("zhangsan")
	newPassword := md5.Sum(password)
	pStr := hex.EncodeToString(newPassword[:])
	fmt.Println(pStr)

	error1 := errors.New("这是一个错误")
	fmt.Println(error1)

	arr1 := pie.Of([]int{2, 3, 1, 5, 9, 4, -1})
	newArr1 := arr1.SortUsing(func(a, b int) bool {
		return a > b
	})
	fmt.Println(newArr1.Result)

	filtArr := arr1.Filter(func(i int) bool {
		return i > 4
	})
	fmt.Println(filtArr.Result)

	arrStr := pie.Of([]string{"a", "i", "f", "x", "b"})
	arrStr = arrStr.Insert(3, "p")
	fmt.Println(arrStr.Result)

	sortArr := pie.Sort([]int{2, 3, 1, 5, 9, 4, -1})
	fmt.Println(sortArr)

	// name := pie.Of([]string{"Bob", "Sally", "John", "Jane"})
	// name.Kyes()

	arr11 := []int{1, 0, 4, 9, 3, 8}
	slice1 := arr11[:]
	fmt.Println(slice1)

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// gresp, gerr := grequests.Get("https://jsonplaceholder.typicode.com/todos/1", nil)
	// if gerr != nil {
	// 	fmt.Println(gresp)
	// }
	// fmt.Println(gresp)

	rune1 := []rune("你好呀，嘿嘿嘿")
	for _, v := range rune1 {
		fmt.Println(string(v))
	}
	fmt.Println(string(rune1))

	r11 := pie.Of(rune1)
	index := r11.FindFirstUsing(func(value rune) bool {
		return value == rune1[1]
	})
	fmt.Println("查到了：", index)

	str11 := "你好呀，嘿嘿嘿"
	for _, v := range str11 {
		fmt.Println(string(v))
	}
	fmt.Println(str11, len(str11), utf8.RuneCountInString(str11))

	strIndex := strings.Index(string(rune1), "好")
	fmt.Println(strIndex)

	t := time.Now()

	tStr := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()) + " " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())

	fmt.Println("当前时间：", tStr)

	byte1 := new(bytes.Buffer)
	byte1.Write([]byte{'a', 'b', 'c'})
	fmt.Println(byte1.String())
	//bute1.buf = ['a', 'b', 'c']

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	copy(slTo, slFrom)
	fmt.Println(slTo, slFrom)

	str2 := "hello word"
	str3 := "你好，世界, 你好，世界"

	fmt.Println("测试测试：", strings.IndexRune(str3, '世'))
	str4 := strings.Replace(str3, "世界", "小伙纸", 2)
	fmt.Println("str4：", str4)

	sar2 := []byte(str2)
	sar3 := []rune(str3)

	fmt.Println(sar2)
	fmt.Println(sar3)
	for _, v := range sar3 {
		fmt.Println(string(v))
	}

	// ch1 := goTest()
	// ch2 := goTest()

	// a := <-ch1
	// b := <-ch2

	// goTestAll(a, b)

	// errorTest()

	// count := 0
	// for {
	// 	data, ok := <-ch1
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("收到数据了：", data)
	// 	count++
	// 	if count == 100 {
	// 		close(ch1)
	// 	}
	// }

	var buff bytes.Buffer
	buff.WriteString("hello，呀")
	fmt.Println("buffer: ", buff.String())

	*/
}

func goTest() chan int {

	ch := make(chan int)
	for i := 0; i < 100; i++ {
		go func(data int) {
			ch <- data
		}(i)
	}
	return ch
}

func goTestAll(a, b int) {
	fmt.Println("收到数据了：", a, b)
}

func errorTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误恢复中...", err)
		}
	}()

	ch3 := make(chan int, 1)
	close(ch3)
	ch3 <- 6
}
