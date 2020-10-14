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
