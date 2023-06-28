package cache

import (
	"WB2/pattern/07_strategy/interfaces"
	"fmt"
)

type lru struct {
}

func NewLru() *lru {
	return &lru{}
}

func (l *lru) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lru 07_strategy")
}
