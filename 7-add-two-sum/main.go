package main

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

func main() {
	values := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	values2 := []int{5, 6, 4}
	l1 := addToList(values)
	l2 := addToList(values2)

	n := addTwoNumbers(l1, l2)

	for n != nil {
		fmt.Printf("%d ->", n.Val)

		if n.Next == nil {
			break
		}

		n = n.Next
	}
	fmt.Println("nil")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Int := getBigIntegerOfListNode(l1)
	l2Int := getBigIntegerOfListNode(l2)
	result := new(big.Int)
	l3Int := result.Add(l1Int, l2Int)
	l3IntSlice := strings.Split(l3Int.String(), "")
	initNo, _ := strconv.Atoi(l3IntSlice[len(l3IntSlice)-1])
	head := &ListNode{
		Val: initNo,
	}
	current := head

	for i := len(l3IntSlice) - 2; i >= 0; i-- {
		val, _ := strconv.Atoi(l3IntSlice[i])
		newNode := &ListNode{
			Val: val,
		}
		current.Next = newNode
		current = newNode
	}

	return head
}

func getBigIntegerOfListNode(ln *ListNode) *big.Int {
	var lnSlice []string
	for ln != nil {
		lnSlice = append(lnSlice, strconv.Itoa(ln.Val))

		if ln.Next == nil {
			break
		}

		ln = ln.Next
	}

	slices.Reverse(lnSlice)
	bigNum := new(big.Int)
	bigInt, _ := bigNum.SetString(strings.Join(lnSlice, ""), 10)

	return bigInt
}

func addToList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, val := range values[1:] {
		newNode := &ListNode{Val: val}
		current.Next = newNode
		current = newNode
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
