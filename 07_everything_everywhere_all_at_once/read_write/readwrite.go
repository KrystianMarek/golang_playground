package read_write

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type date struct {
	Month int
	Day   int
}

type countingWriter struct {
	n int64
	w io.Writer
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.n += int64(n)
	return n, err
}

func (d date) WriteTo(w io.Writer) (int64, error) {
	cw := &countingWriter{w: w}
	err := json.NewEncoder(cw).Encode(d)
	return cw.n, err
}

type ReadWriteExample struct {
}

func (rw ReadWriteExample) Demo() {
	d := date{89, 'z'}
	n, err := d.WriteTo(os.Stdout)
	fmt.Println(n, err)
}
