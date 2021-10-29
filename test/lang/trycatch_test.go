package lang

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func InlineProcess() int {
	defer func() int {
		if err := recover(); err != nil {
			return 1
		}
		return 2
	}()
	return 3
}

func TestDeferFunction(t *testing.T) {
	resultInt := InlineProcess()
	println(resultInt)
}

func InlineProcessPanicNumber() int {
	defer func() int {
		if err := recover(); err != nil {
			return 1
		}
		return 2
	}()

	panic(5)
	return 3
}

func TestDeferPanic(t *testing.T) {
	resultInt := InlineProcessPanicNumber()
	println(resultInt)
}

func InlineProcessPanicError() int {
	defer func() int {
		if err := recover(); err != nil {
			fmt.Println("run 1")
			return 1
		} else {
			fmt.Println("run 2")
			return 2
		}
	}()

	panic(errors.New("this is error"))
	return 3
}

func TestDeferPanicError(t *testing.T) {
	resultInt := InlineProcessPanicError()
	println(resultInt)
}

func InlineProcessPanicErrorRight() (resultInt int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("run 1")
			resultInt = 1
		} else {
			fmt.Println("run 2")
			resultInt = 2
		}
	}()
	resultInt = 3
	panic(errors.New("this is error"))
	return resultInt
}

func TestDeferTrans(t *testing.T) {
	resultInt := InlineProcessPanicErrorRight()
	println(resultInt)
}

func InlineProcessStdResult() (resultInt int) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("run 1")
			resultInt = 1
		} else {
			//fmt.Println("run 2")
			resultInt = 2
		}
	}()
	resultInt = 3
	panic(5)
	return resultInt
}

func OnlyDefer() (resultInt int) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("run 1")
			resultInt = 1
		} else {
			//fmt.Println("run 2")
			resultInt = 2
		}
	}()
	resultInt = 3
	return resultInt
}
func InlineStdResult() (resultInt int) {
	resultInt = 3
	return resultInt
}

func TestDeferStdResult(t *testing.T) {
	resultInt := InlineProcessStdResult()
	println(resultInt)
}

func TestDeferFunctionPerformances(t *testing.T) {
	var t1, t2, t3 int
	t.Run("defer test panic", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			InlineProcessStdResult()
		}
		t1 = time.Now().Nanosecond() - start.Nanosecond()
		t.Logf("defer panic test use time %d ns", t1)
		t.Logf("defer panic test use time %f Microsecond", float64(t1)/1000.0)
		t.Logf("defer panic test use time %f Millisecond", float64(t1)/1000000.0)
	})
	time.Sleep(time.Millisecond * 500)
	t.Run("defer test", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			OnlyDefer()
		}
		t2 = time.Now().Nanosecond() - start.Nanosecond()
		t.Logf("defer test use time %d ns", t2)
	})
	time.Sleep(time.Millisecond * 500)
	t.Run("defer test", func(t *testing.T) {
		start1 := time.Now()
		for i := 0; i < 1000000; i++ {
			InlineStdResult()
		}
		t3 = time.Now().Nanosecond() - start1.Nanosecond()
		t.Logf("simple test use time %d ns", t3)
	})
	time.Sleep(time.Second * 1)
	t.Run("speed rate", func(t *testing.T) {
		t.Logf("speed rate %f and %f", float64(t1)/float64(t2), float64(t2)/float64(t3))
	})
}
