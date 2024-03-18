package timedemo

import (
	"fmt"
	"time"
)

func TimeTest() {

	time1 := time.Now()
	fmt.Println("当前时间：", time1)

	timeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time1.Year(), time1.Month(), time1.Day(), time1.Hour(), time1.Minute(), time1.Second())
	fmt.Println("当前时间：", timeStr)

	timeFstr := time1.Format("2006-01-02 15:04:05")
	fmt.Println("Format时间：", timeFstr)

	timeObj, _ := time.Parse("2006-01-02 15:04:05", timeFstr)
	fmt.Println("时间转换：", timeObj)

	time12 := timeObj.Add(12 * time.Hour)
	fmt.Println("12小时后：", time12.Format("2006-01-02 15:04:05"))

	// afCh := time.After(3 * time.Second)
	// if _, ok := <-afCh; ok {
	// 	fmt.Println("延迟3秒")
	// }

	// time.Sleep(3 * time.Second)
	// fmt.Println("sleep完成")

	// tickCh := time.Tick(1 * time.Second)
	// for {
	// 	if val, ok := <-tickCh; ok {
	// 		fmt.Println("定时器：", val)
	// 	}
	// }

	// tick := time.NewTicker(1 * time.Second)
	// go func() {
	// fmt.Println("执行定时器关闭")
	// time.Sleep(10 * time.Second)
	// tick.Stop()
	// }()

	// count := 0
	// for {
	// 	select {
	// 	case val := <-tick.C:
	// 		fmt.Println("定时器：", val)
	// 		count++
	// 	default:
	// 		// fmt.Println("select default")
	// 	}
	// 	if count == 10 {
	// 		tick.Stop()
	// 		break
	// 	}
	// }

	// tCh := make(chan int, 1)
	// go func(ch chan int) {
	// 	ch <- 888
	// 	// close(ch)
	// }(tCh)

	// go func(t *time.Ticker, ch chan int) {
	// 	count := 0
	// 	for {
	// 		count++

	// 		select {
	// 		case val1 := <-t.C:
	// 			fmt.Println("定时器：", val1)
	// 		case val, ok := <-ch:
	// 			if ok {
	// 				close(ch)
	// 				fmt.Println("另一个channel：", val, ok)
	// 			}

	// 		}
	// 		if count == 10 {
	// 			t.Stop()
	// 		}
	// 	}
	// }(tick, tCh)

	// time.Sleep(20 * time.Second)
}
