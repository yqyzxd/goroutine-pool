/*
Package goroutine_pool
@Time : 2023/1/17 15:39
@Author : wind
@Description : pool
*/
package pool

import (
	"errors"
	"sync"
)

type Pool struct {
	ch chan int
	wg *sync.WaitGroup
}

func NewPool(size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size must great than zero")
	}
	return &Pool{
		ch: make(chan int, size),
		wg: &sync.WaitGroup{},
	}, nil
}
func (p *Pool) Add(size int) {
	negative := false
	absSize := size
	if size < 0 {
		negative = true
		absSize = -size
	}
	for i := 0; i < absSize; i++ {
		if !negative {
			p.ch <- 1
		} else {
			<-p.ch
		}
	}
	p.wg.Add(size)
}

func (p *Pool) Done() {
	<-p.ch
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.ch)
}
