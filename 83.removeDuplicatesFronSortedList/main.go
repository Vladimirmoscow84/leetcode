package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesFromSortedLists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
func main() {
	//заполняем список
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 1}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next.Next = &ListNode{Val: 5}

	//убираем дубликаты из списка
	result := deleteDuplicatesFromSortedLists(list)

	//так как список это указатель на структу *ListNode то обыным принтом его не вывести
	var data []int
	for cur := result; cur != nil; cur = cur.Next {
		data = append(data, cur.Val)
	}
	fmt.Println(data)
}
