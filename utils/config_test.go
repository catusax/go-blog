package utils_test

import (
	"blog/utils"
	"testing"
)

func TestInit(t *testing.T) {
	if utils.C.User.Username != "root" {
		t.Error("读取失败", utils.C)
	}
}

func TestWriteConf(t *testing.T) {
	if err := utils.WriteConf([]byte(`{"Port": "90","user": {"username": "coolrc"}}`)); err != nil {
		t.Error("失败", err)
	}

}
