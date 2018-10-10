package other

import (
	"demo-code/utils"
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	fs, err := utils.GetAllFiles(".")
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range fs {
		fmt.Println(f)
	}
}

func TestGetPwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(dir)
}
