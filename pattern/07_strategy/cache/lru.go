package cache

import (
	"WB2/pattern/strategy/interfaces"
	"fmt"
)

type lru struct {
}

func NewLru() *lru {
	return &lru{}
}

func (l *lru) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lru strategy")
}
