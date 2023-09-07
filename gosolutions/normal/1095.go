package leetcode

//https://leetcode-cn.com/problems/find-in-mountain-array/
func findInMountainArray(target int, mountainArr *MountainArray) int {
	l := mountainArr.length()
	top := findTop(mountainArr, l)
	tmp := binarySearch(mountainArr, target, 0, top, true)
	if tmp >= 0 {
		return tmp
	}
	return binarySearch(mountainArr, target, top+1, l-1, false)
}

func binarySearch(mountainArr *MountainArray, target, head, end int, isUp bool) int {
	for head <= end {
		mid := (head + end) / 2
		midVal := mountainArr.get(mid)
		if head == end && midVal != target {
			break
		}
		if midVal == target {
			return mid
		} else if midVal < target {
			if isUp {
				head = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if isUp {
				end = mid - 1
			} else {
				head = mid + 1
			}
		}
	}
	return -1
}

func findTop(mountainArr *MountainArray, l int) int {
	i, j := 0, l-1
	for i < j {
		mid := (i + j) / 2
		midVal := mountainArr.get(mid)
		preVal := mountainArr.get(mid - 1)
		nextVal := mountainArr.get(mid + 1)
		if midVal > preVal && midVal > nextVal {
			return mid
		}
		if midVal > preVal && midVal < nextVal {
			i = mid + 1
			if j-i == 1 {
				jVal := mountainArr.get(j)
				if jVal > nextVal {
					return j
				}
				return mid + 1
			}
		} else {
			j = mid - 1
			if j-i == 1 {
				iVal := mountainArr.get(i)
				if iVal > preVal {
					return i
				}
				return mid - 1
			}
		}
	}
	return i
}

//MountainArray MountainArray
type MountainArray struct {
}

func (t *MountainArray) get(index int) int { return index }
func (t *MountainArray) length() int       { return 0 }
