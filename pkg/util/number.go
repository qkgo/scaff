package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func UniformityRandString(total int, input string) []int {
	allStr := strings.Split(input, ",")
	minmaxStr := strings.Split(allStr[0], "-")
	medianStr := allStr[1]
	println(minmaxStr)
	min, err := strconv.Atoi(minmaxStr[0])
	if err != nil {
		println(err.Error())
		return nil
	}
	max, err := strconv.Atoi(minmaxStr[1])
	if err != nil {
		println(err.Error())
		return nil
	}
	median, err := strconv.Atoi(medianStr)
	if err != nil {
		println(err.Error())
		return nil
	}
	return UniformityRand(total, min, max, median)
}

func RandString(input string) int {
	allStr := strings.Split(input, ",")
	minmaxStr := strings.Split(allStr[0], "-")
	min, err := strconv.Atoi(minmaxStr[0])
	if err != nil {
		println(err.Error())
		return -1
	}
	max, err := strconv.Atoi(minmaxStr[1])
	if err != nil {
		println(err.Error())
		return -1
	}
	return RandNumber(min, max)
}

func UniformityRand(total int, min int, max int, median int) []int {
	return UniformityRandomNumber(total, min, max, median, 0)
}

func UniformityRandomNumber(total int, min int, max int, median int, average int) []int {
	if average != 0 {
		println("not support now")
		return nil
	}
	println("total:", total, "min", min, "max", max, "median", median)
	if total == 1 {
		return []int{rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min) + min}
	}
	if total == 2 {
		if median == min || median == max {
			return []int{min, max}
		}
		randBase := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min) + min
		randBase2 := median*2 - (randBase)
		return []int{randBase, randBase2}
	}
	resultArrSet := []int{}
	if median >= max || median <= min {
		println("UniformityRandomNumber.median setting error.reset")
		median = (max-min)/2 + min
	}

	odd := total % 2
	if odd == 0 {
		odd = 2
		resultArrSet = append(resultArrSet, median, median)
	} else {
		resultArrSet = append(resultArrSet, median)
	}
	singleRandNumber := (total - odd) / 2
	for i := 0; i < singleRandNumber; i++ {
		rand.Seed(time.Now().Unix())
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		indexNumber := median - min
		if indexNumber < 1 {
			continue
		}
		randRes := r.Intn(indexNumber) + min
		resultArrSet = append(resultArrSet, randRes)
	}
	for i := 0; i < singleRandNumber; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		indexNumber := max - median
		if indexNumber < 1 {
			continue
		}
		randRes := r.Intn(indexNumber) + median
		resultArrSet = append(resultArrSet, randRes)
	}
	return resultArrSet
}

func RandNumber(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
