package txtutil

import (
	"strings"
	"testing"
)

func r(s string, c int) string { return strings.Repeat(s, c) }

func TestTxtDecode(t *testing.T) {
	tests := []struct {
		data     string
		expected []string
	}{
		{``, []string{``}},
		{`""`, []string{``}},
		{`foo`, []string{`foo`}},
		{`"foo"`, []string{`foo`}},
		{`"foo bar"`, []string{`foo bar`}},
		{`foo bar`, []string{`foo`, `bar`}},
		{`"aaa" "bbb"`, []string{`aaa`, `bbb`}},
		{`"a\"a" "bbb"`, []string{`a"a`, `bbb`}},
		// Seen in live traffic:
		{
			"\"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB\"",
			[]string{r("B", 254)},
		},
		{
			"\"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC\"",
			[]string{r("C", 255)},
		},
		{
			"\"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD\" \"D\"",
			[]string{r("D", 255), "D"},
		},
		{
			"\"EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE\" \"EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE\"",
			[]string{r("E", 255), r("E", 255)},
		},
		{
			"\"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF\" \"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF\" \"F\"",
			[]string{r("F", 255), r("F", 255), "F"},
		},
		{
			"\"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\" \"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\" \"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\"",
			[]string{r("G", 255), r("G", 255), r("G", 255)},
		},
		{
			"\"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"H\"",
			[]string{r("H", 255), r("H", 255), r("H", 255), "H"},
		},
		{"\"quo'te\"", []string{`quo'te`}},
		{"\"blah`blah\"", []string{"blah`blah"}},
		{"\"quo\\\"te\"", []string{`quo"te`}},
		{"\"q\\\"uo\\\"te\"", []string{`q"uo"te`}},
		/// Backslashes are meaningless in unquoted strings. Unquoted strings run until they hit a space.
		{`1backs\lash`, []string{`1backs\lash`}},
		{`2backs\\lash`, []string{`2backs\\lash`}},
		{`3backs\\\lash`, []string{`3backs\\\lash`}},
		{`4backs\\\\lash`, []string{`4backs\\\\lash`}},
		/// Inside quotes, a backlash means take the next byte literally.
		{`"q1backs\lash"`, []string{`q1backslash`}},
		{`"q2backs\\lash"`, []string{`q2backs\lash`}},
		{`"q3backs\\\lash"`, []string{`q3backs\lash`}},
		{`"q4backs\\\\lash"`, []string{`q4backs\\lash`}},
		// HETZNER includes a space after the last quote. Make sure we handle that.
		{`"one" "more" `, []string{`one`, `more`}},
		// Edge case: unquoted strings are treated as literals to be joined with no space!
		{`v=spf1 -all`, []string{`v=spf1-all`}},
	}
	for i, test := range tests {
		got, err := txtDecode(test.data)
		if err != nil {
			t.Error(err)
		}

		want := strings.Join(test.expected, "")
		if got != want {
			t.Errorf("%v: expected TxtStrings=(%q) got (%q)", i, want, got)
		}
	}
}

func TestTxtEncode(t *testing.T) {
	tests := []struct {
		data     []string
		expected string
	}{
		{[]string{"simple"}, `"simple"`},
		{[]string{`"quoted"`}, `"\"quoted\""`},
		{[]string{}, `""`},
		{[]string{``}, `""`},
		{[]string{`foo`}, `"foo"`},
		{[]string{`aaa`, `bbb`}, `"aaa" "bbb"`},
		{[]string{`ccc`, `ddd`, `eee`}, `"ccc" "ddd" "eee"`},
		{[]string{`a"a`, `bbb`}, `"a\"a" "bbb"`},
		{[]string{`quo'te`}, "\"quo'te\""},
		{[]string{"blah`blah"}, "\"blah`blah\""},
		{[]string{`quo"te`}, "\"quo\\\"te\""},
		{[]string{`quo"te`}, `"quo\"te"`},
		{[]string{`q"uo"te`}, "\"q\\\"uo\\\"te\""},
		{[]string{`1backs\lash`}, `"1backs\\lash"`},
		{[]string{`2backs\\lash`}, `"2backs\\\\lash"`},
		{[]string{`3backs\\\lash`}, `"3backs\\\\\\lash"`},
		{[]string{strings.Repeat("M", 26), `quo"te`}, `"MMMMMMMMMMMMMMMMMMMMMMMMMM" "quo\"te"`},
	}
	for i, test := range tests {
		got := txtEncode(test.data)

		want := test.expected
		if got != want {
			t.Errorf("%v: expected TxtStrings=v(%v) got (%v)", i, want, got)
		}
	}
}

func TestEncodeSingle(t *testing.T) {
	// EncodeSingle should render the full string as a single quoted chunk,
	// without splitting at 255-octet boundaries (unlike EncodeQuoted).
	long := strings.Repeat("X", 300)
	tests := []struct {
		input    string
		expected string
	}{
		// Short strings behave the same as EncodeQuoted
		{"simple", `"simple"`},
		{`a"b`, `"a\"b"`},
		// A 300-char string must NOT be split into two chunks
		{long, `"` + long + `"`},
	}
	for _, tt := range tests {
		got := EncodeSingle(tt.input)
		if got != tt.expected {
			t.Errorf("EncodeSingle(%q): got %q, want %q", tt.input, got, tt.expected)
		}
		// Verify EncodeQuoted *would* split the 300-char string
		if len(tt.input) == 300 {
			chunked := EncodeQuoted(tt.input)
			if chunked == tt.expected {
				t.Errorf("EncodeQuoted(%d chars): expected chunked output but got single chunk", len(tt.input))
			}
		}
	}
}
