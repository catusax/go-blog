package normalize

import "bytes"

// LinesToLF 把windows和mac的换行符格式统一转为unix格式
func LinesToLF(b *[]byte) {
	// Win -> Unix: replace CR LF with LF & remove BOM
	*b = bytes.ReplaceAll(*b, []byte("\uFEFF"), []byte(""))
	*b = bytes.ReplaceAll(*b, []byte("\r\n"), []byte("\n"))
	// Mac -> Unix: replace CF with LF
	*b = bytes.ReplaceAll(*b, []byte("\r"), []byte("\n"))
}
