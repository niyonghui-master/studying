package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表头
type HeadLinkNode struct {
	Length int
	Node   *ListNode
}

func New(val int) *HeadLinkNode {
	return &HeadLinkNode{Length: 1, Node: &ListNode{Val: val, Next: nil}}
}

//增加末尾节点
func (h *HeadLinkNode) Add(node int) {
	l := h.Node
	for {
		if l.Next == nil {
			newNode := &ListNode{Val: node, Next: nil}
			l.Next = newNode
			break
		} else {
			l = l.Next
		}
	}
	h.Length++
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	L1Arr := GetVal(l1)
	L2Arr := GetVal(l2)

	if len(L1Arr) >= len(L2Arr) {
		return Add(L1Arr, L2Arr)
	} else {
		return Add(L2Arr, L1Arr)
	}
}

func Add(L1Arr, L2Arr []int) *ListNode {
	var decade int
	var Nodes *HeadLinkNode
	for i := 0; i < len(L1Arr); i++ {
		L1Num := L1Arr[i]
		L2Num := 0
		if i < len(L2Arr) {
			L2Num = L2Arr[i]
		}
		sum := L1Num + L2Num + decade
		decade = 0
		if sum >= 10 {
			decade = 1
			sum = sum - 10
		}
		if i == 0 {
			Nodes = New(sum)
		} else {
			Nodes.Add(sum)
		}

	}

	if decade > 0 {
		Nodes.Add(decade)
	}

	if Nodes == nil {
		return nil
	}

	return Nodes.Node
}

func GetVal(l *ListNode) (LArr []int) {
	for {
		if l == nil {
			break
		}
		LArr = append(LArr, l.Val)
		l = l.Next
	}

	return
}

func main() {
	L1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:8,
			Next:nil,
		},
	}

	L2 := &ListNode{
		Val: 0,
		Next: nil,
	}
	addTwoNumbers(L1, L2)
	fmt.Println(addTwoNumbers(L1, L2).Val)
	fmt.Println(addTwoNumbers(L1, L2).Next.Val)
	//fmt.Println(addTwoNumbers(L1, L2).Next.Next.Val)
	//fmt.Println(addTwoNumbers(L1, L2).Next.Next.Next.Val)
}
