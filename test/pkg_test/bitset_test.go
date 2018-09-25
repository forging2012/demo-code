package pkg_test

import (
	"encoding/base64"
	"fmt"
	"github.com/dncmn/bitset"
	"testing"
)

func TestJsonBitset(t *testing.T) {
	b := &bitset.BitSet{}
	b.Set(uint(3))
	data, err := b.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
	data, err = base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}
