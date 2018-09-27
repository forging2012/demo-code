package other

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
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

func TestMD5Encode(t *testing.T) {
	m := md5.New()
	s := "helloadfadfasdadfadfadsfaasdfasdfa"
	io.WriteString(m, s)
	bt := m.Sum(nil)
	fmt.Println(hex.EncodeToString(bt))
}

func TestSha1Encode(t *testing.T) {
	s := sha1.New()
	io.WriteString(s, "helldfdfaskdfjalksdfjalasfasdo")
	s.Sum(nil)
	//fmt.Println(s)
	//fmt.Println(hex.EncodeToString(s.Sum(nil)), len(hex.EncodeToString(s.Sum(nil))))
	enStr := hex.EncodeToString(s.Sum(nil))
	fmt.Println("encodeStr=", enStr)
	byt, err := hex.DecodeString(enStr)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(byt, string(byt))
}

func TestSwitchDemo(t *testing.T) {
	num := 40

	ok := GetBoolean(num)
	fmt.Println(ok)
}

func GetBoolean(num int) (ok bool) {
	switch {
	case num < 20:
		fmt.Println("num<20")
		ok = true
	case num < 50:
		fmt.Println("num<50")
		ok = true
	}

	return
}

type St struct {
	Name string
	Age  int
}

func (this *St) Play() {
	return
}

func TestBase64(t *testing.T) {
	str := "你好"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
}
