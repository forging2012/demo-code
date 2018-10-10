package other

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	fs, err := GetAllFiles(".")
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range fs {
		fmt.Println(f)
	}
}

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			//files = append(files, dirPth+PthSep+fi.Name())
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if !ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func TestGetPwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(dir)
}
