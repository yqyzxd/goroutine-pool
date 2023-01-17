/*
Package pool
@Time : 2023/1/17 15:49
@Author : wind
@Description : pool_test
*/
package pool

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	size := 10
	//1. create pool
	pool, _ := NewPool(size)
	fmt.Printf("num:%d\n", runtime.NumGoroutine())
	loopSize := 100
	var values []string
	//使用chan在goroutine之间共享数据
	ch := make(chan string)
	go func() {
		for s := range ch {
			values = append(values, s)
		}
	}()
	for i := 0; i < loopSize; i++ {
		//2. add to pool
		pool.Add(1)
		go work(pool, i, ch)
	}
	fmt.Println("Wait")
	pool.Wait()
	//务必在wait之后操作数据
	fmt.Printf("values length:%d\n", len(values))
	fmt.Println("Finish")
}
func work(pool *Pool, index int, ch chan string) {
	//3. release from pool
	defer pool.Done()
	time.Sleep(time.Second)
	ch <- strconv.Itoa(index)

}
