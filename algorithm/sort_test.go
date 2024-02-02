package algorithm

// People test mock data
type people struct {
	Name string
	Age  int
}

// PeopleAageComparator sort people slice by age field
type peopleAgeComparator struct{}

// Compare implements github.com/duke-git/lancet/v2/constraints/constraints.go/Comparator
func (pc *peopleAgeComparator) Compare(v1 any, v2 any) int {
	p1, _ := v1.(people)
	p2, _ := v2.(people)

	//ascending order
	if p1.Age < p2.Age {
		return -1
	} else if p1.Age > p2.Age {
		return 1
	}
	return 0
}

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	//ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}
