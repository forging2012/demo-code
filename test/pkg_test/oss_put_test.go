package pkg

import (
	"demo-code/utils"
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestInitOss(t *testing.T) {
	client, err := utils.GetOss()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(client.Config.AccessKeyID)
	res, err := client.GetBucketInfo(utils.OssBucketName)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res.BucketInfo)
}

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
