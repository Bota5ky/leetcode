package question_201_to_300

//和offer41相同
//https://leetcode.cn/problems/find-median-from-data-stream/
import (
	"container/heap"
)

// MaxHeap 最大堆
type MaxHeap []int

// MinHeap 最小堆
type MinHeap []int

// Len Len
func (h MaxHeap) Len() int { return len(h) }

// Less Less
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

// Swap Swap
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push Push
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop Pop
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Len Len
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push Push
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop Pop
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder MedianFinder
type MedianFinder struct {
	lower MaxHeap
	upper MinHeap
}

//Constructor5 Constructor5
/** initialize your data structure here. */
func Constructor5() MedianFinder {
	return MedianFinder{make([]int, 0), make([]int, 0)}
}

// AddNum AddNum
func (t *MedianFinder) AddNum(num int) { // len(lower) >= len(upper)
	if len(t.lower) == 0 || t.lower[0] >= num {
		heap.Push(&t.lower, num)
	} else {
		heap.Push(&t.upper, num)
	}

	if len(t.upper) > len(t.lower) {
		v := heap.Pop(&t.upper)
		heap.Push(&t.lower, v)
	} else if len(t.lower) > len(t.upper)+1 {
		v := heap.Pop(&t.lower)
		heap.Push(&t.upper, v)
	}
}

// FindMedian FindMedian
func (t *MedianFinder) FindMedian() float64 {
	if len(t.lower) > len(t.upper) {
		return float64(t.lower[0])
	}
	return float64(t.lower[0]+t.upper[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/* MAX & MIN heap written by myself
type MedianFinder struct {
    maxheap []int
    minheap []int
}
//initialize your data structure here.
func Constructor() MedianFinder {
    return MedianFinder{}
}
func (t *MedianFinder) AddNum(num int)  {
    //插入数字
    if len(t.minheap)==0 || num>=t.minheap[0] {
        t.minheap=append(t.minheap,num)
    }else{
        t.maxheap=append(t.maxheap,num)
    }
    //整理
    heapifyMax(t.maxheap)
    heapifyMin(t.minheap)
    if len(t.maxheap)>len(t.minheap) {
        temp:=t.maxheap[0]
        t.maxheap=t.maxheap[1:]
        t.minheap=append(t.minheap,temp)
    }else if len(t.minheap)-len(t.maxheap)>1 {
        temp:=t.minheap[0]
        t.minheap=t.minheap[1:]
        t.maxheap=append(t.maxheap,temp)
    }
}
func (t *MedianFinder) FindMedian() float64 {
	heapifyMax(t.maxheap)
    heapifyMin(t.minheap)
    l:=len(t.minheap)+len(t.maxheap)
    if l%2==1 {
        return float64(t.minheap[0])
    }
    return float64(t.minheap[0]+t.maxheap[0])/2.0
}
func heapifyMax(nums []int) {
    last:=len(nums)/2-1
    for i:=last;i>=0;i--{
        left:=2*i+1
        right:=2*i+2
        max:=i
        if left<len(nums) && nums[left]>nums[max] {max=left}
        if right<len(nums) && nums[right]>nums[max] {max=right}
        nums[i],nums[max]=nums[max],nums[i]
    }
}
func heapifyMin(nums []int) {
    last:=len(nums)/2-1
    for i:=last;i>=0;i--{
        left:=2*i+1
        right:=2*i+2
        min:=i
        if left<len(nums) && nums[left]<nums[min] {min=left}
        if right<len(nums) && nums[right]<nums[min] {min=right}
        nums[i],nums[min]=nums[min],nums[i]
    }
}
*/
