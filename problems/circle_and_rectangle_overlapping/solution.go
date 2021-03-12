func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {

	if x_center < x2 && x_center > x1 && y_center < y2 && y_center > y1 {
		// Checking whether the circle is in the middle of the square
		return true
	}

	if ((y1 <= y_center+radius && y1 >= y_center) || (y2 >= y_center-radius && y2 <= y_center)) &&
		x_center >= x1 && x_center <= x2 {
		return true
	}

	if y_center <= y2 && y_center >= y1 && ((x_center-radius <= x2 && x_center >= x2) ||
		(x_center+radius >= x1 && x_center <= x1)) {
		return true
	}

	if powerOfTwo(x_center-x1)+powerOfTwo(y_center-y1) <= powerOfTwo(radius) ||
		powerOfTwo(x_center-x1)+powerOfTwo(y_center-y2) <= powerOfTwo(radius) ||
		powerOfTwo(x_center-x2)+powerOfTwo(y_center-y1) <= powerOfTwo(radius) ||
		powerOfTwo(x_center-x2)+powerOfTwo(y_center-y2) <= powerOfTwo(radius) {
		return true
	}

	return false
}

func powerOfTwo(n int) int {
	return n * n
}