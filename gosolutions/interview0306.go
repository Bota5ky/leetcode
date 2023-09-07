package gosolutions

// AnimalShelf AnimalShelf
// https://leetcode-cn.com/problems/animal-shelter-lcci/
type AnimalShelf struct {
	cat []int
	dog []int
}

// Constructor8 Constructor8
func Constructor8() AnimalShelf {
	return AnimalShelf{}
}

// Enqueue Enqueue
func (t *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 { //is cat
		t.cat = append(t.cat, animal[0])
	} else {
		t.dog = append(t.dog, animal[0])
	}
}

// DequeueAny DequeueAny
func (t *AnimalShelf) DequeueAny() []int {
	if len(t.cat) == 0 && len(t.dog) == 0 {
		return []int{-1, -1}
	} else if len(t.cat) == 0 || len(t.dog) > 0 && t.cat[0] > t.dog[0] {
		temp := t.dog[0]
		t.dog = t.dog[1:]
		return []int{temp, 1}
	}
	temp := t.cat[0]
	t.cat = t.cat[1:]
	return []int{temp, 0}
}

// DequeueDog DequeueDog
func (t *AnimalShelf) DequeueDog() []int {
	if len(t.dog) == 0 {
		return []int{-1, -1}
	}
	temp := t.dog[0]
	t.dog = t.dog[1:]
	return []int{temp, 1}
}

// DequeueCat DequeueCat
func (t *AnimalShelf) DequeueCat() []int {
	if len(t.cat) == 0 {
		return []int{-1, -1}
	}
	temp := t.cat[0]
	t.cat = t.cat[1:]
	return []int{temp, 0}

}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
