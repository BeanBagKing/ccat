package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type HtmlCodes map[string]string

func (c HtmlCodes) String() string {
	cc := make([]string, 0, len(c))
	for k := range c {
	    if k != "" {
	        cc = append(cc, k)
	    }
	}
	sort.Strings(cc)

	var s []string
	for _, ss := range cc {
		s = append(s, Htmlize(ss, ss))
	}

	return strings.Join(s, ", ")
}

var htmlCodes = HtmlCodes{
	"":          "",
	"reset":     `</span>`,
	"bold":      `<span class="bold">`,
	"faint":     `<span class="faint">`,
	"standout":  `<span class="standout">`,
	"underline": `<span class="underline">`,
	"blink":     `<span class="blink">`,
	"overline":  `<span class="overline">`,
}

func init() {
	darkHtmls := []string{
		"black",
		"darkred",
		"darkgreen",
		"brown",
		"darkblue",
		"purple",
		"teal",
		"lightgray",
	}

	lightHtmls := []string{
		"darkgray",
		"red",
		"green",
		"yellow",
		"blue",
		"fuchsia",
		"turquoise",
		"white",
	}

	for i, x := 0, 30; i < len(darkHtmls); i, x = i+1, x+1 {
		tag := `<span class="%s">`
		htmlCodes[darkHtmls[i]] = fmt.Sprintf(tag, darkHtmls[i])
		htmlCodes[lightHtmls[i]] = fmt.Sprintf(tag, lightHtmls[i])
	}

	htmlCodes["darkteal"] = htmlCodes["turquoise"]
	htmlCodes["darkyellow"] = htmlCodes["brown"]
	htmlCodes["fuscia"] = htmlCodes["fuchsia"]
	htmlCodes["white"] = htmlCodes["bold"]
}

func Htmlize(attr, text string) string {
	if attr == "" {
		return text
	}

	result := new(bytes.Buffer)

	if strings.HasPrefix(attr, "+") && strings.HasSuffix(attr, "+") {
		result.WriteString(htmlCodes["blink"])
		attr = strings.Trim(attr, "+")
	}

	if strings.HasPrefix(attr, "*") && strings.HasSuffix(attr, "*") {
		result.WriteString(htmlCodes["bold"])
		attr = strings.Trim(attr, "*")
	}

	if strings.HasPrefix(attr, "_") && strings.HasSuffix(attr, "_") {
		result.WriteString(htmlCodes["underline"])
		attr = strings.Trim(attr, "_")
	}

	result.WriteString(htmlCodes[attr])
	result.WriteString(text)
	result.WriteString(htmlCodes["reset"])

	return result.String()
}
