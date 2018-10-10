package pkg

import (
	"demo-code/utils"
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestOssPutFile(t *testing.T) {
	bkt, err := utils.GetOssBucket(utils.OssBucketName)
	err = bkt.PutObjectFromFile("oss_put_manan_test/test1", "./test1.txt")
	if err != nil {
		t.Error(err)
		return
	}

	err = bkt.PutObjectFromFile("oss_put_manan_test/test2", "./test2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("upload success")
}

func TestOssPutDirctoryAndFiles(t *testing.T) {
	dir_path := "../other_test/homeworkTest"
	err := utils.PutFilesToOSS(dir_path)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("uplode success")
}

func TestPathBasic(t *testing.T) {
	src := "../other_test/homeworkTest/level1/hello.txt"
	dir_path := "../other_test/homeworkTest"
	idx := strings.Index(src, path.Base(dir_path))
	fmt.Println(string(([]byte(src)[idx:])))
}
