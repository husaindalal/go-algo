package hash


/*
	Duplicates

 */
type TwoSum struct {
	lookup map[int]int // value -> count
}

func NewTwoSum() TwoSum {
	return TwoSum{
		lookup: map[int]int{},
	}
}

func (t *TwoSum) Add(value int) {
	t.lookup[value] = t.lookup[value] + 1
}

func (t *TwoSum) Find(value int) bool {
	for k, v := range t.lookup {
		second := value - k
		if second == k {
			if v > 1 {
				return true
			}
		} else {
			if t.lookup[second] > 0 {
				return true
			}
		}
	}
	return false
}

//func twoSumDS ()

/*
	twoSum := hash.NewTwoSum()
	twoSum.Add(3)
	twoSum.Add(1)
	twoSum.Add(2)
	result := twoSum.Find(7)
	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")


*/