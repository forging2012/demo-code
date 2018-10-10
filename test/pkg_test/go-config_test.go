package pkg

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/encoder/toml"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"testing"
)

func TestReadConfig(t *testing.T) {
	conf := config.NewConfig()

	enc := toml.NewEncoder()

	err := conf.Load(file.NewSource(file.WithPath("../../tmp_files/conf.yaml"), source.WithEncoder(enc)))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(enc.String())

}
