package main

import "WB2/pattern/strategy/cache"

func main() {
	lfu := cache.NewLfu()
	cache1 := cache.InitCache(lfu)
	cache1.Add("a", "dev01")
	cache1.Add("b", "dev02")
	cache1.Add("c", "dev03")
	lru := cache.NewLru()
	cache1.SetEvictionAlgo(lru)
	cache1.Add("d", "dev04")
	fifo := cache.NewFifo()
	cache1.SetEvictionAlgo(fifo)
	cache1.Add("e", "dev05")
}
