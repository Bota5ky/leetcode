package shu_ju_liu_zhong_de_zhong_wei_shu_lcof

// MedianFinder 剑指 Offer 41. 数据流中的中位数 https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
type MedianFinder struct {
	maxheap []int
	minheap []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (t *MedianFinder) AddNum(num int) {
	if len(t.minheap) == 0 || num >= t.minheap[0] {
		insertNum(&t.minheap, num)
	} else {
		insertNum(&t.maxheap, num)
	}
	if len(t.minheap) < len(t.maxheap) {
		tmp := t.maxheap[len(t.maxheap)-1]
		t.maxheap = t.maxheap[:len(t.maxheap)-1]
		t.minheap = append([]int{tmp}, t.minheap...)
	} else if len(t.minheap)-len(t.maxheap) > 1 {
		t.maxheap = append(t.maxheap, t.minheap[0])
		t.minheap = t.minheap[1:]
	}
}

func (t *MedianFinder) FindMedian() float64 {
	if (len(t.maxheap)+len(t.minheap))%2 == 1 {
		return float64(t.minheap[0])
	}
	return (float64(t.minheap[0]) + float64(t.maxheap[len(t.maxheap)-1])) / 2.0
}

func insertNum(nums *[]int, target int) {
	*nums = append(*nums, target)
	for i := len(*nums) - 2; i >= 0; i-- {
		if (*nums)[i] > (*nums)[i+1] {
			(*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
		} else {
			break
		}
	}
}
