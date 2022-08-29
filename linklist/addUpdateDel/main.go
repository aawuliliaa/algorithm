package main

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}
type LinkList struct {
	dummyHead *Node
	size      int
}

func NewLinkList() *LinkList {
	return &LinkList{
		dummyHead: &Node{},
		size:      0,
	}
}

func (l *LinkList) verifyIndex(index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of range")
	}
	return nil
}
func (l *LinkList) add(index int, e interface{}) {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	pre.next = &Node{e: e, next: pre.next}
	l.size++
}
func (l *LinkList) AddFirst(e interface{}) {
	l.add(0, e)
}
func (l *LinkList) AddLast(e interface{}) {
	l.add(l.size, e)
}
func (l *LinkList) get(index int) interface{} {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}
func (l *LinkList) GetFirst() interface{} {
	return l.get(0)
}
func (l *LinkList) GetLast() interface{} {
	return l.get(l.size)
}
func (l *LinkList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for {
		if cur == nil {
			break
		}
		if e == cur.e {
			return true
		}
		cur = cur.next
	}
	return false

}
func (l *LinkList) ToString() string {
	cur := l.dummyHead.next
	str := ""
	for cur != nil {
		str += fmt.Sprintf("%d-->", cur.e.(int))
		cur = cur.next
	}
	return str
}
func (l *LinkList) update(index int, e interface{}) {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}
func (l *LinkList) UpdateFirst(e interface{}) {
	l.update(0, e)
}
func (l *LinkList) UpdateLast(e interface{}) {
	l.update(l.size, e)
}
func (l *LinkList) remove(index int) {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	delNode := pre.next
	pre.next = delNode.next
	delNode.next = nil
	l.size--
}
func main()  {
	l :=NewLinkList()
	for i:=0;i<3;i++{
		l.AddLast(i)
		fmt.Println(l.ToString())
	}
	l.update(1,5)
	fmt.Println(l.ToString())
	l.remove(1)
	fmt.Println(l.ToString())

}
