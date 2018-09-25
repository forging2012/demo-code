package other_test

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test2(t *testing.T) {
	name := flag.String("name", "helloKitty", "hello")
	flag.Parse()
	//fmt.Println(flag.Parsed())
	fmt.Println("name=", *name)
	others := flag.Args()
	fmt.Println("others=", others)
}

var (
	period = flag.Duration("period", 3*time.Second, "sleep period")
)

func TestGetPeriod(t *testing.T) {
	flag.Parse()
	fmt.Printf("Sleeping for %v....\n", *period)
}
func TestSayHello(t *testing.T) {
	flag.String("env", "test", "for test")
	err := flag.Set("env", "test")
	if err != nil {
		t.Error(err)
		return
	}
	flag.Parse()
	fmt.Println(os.Getenv("env"))
	fmt.Println("helloWrold")
	fmt.Println("helloWrold")
	fmt.Println("helloWrold")
	fmt.Println("helloWrold")
	fmt.Println("helloWrold")
}
