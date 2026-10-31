package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//созадем загушку
	dummy := &ListNode{}
	//создаем хвост в этой заглушке
	tail := dummy

	//пока один из списков не станет nil
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next

	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next

}
func main() {
	//заполняем первый список
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}

	//заполняем второй список
	list2 := &ListNode{Val: 2}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 6}

	//объединяем два списка
	list3 := mergeTwoLists(list1, list2)

	//так как список это указатель на структу *ListNode то обыным принтом его не вывести
	var data []int
	for cur := list3; cur != nil; cur = cur.Next {
		data = append(data, cur.Val)
	}
	fmt.Println(data)

}
