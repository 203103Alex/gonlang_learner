package Lesson_1

func Dichotomy(dirNum int) int {
	a, b := 0, 100
	for {
		temp := (a + b) / 2
		if temp < dirNum {
			a++
		} else if temp > dirNum {
			b--
		} else {
			return temp
		}
	}
}
