package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13byte(sb byte) byte {
    int_s := int(sb)
    
    if int_s >= int('A') && int_s <= int('M') || int_s >= int('a') && int_s <= int('m') {
		sb += 13
    } else {
		sb -= 13
	}

	return sb
}

func (reader rot13Reader) Read(stream []byte) (n int, err error) {
    length, err := reader.r.Read(stream)
    for i := 0; i < length; i++ {
        stream[i] = rot13byte(stream[i])
    }

    return length, err
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

