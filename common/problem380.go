package common

import "math/rand"

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
type randomizedSet struct {
	idx  map[int]int
	nums []int
}

/** Initialize your data structure here. */
func constructor3() randomizedSet {
	return randomizedSet{idx: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *randomizedSet) Insert(val int) bool {
	if _, ok := r.idx[val]; !ok {
		r.nums = append(r.nums, val)
		r.idx[val] = len(r.nums) - 1
		return !ok
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *randomizedSet) Remove(val int) bool {
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

/** Get a random element from the set. */
func (r *randomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}

/**
 * Your randomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
