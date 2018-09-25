package main

import (
	"flag"
	"fmt"
)

func main() {

	nameStr := flag.String("name", "tomCat", "user name")
	ageStr := flag.Int("age", 18, "user age")

	var email string
	flag.StringVar(&email, "emailURL", "http://www.baidu.com", "user email")

	flag.Parse()
	others := flag.Args()
	nf := flag.Lookup("name")
	fmt.Println("nf.name=", nf.Name, " nf.Usage=", nf.Usage, " nf.Value=", nf.Value, " nf.Default=", nf.DefValue)
	fmt.Println("name=", *nameStr, "  age=", *ageStr, " email=", email, " others=", others)

}
