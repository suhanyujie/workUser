package stringx

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Str(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
