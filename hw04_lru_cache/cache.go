package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Item struct {
	Key   Key
	Value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	item := &Item{
		Key:   key,
		Value: value,
	}

	if element, ok := cache.items[key]; ok {
		cache.items[key].Value.(*Item).Value = value
		cache.queue.MoveToFront(element)
		return true
	}

	cache.items[key] = cache.queue.PushFront(item)
	if cache.queue.Len() > cache.capacity {
		lastElement := cache.queue.Back()
		cache.queue.Remove(lastElement)
		delete(cache.items, item.Key)
	}

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if element, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(element)
		return element.Value.(*Item).Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = map[Key]*ListItem{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
