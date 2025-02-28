func maxSatisfied(customers []int, grumpy []int, X int) int {
	max := 0
	satisfied := 0
	temp := 0

	for i := 0; i < X; i++ {
		satisfied, temp = AddToSatisfiedOrTemp(customers, grumpy, i, satisfied, temp)
	}

	max = temp

	for i := X; i < len(customers); i++ {
		if grumpy[i - X] == 1 {
			temp -= customers[i - X]
		}

		satisfied, temp = AddToSatisfiedOrTemp(customers, grumpy, i, satisfied, temp)
        
		max = int(math.Max(float64(max), float64(temp)))
	}

	return satisfied + max
}

func AddToSatisfiedOrTemp(customers []int, grumpy []int, i int, satisfied int, temp int) (int, int) {
	if grumpy[i] == 0 {
		satisfied += customers[i]
	} else {
		temp += customers[i]
	}
    
	return satisfied, temp
}