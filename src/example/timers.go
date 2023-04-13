package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

func periodicTask() {
	t := time.NewTicker(5 * time.Second) // 创建一个每5秒触发一次的定时器
	for {
		select {
		case <-t.C:
			// 执行周期性任务的代码
			fmt.Println("starting...")
		}
	}
}

func timeoutTask() {
	done := make(chan bool, 1)
	timer := time.NewTimer(1 * time.Second) // 创建一个1秒的定时器
	go func() {
		// 执行需要控制超时时间的任务
		// 如果任务执行完毕，则向done通道发送true
		done <- true
	}()

	select {
	case <-done:
		// 任务完成，无需超时处理
	case <-timer.C:
		// 超时处理
	}
}

func delayedTask() {
	timer := time.NewTimer(1 * time.Minute) // 创建一个1分钟的定时器
	<-timer.C
	// 延迟执行的任务
}

func scheduleTask() {
	t := time.Now()
	next := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1) // 获取明天的时间
	timer := time.NewTimer(next.Sub(t))                                                        // 计算距离明天凌晨的时间差

	for {
		select {
		case <-timer.C:
			// 执行定时任务 do something

			// 更新timer，计算下一次任务执行的时间
			next = next.AddDate(0, 0, 1)
			timer.Reset(next.Sub(time.Now()))
		}
	}
}
