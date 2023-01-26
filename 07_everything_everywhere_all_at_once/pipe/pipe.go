package pipe

import (
	"fmt"
	"io"
	"log"
	"os"
)

type PipeExample struct {
}

func (p PipeExample) Demo() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "Hello there\n")
		w.Close()
	}()

	_, err := io.Copy(os.Stdout, r)

	if err != nil {
		log.Fatal(err)
	}
}
