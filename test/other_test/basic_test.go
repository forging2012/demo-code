package other

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestReflect(t *testing.T) {
	var a string
	ta := reflect.TypeOf(a)
	fmt.Println("type of a=", ta.Kind())

	var b int8
	tb := reflect.TypeOf(b)
	fmt.Println("type of b=", tb.Kind())

	var c func()
	tc := reflect.TypeOf(c)
	fmt.Println("type of c=", tc.Kind())
}

type Student struct {
	Say func(name string)
}

func TestDefineFuncFieldOfStructField(t *testing.T) {
	a := Student{}
	a.Say = func(name string) {
		fmt.Println("I am", name)
	}

	b := Student{}
	b.Say = func(name string) {
		fmt.Println("I am", name)
	}

	a.Say("a")
	b.Say("b")
}

var NowFunc = func() time.Time {
	return time.Now()
}

func TestGetTimeNow(t *testing.T) {
	fmt.Println(NowFunc().In(time.Local).Format("2006-01-02"))
	fmt.Println(reflect.TypeOf(NowFunc).Kind())
	fmt.Println(NowFunc().String())
}

var QQ func(name, age string)

func TestQQ(t *testing.T) {
	a := QQ
	a = func(name, age string) {
		fmt.Println("hello ", name, "your age is ", age)
	}
	a("tom", "18")
}
