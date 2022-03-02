package util

import (
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func RandomDay(allMonthDay []string, randomNumber int) []string {
	tempMonthDayArr := allMonthDay
	returnMonthArr := []string{}
	ll := sync.Mutex{}
	defer func() []string {
		if err := recover(); err != nil {
			ll.Unlock()
		}
		return returnMonthArr
	}()
	ll.Lock()
	randomTimes := 0
	for true {
		randomTimes++
		if len(tempMonthDayArr) < 1 {
			break
		}
		if len(tempMonthDayArr) == 1 {
			returnMonthArr = append(returnMonthArr, tempMonthDayArr[0])
			break
		}
		rand.Seed(time.Now().UnixNano())
		maxLength := len(tempMonthDayArr) - 1
		randomNumberTemp := rand.Intn(maxLength)
		if randomNumberTemp == 0 {
			randomNumberTemp = 1
		}
		returnMonthArr = append(returnMonthArr, tempMonthDayArr[randomNumberTemp])
		tempMonthDayArr = append(tempMonthDayArr[:randomNumberTemp-1], tempMonthDayArr[randomNumberTemp:]...)
		if len(returnMonthArr) >= randomNumber {
			break
		}
	}
	ll.Unlock()
	deleteDuplicationMap := map[string]bool{}
	for _, row := range returnMonthArr {
		deleteDuplicationMap[row] = true
	}
	finallyReturnArray := []string{}
	for key, _ := range deleteDuplicationMap {
		finallyReturnArray = append(finallyReturnArray, key)
	}

	return finallyReturnArray
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	log.Println(reflect.TypeOf(arg))
	log.Println(reflect.ValueOf(arg))
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func DeleteNumberAtArray(listParameter interface{}, number int) interface{} {
	println(reflect.TypeOf(listParameter).String())
	println(reflect.TypeOf(listParameter).Kind())
	slice, ok := takeArg(listParameter, reflect.Slice)
	if !ok {
		return reflect.Value{}
	}
	log.Println(slice)
	list := listParameter.([]string)
	arrLen := len(list)
	if len(list) < 1 {
		return list
	}
	if number > arrLen {
		return list
	}
	if arrLen == number {
		return list[:arrLen-1]
	}
	if number == 0 {
		return list[number]
	}
	return append(list[:arrLen-1], list[arrLen:]...)
}

func DeleteNumberAtArr(number int, list ...interface{}) []interface{} {
	arrLen := len(list)
	if len(list) < 1 {
		return list
	}
	if number < arrLen {
		return list
	}
	if arrLen == number {
		return list[:arrLen-1]
	}
	if number == 0 {
		return list[1:arrLen]
	}
	return append(list[:arrLen-1], list[arrLen+1:])
}
