package main

import "fmt"

// DoublyLinkedList의 리소스 선언
type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
}

// DoublyNode의 리소스 선언
type DoublyNode struct {
	// 노드가 담고있을 값
	DataField int
	// 다음 노드 주소
	LinkField *DoublyNode
	// 이전 노드 주소
	Prev *DoublyNode
}

func main() {
	list := DoublyLinkedList{}
	//
	list.PushFront(1)
	list.FindAll()
	list.PushFront(2)
	list.FindAll()
	list.PushRear(3)
	list.FindAll()
	list.PushRear(4)
	list.FindAll()
	list.PopFront()
	list.FindAll()
	list.PopFront()
	list.FindAll()
	list.PopRear()
	list.FindAll()
	list.PopRear()
	list.FindAll()
}

// 1-1. Push_Front
func (l *DoublyLinkedList) PushFront(data int) {
	if l.Head == nil {
		newNode := &DoublyNode{
			DataField: data,
			LinkField: nil,
			Prev:      nil,
		}
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode := &DoublyNode{
			DataField: data,
			LinkField: l.Head,
			Prev:      nil,
		}
		l.Head.Prev = newNode
		l.Head = newNode
	}
}

// 1-2. Push_Rear
func (l *DoublyLinkedList) PushRear(data int) {
	if l.Tail == nil {
		newNode := &DoublyNode{
			DataField: data,
			LinkField: nil,
			Prev:      nil,
		}
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode := &DoublyNode{
			DataField: data,
			LinkField: nil,
			Prev:      l.Tail,
		}
		l.Tail.LinkField = newNode
		l.Tail = newNode
	}
}

// 1-3. Pop_Front
func (l *DoublyLinkedList) PopFront() {
	// Head 비어있을때
	if l.Head == nil {
		fmt.Println("err")
		return

		// 리스트에 값이 한 개만 있을 때
	} else if l.Tail == l.Head {
		l.Head = nil
		l.Tail = nil
		return

		// Head값을 Head값의 다음 노드
	} else {
		l.Head = l.Head.LinkField
		l.Head.Prev = nil
	}
}

// 1-4. Pop_Rear
func (l *DoublyLinkedList) PopRear() {
	// Tail이 비어있을때
	if l.Tail == nil {
		fmt.Println("err")
		return
	}
	// 리스트에 값이 한 개만 있을 때
	if l.Tail == l.Head {
		l.Head = nil
		l.Tail = nil
		return
	} else {
		l.Tail = l.Tail.Prev
		l.Tail.LinkField = nil
	}
}

// 2-1. Head 조회
func (l *DoublyLinkedList) FindHead() {
	current := l.Head
	if current != nil {
		fmt.Printf("Head : %d\n", current.DataField)
	}
}

// 2-2. Tail 조회
func (l *DoublyLinkedList) FindTail() {
	current := l.Tail
	if current != nil {
		fmt.Printf("Tail : %d\n", current.DataField)
	}
}

// 2-3. 모든 값 조회
func (l *DoublyLinkedList) FindAll() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ", current.DataField)
		current = current.LinkField
	}
	fmt.Println()
}

// 2-4. 크기 조회
func (l *DoublyLinkedList) CountNodes() {
	current := l.Head
	var size int
	for current != nil {
		current = current.LinkField
		size++
	}
	fmt.Printf("길이 : %d\n", size)
}
