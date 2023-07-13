package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Item struct {
	Key   Key
	Value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
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
	if element, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(element)
		return element.Value.(*Item).Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.capacity = 0
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
