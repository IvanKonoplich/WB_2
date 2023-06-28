package cache

import (
	"WB2/patterns/strategy/interfaces"
	"fmt"
)

type fifo struct {
}

func NewFifo() *fifo {
	return &fifo{}
}

func (l *fifo) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by fifo strategy")
}
