package ch2p2

import "sedgewick/ll"

// Source: https://stackoverflow.com/a/56240818/605022
// > Time Complexity: O(nLogn) n = number of nodes.
// > Each iteration through the linked list doubles
// > the size of the sorted smaller linked lists.
// > For example, after the first iteration, the linked
// > list will be sorted into two halves. After the
// > second iteration, the linked list will be sorted
// > into four halves. It will keep sorting up to the
// > size of the linked list. This will take O(logn)
// > doublings of the small linked lists' sizes to
// > reach the original linked list size. The n in
// > nlogn is there because each iteration of the
// > linked list will take time proportional to the
// > number of nodes in the originial linked list.

func LLMergeSort[T LessThan](head *ll.Node[T]) *ll.Node[T] {
	if head == nil || head.Next == nil {
		return head
	}
	middle := llGetMiddle(head)
	afterMiddle := middle.Next
	middle.Next = nil
	return llMerge(LLMergeSort(head), LLMergeSort(afterMiddle))
}
func llMerge[T LessThan](a, b *ll.Node[T]) (result *ll.Node[T]) {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Value <= b.Value {
		result = a
		result.Next = llMerge(a.Next, b)
	} else {
		result = b
		result.Next = llMerge(a, b.Next)
	}
	return result
}
func llGetMiddle[T LessThan](head *ll.Node[T]) *ll.Node[T] {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return slow
}
