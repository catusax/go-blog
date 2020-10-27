package config_test

import (
	"blog/utils/config"
	"testing"
)

func TestInit(t *testing.T) {
	if config.C.User.Username != "root" {
		t.Error("读取失败", config.C)
	}
}

func TestWriteConf(t *testing.T) {
	if err := config.WriteConf([]byte(`{"Port": "90","user": {"username": "coolrc"}}`)); err != nil {
		t.Error("失败", err)
	}

}
