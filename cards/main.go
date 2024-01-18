package main

import (
	"fmt"
	"reflect"
	"time"
)

type MailCategrory int

const (
	Unrecognized MailCategrory = iota
	Personal
	Spam
	Social
	Advertisements
)

func copySlice(originalSlice []int) []int {
	sliceKopie := make([]int, len(originalSlice))
	copy(sliceKopie, originalSlice)
	for i, wert := range sliceKopie {
		sliceKopie[i] = wert + 2
	}
	return sliceKopie
}

func modifySlice(originalSlice []int) []int {
	for i, wert := range originalSlice {
		originalSlice[i] = wert + 2
	}
	return originalSlice
}

func (m MailCategrory) printIt() {
	fmt.Sprintf("MailCategrory: %d", m)
}

// 1. generics on Go
func isEqual[T comparable](a, b T) bool {
	return a == b
}

// 2. generics on Go
func isEqual2[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

func main() {
	nbrOfRuns := 1000000
	start := time.Now()
	for i := 0; i < nbrOfRuns; i++ {
		slice := make([]int, 10)
		_ = copySlice(slice)
	}
	duration := time.Since(start)
	fmt.Printf("Die Ausführung für copySlice dauerte %v\n", duration)
	fmt.Println("-------------------------------------")
	start = time.Now()
	for i := 0; i < nbrOfRuns; i++ {
		slice := make([]int, 10)
		_ = modifySlice(slice)
	}
	duration = time.Since(start)
	fmt.Printf("Die Ausführung für modifySlice dauerte %v\n", duration)

	/*fmt.Println(Unrecognized)
	fmt.Println(">>>>")
	Unrecognized.printIt()
	fmt.Println(">>>>")
	fmt.Println(Personal)
	fmt.Println(Spam)
	fmt.Println(Social)
	fmt.Println(Advertisements)*/

	fmt.Println(isEqual2(5, 5))             //  true
	fmt.Println(isEqual2("Hello", "World")) //  false
	fmt.Println(isEqual2(3.14, 3.14))       //  true
}
