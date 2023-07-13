package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem // указатель на следующий элемент по списку
	Prev  *ListItem // указатель на предыдущий элемент по списку
}

type LinkedList struct {
	len   int
	first *ListItem // указатель на первый элемент списка
	last  *ListItem // указатель на последний элемент списка
}

func (list *LinkedList) Len() int { // получить длину списка
	return list.len
}

func (list *LinkedList) Front() *ListItem { // получить первый элемент списка
	return list.first
}

func (list *LinkedList) Back() *ListItem { // получить последний элемент списка
	return list.last
}

func (list *LinkedList) PushFront(v interface{}) *ListItem { // добавить значение в начало
	item := ListItem{ // item.Prev и item.Next будут nil по умолчанию
		Value: v,
	}

	if list.len != 0 {
		item.Next = list.first
		list.first.Prev = &item // обновляем указатель первого элемента на новый элемент
		list.first = &item      // обновляем указатель списка на первый элемент списка
	} else {
		list.first = &item
		list.last = &item
	}

	list.len++
	return &item
}

func (list *LinkedList) PushBack(v interface{}) *ListItem { // добавить значение в конец
	item := ListItem{ // item.Prev и item.Next будут nil по умолчанию
		Value: v,
	}

	if list.len != 0 {
		item.Prev = list.last
		list.last.Next = &item // обновляем указатель последнего элемента на новый элемент
		list.last = &item      // обновляем указатель списка на последний элемент списка
	} else {
		list.first = &item
		list.last = &item
	}

	list.len++
	return &item
}

func (list *LinkedList) Remove(i *ListItem) { // удалить элемент
	prev := i.Prev // указатель элемента на предыдущий элемент
	next := i.Next // указатель элемента на следующий элемент

	if prev != nil {
		prev.Next = next // обновляем указатель следующего элемента на предыдущий
	} else {
		list.first = next
		list.first.Prev = nil
	}
	if next != nil {
		next.Prev = prev // обновляем указатель предыдущего элемента на следующий
	} else {
		list.last = prev
		list.last.Next = nil
	}

	list.len--
}

func (list *LinkedList) MoveToFront(i *ListItem) { // переместить элемент в начало
	list.Remove(i)
	list.PushFront(i.Value)
}

func NewList() *LinkedList {
	return new(LinkedList)
}
