package main

import "fmt"

type ElementExistsError[T comparable] struct {
	value T
}

func (e *ElementExistsError[T]) Error() string {
	return fmt.Sprintf("Element %v already exists", e.value)
}

type ListIsEmptyError struct{}

func (e *ListIsEmptyError) Error() string {
	return "List is empty"
}

type Node[T comparable] struct {
	next  *Node[T]
	value T
}

type SinglyLinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

func (l *SinglyLinkedList[T]) AddFront(value T) *Node[T] {
	node := &Node[T]{nil, value}
	if l.length == 0 {
		l.head = node
		l.length++
	} else {
		node.next = l.head
		l.head = node
		l.length++
	}
	return node
}

func (l *SinglyLinkedList[T]) AddBack(value T) *Node[T] {
	node := &Node[T]{nil, value}
	if l.length == 0 {
		l.head = node
		l.length++
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
		l.length++
	}
	return node
}

func (l *SinglyLinkedList[T]) RemoveFront() *Node[T] {
	if l.length == 0 {
		panic(&ListIsEmptyError{})
	} else {
		node := l.head
		l.head = node.next
		l.length--
		return node
	}
}

func (l *SinglyLinkedList[T]) RemoveBack() *Node[T] {
	if l.length == 0 {
		panic(&ListIsEmptyError{})
	} else {
		current := l.head
		for current.next.next != nil {
			current = current.next
		}
		node := current.next
		current.next = nil
		l.length--
		return node
	}
}

func (l *SinglyLinkedList[T]) Contains(value T) bool {
	current := l.head
	for current != nil {
		if current.value == value {
			return true
		}
	}
	return false
}

func (l *SinglyLinkedList[T]) IndexOf(value T) int {
	current := l.head
	for i := 0; current != nil; i++ {
		if current.value == value {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *SinglyLinkedList[T]) RemoveAt(index int) *Node[T] {
	if l.length == 0 {
		panic(&ListIsEmptyError{})
	} else {
		if index == 0 {
			node := l.head
			l.head = node.next
			l.length--
			return node
		} else {
			current := l.head
			for i := 0; i < index-1; i++ {
				current = current.next
			}
			node := current.next
			current.next = node.next
			l.length--
			return node
		}
	}
}

func (l *SinglyLinkedList[T]) Remove(value T) *Node[T] {
	if l.length == 0 {
		panic(&ListIsEmptyError{})
	} else {
		current := l.head
		for current.next != nil {
			if current.value == value {
				node := current.next
				current.next = current.next.next
				return node
			}
		}
		l.length--
	}
	return nil
}

func (l *SinglyLinkedList[T]) Display() {
	if l.length == 0 {
		fmt.Println(&ListIsEmptyError{})
	} else {
		current := l.head
		for current != nil {
			fmt.Println(current.value)
			current = current.next
		}
	}
}

func main() {
	list := SinglyLinkedList[int]{}
	list.AddFront(1)
	list.AddFront(2)
	list.AddFront(3)
	list.AddBack(4)
	list.AddBack(5)
	list.AddBack(6)
	list.Display()
	list.RemoveFront()
	defer list.Display()
}
