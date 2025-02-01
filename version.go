package main

import (
	"fmt"
	"io"
)

var Version = "1.2.0"

func displayVersion(w io.Writer) {
    _, _ = fmt.Fprintf(w, "ccat version %s\n", Version)
}

