package cache

import (
	"WB2/pattern/07_strategy/interfaces"
	"fmt"
)

type lfu struct {
}

func NewLfu() *lfu {
	return &lfu{}
}

func (l *lfu) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lfu 07_strategy")
}
