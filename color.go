package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

const esc = "\033["

type ColorCodes map[string]string

func (c ColorCodes) String() string {
	cc := make([]string, 0, len(c))
	for k := range c {
	    if k != "" {
	        cc = append(cc, k)
	    }
	}
	sort.Strings(cc)

	var s []string
	for _, ss := range cc {
		s = append(s, Colorize(ss, ss))
	}

	return strings.Join(s, ", ")
}

var colorCodes = ColorCodes{
	"":          "",
	"reset":     esc + "39;49;00m",
	"bold":      esc + "01m",
	"faint":     esc + "02m",
	"standout":  esc + "03m",
	"underline": esc + "04m",
	"blink":     esc + "05m",
	"overline":  esc + "06m",
}

func init() {
	darkColors := []string{
		"black",
		"darkred",
		"darkgreen",
		"brown",
		"darkblue",
		"purple",
		"teal",
		"lightgray",
	}

	lightColors := []string{
		"darkgray",
		"red",
		"green",
		"yellow",
		"blue",
		"fuchsia",
		"turquoise",
		"white",
	}

	for i, x := 0, 30; i < len(darkColors); i, x = i+1, x+1 {
		colorCodes[darkColors[i]] = fmt.Sprintf("%s%dm", esc, x)
		colorCodes[lightColors[i]] = fmt.Sprintf("%s%d;01m", esc, x)
	}

	colorCodes["darkteal"] = colorCodes["turquoise"]
	colorCodes["darkyellow"] = colorCodes["brown"]
	colorCodes["fuscia"] = colorCodes["fuchsia"]
	colorCodes["white"] = colorCodes["bold"]
}

/*
	Format ``text`` with a color and/or some attributes::

		color       normal color
		*color*     bold color
		_color_     underlined color
		+color+     blinking color
*/
func Colorize(attr, text string) string {
	if attr == "" {
		return text
	}

	result := new(bytes.Buffer)

	if strings.HasPrefix(attr, "+") && strings.HasSuffix(attr, "+") {
		result.WriteString(colorCodes["blink"])
		attr = strings.Trim(attr, "+")
	}

	if strings.HasPrefix(attr, "*") && strings.HasSuffix(attr, "*") {
		result.WriteString(colorCodes["bold"])
		attr = strings.Trim(attr, "*")
	}

	if strings.HasPrefix(attr, "_") && strings.HasSuffix(attr, "_") {
		result.WriteString(colorCodes["underline"])
		attr = strings.Trim(attr, "_")
	}

	result.WriteString(colorCodes[attr])
	result.WriteString(text)
	result.WriteString(colorCodes["reset"])

	return result.String()
}
