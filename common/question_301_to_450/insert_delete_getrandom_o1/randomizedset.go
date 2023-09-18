package insert_delete_getrandom_o1

import "math/rand"

// RandomizedSet 380. O(1) 时间插入、删除和获取随机元素 https://leetcode.cn/problems/insert-delete-getrandom-o1/
type RandomizedSet struct {
	idx  map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{idx: make(map[int]int)}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.idx[val]; !ok {
		r.nums = append(r.nums, val)
		r.idx[val] = len(r.nums) - 1
		return !ok
	}
	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.idx[val]; ok {
		index := r.idx[val]
		r.idx[r.nums[len(r.nums)-1]] = index
		r.nums[index] = r.nums[len(r.nums)-1]
		delete(r.idx, val)
		r.nums = r.nums[:len(r.nums)-1]
		return ok
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
