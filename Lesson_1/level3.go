package Lesson_1

func IsPrime(num int) bool {
	var i, j int
	j = 0
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			j++
		}
	}
	if j == 2 {
		return true
	} else {
		return false
	}
}
