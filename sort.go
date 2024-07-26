package collection

import (
	"container/list"
)

// mergeSort 好像链表排序中的归并效率比较高，为O(nlogn)，但是空间开销比较大
func mergeSort[T item](lst *list.List, f func(i, j T) bool) *list.List {
	if lst.Len() <= 1 {
		return lst
	}
	// 将链表分为两半
	left, right := cut(lst)
	// 递归地对左右两半进行排序
	left = mergeSort(left, f)
	right = mergeSort(right, f)

	// 合并左右两个已排序的链表
	return merge(left, right, f)
}

// cut 将链表对半分
func cut(lst *list.List) (left, right *list.List) {
	left = list.New()
	right = lst
	if lst.Len() <= 1 {
		return
	}
	// 快慢指针，快的走两个如果快的走完了就停下，慢的就是一半了，剩下都就是另一半
	for slow, fast := lst.Front(), lst.Front(); fast != nil; fast, slow = slow.Next(), fast.Next() {
		fast = fast.Next()
		if fast == nil {
			break
		}
		left.PushBack(slow.Value)
		right.Remove(slow)
	}
	return
}

// 合并两个已排序链表
func merge[T item](left, right *list.List, f func(i, j T) bool) *list.List {

	result := list.New()

	for left.Len() > 0 && right.Len() > 0 {
		if f(left.Front().Value.(T), right.Front().Value.(T)) {
			result.PushBack(left.Front().Value)
			left.Remove(left.Front())
		} else {
			result.PushBack(right.Front().Value)
			right.Remove(right.Front())
		}
	}

	// 将剩余的元素添加到结果链表中
	if left.Len() > 0 {
		result.PushBackList(left)
	}

	if right.Len() > 0 {
		result.PushBackList(right)
	}

	return result
}
