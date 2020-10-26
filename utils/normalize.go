package utils

import "bytes"

func NormalizeLines(b []byte) []byte {
	// Win -> Unix: replace CR LF with LF & remove BOM
	b = bytes.Replace(b, []byte("\uFEFF"), []byte(""), -1)
	b = bytes.Replace(b, []byte("\r\n"), []byte("\n"), -1)
	// Mac -> Unix: replace CF with LF
	b = bytes.Replace(b, []byte("\r"), []byte("\n"), -1)
	return b
}
