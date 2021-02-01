func sortColors(nums []int) {
	white := 0
	blue := 0
	counter := 0
	
	for _, num := range nums {
		if num == 0 {
			nums[counter] = 0
			counter ++
		} else if num == 1 {
			white++
		} else {
			blue ++
		}
	}

	for i := 0; i < white; i++ {
		nums[counter] = 1
		counter ++
	}
    fmt.Println(nums)
	for i := 0; i < blue; i++ {
		nums[counter] = 2
		counter ++
	}
}