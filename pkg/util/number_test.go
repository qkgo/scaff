package util

import (
	"log"
	"testing"
)

func TestUniformityRandomNumber(t *testing.T) {
	total := 2
	max := 3000
	resIntArr := UniformityRandomNumber(total, 50, max, 51, 0)
	log.Println(resIntArr)
	log.Println(len(resIntArr) == total)
}

func TestA(t *testing.T) {
	total := 30
	resIntArr := UniformityRandString(total, "1-50,35")
	log.Println(resIntArr)
	log.Println(len(resIntArr) == total)

}
