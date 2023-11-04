package main

import (
	"fmt"
	"golang_learner/Lesson_1"
	"math/rand"
	"time"
)

func main() {
	var a, b, d, e int
	var c float64
	rand.Seed(time.Now().UnixNano())
	a, b, c, d, e = rand.Intn(100), rand.Intn(100), float64(rand.Intn(100)), rand.Intn(100), rand.Intn(100)
	ret1 := Lesson_1.Sum(a, b)
	fmt.Println("Num a is: ", a, ", Num b is: ", b, ", Sum is: ", ret1)
	ret2 := Lesson_1.CircleArea(c)
	fmt.Println("R is: ", c, ", Area is: ", ret2)
	ret3 := Lesson_1.IsPrime(d)
	fmt.Println("Num is: ", d, ", Result is: ", ret3)
	retX := Lesson_1.Dichotomy(e)
	fmt.Println("Direct Num is: ", retX)
}
