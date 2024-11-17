package list

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int       // data in a node
	Next *ListNode // address of next node
}

type LinkedList struct {
	head   *ListNode
	length int
}

func (l *LinkedList) Prepend(value int) {
	newNode := new(ListNode)
	newNode.Val = value // create newNode with desired value
	// head ->  nil
	// head -> newNode -> nil
	old := l.head
	l.head = newNode
	newNode.Next = old
	l.length++
}

func (l *LinkedList) Remove(value int) {
	if l.length == 0 {
		fmt.Println("Nothing to remove...")
		return
	}
	// value to delete is at head
	if l.head.Val == value {
		l.head = l.head.Next
		l.length--
	}
	// value to delete is in middle or at last
	curr := l.head.Next
	prev := curr
	for curr.Val != value && curr.Next != nil {
		curr = curr.Next
		prev = prev.Next
	}
	prev.Next = curr.Next
	l.length--

}

func (l *LinkedList) PrintList() {
	var builder strings.Builder
	curr := l.head
	for curr != nil {
		builder.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			builder.WriteString(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println(builder.String())
}

//func main() {
//	l := LinkedList{}
//	l.Prepend(1)
//	l.Prepend(2)
//	l.Prepend(3)
//	l.Prepend(4)
//
//	l.PrintList()
//	l.Remove(3)
//	l.PrintList()
//
//}
