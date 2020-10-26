package utils_test

import (
	"blog/utils"
	"fmt"
	"testing"
)

func TestMDParse(t *testing.T) {
	input := []byte("---\nyaml配置\n---\nmarkdown内容")
	yaml, md := utils.MDCut(input)
	fmt.Printf("%s\n%s\n", yaml, md)
}

func TestGetDescription(t *testing.T) {
	descmd := utils.GetDescription([]byte("123123<!--more-->321321"))
	fmt.Println(string(descmd))
}
