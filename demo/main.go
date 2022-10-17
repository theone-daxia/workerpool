package main

import (
	"fmt"
	"github.com/theone-daxia/workerpool"
	"time"
)

func main()  {
	p := workerpool.New(
		5,
		workerpool.WithPreAllocWorkers(false),
		workerpool.WithBlock(false))

	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 1)
		})
		if err != nil {
			fmt.Printf("task[%d]: error:%s\n", i, err)
		}
	}
	p.Free()
}