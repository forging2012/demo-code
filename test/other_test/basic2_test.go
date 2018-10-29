package other

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestStructTomap(t *testing.T){
	s:=Stu{
		ID:"123456",
		Name:"",
		Age:0,
	}


	bte,err:=json.Marshal(s)
	if err!=nil{
		t.Error(err)
		return
	}

	m:=make(map[string]interface{},0)
	err=json.Unmarshal(bte,&m)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(m)

	fmt.Println(s,bte)

}


type Stu struct{
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
}
