package main

import "testing"

func Test_ColorPalette_Set(t *testing.T) {
	palettes := ColorPalettes{
		stringKind: "blue",
	}

	ok := palettes.Set("foo", "bar")
	if ok {
		t.Errorf("setting color code foo should not be ok")
	}

	ok = palettes.Set("String", "baz")
	if !ok {
		t.Fatalf("failed to set color code 'String'")
	}

	if palettes[stringKind] != "baz" {
	    t.Errorf("expected color code of String to be 'baz', got '%s'", palettes[stringKind])
	}

}

func TestColorize(t *testing.T) {
	cases := []struct {
		Color, Output string
	}{
		{
			Color:  "",
			Output: "hello",
		},

		{
			Color:  "blue",
			Output: "\033[34;01mhello\033[39;49;00m",
		},
		{
			Color:  "_blue_",
			Output: "\033[04m\033[34;01mhello\033[39;49;00m",
		},
		{
			Color:  "bold",
			Output: "\033[01mhello\033[39;49;00m",
		},
	}

	for _, tc := range cases {
	    t.Run(tc.Color, func(t *testing.T) {
	        actual := Colorize(tc.Color, "hello")
	        if actual != tc.Output {
	            t.Errorf("for color %q, expected %q but got %q", tc.Color, tc.Output, actual)
	        }
	    })
	}
}

func TestColorizeMultiByte(t *testing.T) {
	cases := []struct {
		Color, Output string
	}{
		// Japanese
		{
			Color:  "",
			Output: "こんにちは",
		},

		{
			Color:  "blue",
			Output: "\033[34;01mこんにちは\033[39;49;00m",
		},
		{
			Color:  "_blue_",
			Output: "\033[04m\033[34;01mこんにちは\033[39;49;00m",
		},
		{
			Color:  "bold",
			Output: "\033[01mこんにちは\033[39;49;00m",
		},
	}

	for _, tc := range cases {
		actual := Colorize(tc.Color, "こんにちは")
		if actual != tc.Output {
			t.Errorf(
				"Color: %#v\n\nOutput: %#v\n\nExpected: %#v",
				tc.Color,
				actual,
				tc.Output)
		}
	}
}
