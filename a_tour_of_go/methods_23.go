package main

import (
	"io"
	"os"
	"strings"
)

func SwapRuneRot13(b byte) byte {
    switch {
	case 65 <= b && b < 65+13:
        return b + 13
	case b <= 90:
        return b - 13
    case 97 <= b && b < 97+13:
        return b + 13
	case b <= 122:
        return b - 13
    default:
        return b
    }
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = rot.r.Read(bytes)
	for i := range bytes {
		bytes[i] = SwapRuneRot13(bytes[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// Youcrackedthecode
