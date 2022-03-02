package util

import (
	"log"
	"testing"
)

func TestRandomDay(t *testing.T) {
	res := DeleteNumberAtArray([]int{1, 2, 3}, 2)
	log.Println(res)
}

func TestSplice(t *testing.T) {
	println([]int{1, 2}[:1])
	println([]int{1, 2}[:0])
	println([]int{1, 2}[1:])
	p := []int{1, 2, 3}
	println(p[:2-1])
	println(p[2:])
	//println(append([]int{1, 2}[:0]))
}

func TestSplice1(t *testing.T) {
	println([]int{1, 2}[:1])
	println([]int{1, 2}[:0])
	println([]int{1, 2}[1:])
	p := []int{1, 2}
	println(p[:2-1]) //删除第二个  :n-1
	println(p[2:])   // 删除第二个 n:
}

func TestSpliceZero(t *testing.T) {
	p := []int{1, 2}
	deleteNumber := 1
	arrA := p[:deleteNumber-1]
	arrB := p[deleteNumber:]
	log.Println(append(arrA, arrB...))
}
