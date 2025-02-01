package main

import (
	"io"
	"os"

	"fmt"

	"github.com/mattn/go-isatty"
)

type CCatPrinter interface {
	Print(r io.Reader, w io.Writer) error
}

type AutoColorPrinter struct {
	ColorPalettes ColorPalettes
}

func (a AutoColorPrinter) Print(r io.Reader, w io.Writer) error {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return ColorPrinter{a.ColorPalettes}.Print(r, w)
	} else {
		return PlainTextPrinter{}.Print(r, w)
	}
}

type ColorPrinter struct {
	ColorPalettes ColorPalettes
}

func (c ColorPrinter) Print(r io.Reader, w io.Writer) error {
	return CPrint(r, w, c.ColorPalettes)
}

type PlainTextPrinter struct {
}

func (p PlainTextPrinter) Print(r io.Reader, w io.Writer) error {
	_, err := io.Copy(w, r)
	return err
}

type HtmlPrinter struct {
	ColorPalettes ColorPalettes
}

func (c HtmlPrinter) Print(r io.Reader, w io.Writer) error {
	return HtmlPrint(r, w, c.ColorPalettes)
}

func CCat(fname string, p CCatPrinter, w io.Writer) error {
	var r io.Reader

	if fname == readFromStdin {
		r = os.Stdin

	} else {
		file, err := os.Open(fname)
		if err != nil {
		    return err
		}
		defer func() {
		    _ = file.Close() // Ensuring no accidental panic
		}()

		fi, err := file.Stat()
		if err != nil {
		    return err
		}

		if fi.Mode().IsDir() {
			return fmt.Errorf("%s is a directory", file.Name())
		}

		r = file
	}

	return p.Print(r, w)
}
