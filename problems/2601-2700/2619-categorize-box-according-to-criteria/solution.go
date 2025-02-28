func categorizeBox(length int, width int, height int, mass int) string {
    volume := length * width * height
    pow9 := int(math.Pow(float64(10), float64(9)))
    pow4 := int(math.Pow(float64(10), float64(4)))
    bulky, heavy := false, false


    if volume >= pow9 || length >= pow4 || width >= pow4 || height >= pow4 {
        bulky = true
    }

    if mass >= 100 {
        heavy = true
    }

    if bulky && heavy {
        return "Both"
    } 
    if bulky {
        return "Bulky"
    }
    if heavy {
        return "Heavy"
    }
    return "Neither"
}