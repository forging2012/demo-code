package other

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestNumConvert(t *testing.T) {
	hex := "AB"
	fmt.Println(HexDec(hex))
}

// 16进制到10进制
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

func TestStructTomap(t *testing.T) {
	s := Stu{
		ID:   "123456",
		Name: "",
		Age:  0,
	}

	bte, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
		return
	}

	m := make(map[string]interface{}, 0)
	err = json.Unmarshal(bte, &m)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m)

	fmt.Println(s, bte)

}

type Stu struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
