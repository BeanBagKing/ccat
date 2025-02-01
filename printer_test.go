package main

import (
	"bytes"
	"testing"
)

func TestCPrint(t *testing.T) {
	r := bytes.NewBufferString("hello")
	var w bytes.Buffer

	err := CPrint(r, &w, LightColorPalettes)
	if err != nil {
		t.Errorf("error should be nil, but it's %s", err)
	}

	s := w.String()
	expected := "\033[34mhello\033[39;49;00m"
	if s != expected {
	    t.Errorf("CPrint output mismatch:\nExpected: %q\nGot: %q", expected, s)
	}

}

func TestHtmlPrint(t *testing.T) {
	r := bytes.NewBufferString("hello")
	var w bytes.Buffer

	err := HtmlPrint(r, &w, LightColorPalettes)
	if err != nil {
		t.Errorf("error should be nil, but it's %s", err)
	}

	expect := `<style>
.black { color: black; }
.blink { color: blink; }
.blue { color: blue; }
.bold { color: bold; }
.brown { color: brown; }
.darkblue { color: darkblue; }
.darkgray { color: darkgray; }
.darkgreen { color: darkgreen; }
.darkred { color: darkred; }
.darkteal { color: darkteal; }
.darkyellow { color: darkyellow; }
.faint { color: faint; }
.fuchsia { color: fuchsia; }
.fuscia { color: fuscia; }
.green { color: green; }
.lightgray { color: lightgray; }
.overline { color: overline; }
.purple { color: purple; }
.red { color: red; }
.reset { color: reset; }
.standout { color: standout; }
.teal { color: teal; }
.turquoise { color: turquoise; }
.underline { color: underline; }
.white { color: white; }
.yellow { color: yellow; }
</style>
<pre>
<span class="darkblue">hello</span>
</pre>
`

	t.Run("HtmlPrint Output", func(t *testing.T) {
	    s := w.String()
		if !strings.Contains(s, `<span class="darkblue">hello</span>`) {
		    t.Errorf("HtmlPrint output does not contain expected styling. Got:\n%s", s)
		}
	})
}
