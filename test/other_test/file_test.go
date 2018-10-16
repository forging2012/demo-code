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

func TestFileExist(t *testing.T) {
	fmt.Println(Exists("homeworkTest"))
}
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
