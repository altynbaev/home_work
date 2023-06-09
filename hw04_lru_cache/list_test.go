package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push one element at front", func(t *testing.T) {
		l := NewList()

		front := l.PushFront(10)

		require.Equal(t, 1, l.Len())
		require.Equal(t, front, l.Front())
		require.Equal(t, front, l.Back())
	})

	t.Run("push one element at back", func(t *testing.T) {
		l := NewList()

		back := l.PushBack(10)

		require.Equal(t, 1, l.Len())
		require.Equal(t, back, l.Front())
		require.Equal(t, back, l.Back())
	})

	t.Run("push and remove one element", func(t *testing.T) {
		l := NewList()

		front := l.PushFront(10)
		l.Remove(front)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push one element and remove another element", func(t *testing.T) {
		l := NewList()

		front := l.PushFront(10)
		l.Remove(&ListItem{})

		require.Equal(t, 1, l.Len()) // элемент не удаляется
		require.Equal(t, front, l.Front())
		require.Equal(t, front, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
