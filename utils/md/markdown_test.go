package md_test

import (
	md2 "blog/utils/md"
	"fmt"
	"testing"
)

func TestMDParse(t *testing.T) {
	input := []byte("---\nyaml配置\n---\nmarkdown内容")
	yaml, md := md2.Cut(input)
	fmt.Printf("%s\n%s\n", yaml, md)
}

func TestGetDescription(t *testing.T) {
	descmd := md2.GetDescription([]byte("123123<!--more-->321321"))
	fmt.Println(string(descmd))
}
