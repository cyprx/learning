package main

import "container/list"

type Pair struct {
	first  int
	second *list.Element
}

type LRUCache struct {
	c int
	m map[int]Pair
	l *list.List
}

func NewLRUCache(c int) *LRUCache {
	return &LRUCache{
		c: c,
		m: make(map[int]Pair),
		l: list.New(),
	}
}

// Put will add/update value to cache in O(1)
func (c *LRUCache) Put(int key, int val) {
	if p, ok := m[key]; ok {
		c.l.Remove(p.second)
	}
	if c.c == c.l.Len() {
		c.l.Remove(c.l.Front())
	}
	c.l.PushBack(key)
	u := Pair{first: val, second: c.l.Back()}
	c.m[key] = u
}

func (c *LRUCache) Get(int key) int {
	if p, ok := m[key]; !ok {
		return -1
	} else {
		c.l.Remove(p.second)
		c.l.PushBack(key)
		u := Pair{first: p.first, second: c.l.Back()}
		c.m[key] = u
		return p.first
	}
}
