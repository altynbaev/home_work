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

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if element, ok := cache.items[key]; ok {
		cache.items[key].Value = value
		cache.queue.MoveToFront(element)
		return true
	}

	// Добавление элемента в кэш
	listItem := cache.queue.PushFront(value)
	cache.items[key] = listItem

	if cache.queue.Len() > cache.capacity {
		// Удаляем лишний элемент
		lastElement := cache.queue.Back()
		cache.queue.Remove(lastElement)
		delete(cache.items, key)
	}

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// Если элемент присутствует в кэше, перемещаем на первое место
	if element, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(element)
		return element.Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
